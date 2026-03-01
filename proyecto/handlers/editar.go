package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"gestion-libros/database"
	"gestion-libros/models"
)

// MostrarEditarHandler responde a peticiones GET.
// Su función es capturar el ID de la URL, buscar el registro en la BD y precargar el formulario.
func MostrarEditarHandler(w http.ResponseWriter, r *http.Request) {
	// Extracción del parámetro 'id' de la URL (ej. localhost:8080/editar?id=5)
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	// Apertura de conexión a la base de datos
	db, _ := database.Conectar()
	defer db.Close() // Garantiza la liberación del pool de conexiones

	// Consulta del registro específico (Operación Read del CRUD)
	libro, err := models.ObtenerLibroPorID(db, id)
	if err != nil {
		// Manejo de errores: Si el ID no existe o falla la BD, aborta y vuelve al inicio
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Carga de la vista y renderizado
	plantilla, _ := template.ParseFiles("templates/editar.html")

	// Inyección del objeto libro
	plantilla.Execute(w, libro)
}

// Procesa la petición POST enviada desde el formulario de edición
// Su función es sobrescribir los datos del registro existente en MySQL
func ActualizarHandler(w http.ResponseWriter, r *http.Request) {
	// Validación de seguridad(peticiones GET)
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Extracción y parseo de datos del request HTTP
	id, _ := strconv.Atoi(r.FormValue("id"))
	anio, _ := strconv.Atoi(r.FormValue("anio"))

	// Reconstrucción del modelo con los datos modificados por el usuario
	libroEditado := models.Libro{
		ID:      id,
		Titulo:  r.FormValue("titulo"),
		Autor:   r.FormValue("autor"),
		Genero:  r.FormValue("genero"),
		Anio:    anio,
		Formato: r.FormValue("formato"),
	}

	// Conexión a la base de datos
	db, _ := database.Conectar()
	defer db.Close()

	// Update
	models.ActualizarLibro(db, libroEditado)

	// Redirección limpia a la tabla principal
	http.Redirect(w, r, "/", http.StatusSeeOther)
}