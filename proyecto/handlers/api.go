package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gestion-libros/database"
	"gestion-libros/models"
)

// SERVICIO 1: Obtener todos los libros (GET)
func APIObtenerTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Avisamos que responderemos en JSON
	db, _ := database.Conectar()
	defer db.Close()

	libros, _ := models.ObtenerTodos(db)
	json.NewEncoder(w).Encode(libros) // Magia: Convierte la lista de Go a JSON automáticamente
}

// SERVICIO 2: Obtener un libro por ID (GET)
func APIObtenerPorID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	db, _ := database.Conectar()
	defer db.Close()

	libro, err := models.ObtenerLibroPorID(db, id)
	if err != nil {
		http.Error(w, `{"error": "Libro no encontrado"}`, http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(libro)
}

// SERVICIO 3: Eliminar un libro (DELETE)
func APIEliminarLibro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	db, _ := database.Conectar()
	defer db.Close()

	models.EliminarLibro(db, id)
	// Respondemos con un JSON de confirmación
	json.NewEncoder(w).Encode(map[string]string{"mensaje": "Libro eliminado correctamente"})
}

// SERVICIO 4: Obtener el total de libros (Estadística) (GET)
func APITotalLibros(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, _ := database.Conectar()
	defer db.Close()

	libros, _ := models.ObtenerTodos(db)
	total := len(libros) // Contamos cuántos hay en la lista
	json.NewEncoder(w).Encode(map[string]int{"total_libros": total})
}

// SERVICIO 5: Estado del Servidor (Healthcheck) (GET)
func APIEstadoServidor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	respuesta := map[string]string{
		"estado": "Online",
		"version": "2.0 MVC",
		"concurrencia": "Activa",
	}
	json.NewEncoder(w).Encode(respuesta)
}

// SERVICIO 6: Información del Proyecto (GET)
func APIInfoProyecto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	info := map[string]string{
		"materia": "Programación Orientada a Objetos",
		"tema": "Gestión de Libros Electrónicos",
		"lenguaje": "Golang",
	}
	json.NewEncoder(w).Encode(info)
}

// SERVICIO 7: Buscar por formato PDF o EPUB (Filtro simulado) (GET)
// En la vida real haríamos un SELECT WHERE formato=?, aquí filtramos en memoria por simplicidad
func APIFiltrarPorFormato(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	formatoBuscado := r.URL.Query().Get("formato")

	db, _ := database.Conectar()
	defer db.Close()
	libros, _ := models.ObtenerTodos(db)

	var filtrados []models.Libro
	for _, lib := range libros {
		if lib.Formato == formatoBuscado {
			filtrados = append(filtrados, lib)
		}
	}
	json.NewEncoder(w).Encode(filtrados)
}

// SERVICIO 8: Crear Libro vía JSON (POST)
func APICrearLibro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, `{"error": "Método no permitido, use POST"}`, http.StatusMethodNotAllowed)
		return
	}

	var nuevoLibro models.Libro
	// Decodificamos el JSON que envía el cliente y lo metemos en nuestro Struct
	err := json.NewDecoder(r.Body).Decode(&nuevoLibro)
	if err != nil {
		http.Error(w, `{"error": "JSON inválido"}`, http.StatusBadRequest)
		return
	}

	db, _ := database.Conectar()
	defer db.Close()

	models.InsertarLibro(db, nuevoLibro)
	json.NewEncoder(w).Encode(map[string]string{"mensaje": "Libro creado exitosamente vía API"})
}