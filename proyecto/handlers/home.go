package handlers

import (
	"html/template"
	"log"
	"net/http"

	"gestion-libros/database"
	"gestion-libros/models"
)

// Handler es el controlador principal que responde a la ruta raíz
// Actúa como orquestador: extrae los datos del Modelo y los inyecta en la Vista
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Establecer conexión MySQL
	db, err := database.Conectar()
	if err != nil {
		// Retorna un código HTTP 500 (Internal Server Error) si la base de datos está caída
		http.Error(w, "Error crítico conectando a la base de datos", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Ejecutar la obtención de datos Consulta SELECT
	// Delegamos la complejidad de la consulta SQL a nuestro paquete 'models'
	libros, err := models.ObtenerTodos(db)
	if err != nil {
		// Si hay un error SQL, lo registramos internamente para el desarrollador,
		// pero permitimos que la página cargue (probablemente mostrando una tabla vacía)
		log.Println("Fallo al ejecutar la consulta de libros:", err)
	}

	// ParseFiles lee el archivo HTML del disco y lo prepara para recibir datos dinámicos
	plantilla, err := template.ParseFiles("templates/base.html")
	if err != nil {
		http.Error(w, "Error del servidor al intentar cargar la interfaz principal", http.StatusInternalServerError)
		return
	}

	// Renderizado dinámico y respuesta al cliente
	// Execute fusiona el HTML estático con nuestro Slice de Go ('libros')
	// y envía el resultado final compilado al navegador del usuario
	plantilla.Execute(w, libros)
}