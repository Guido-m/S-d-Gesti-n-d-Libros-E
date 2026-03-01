package handlers

import (
	"log"
	"net/http"
	"strconv"

	"gestion-libros/database"
	"gestion-libros/models"
	"gestion-libros/utils"
)

// EliminarHandler procesa la petición de borrado de un registro específico.
// Su función es capturar el ID enviado por la vista, interactuar con el modelo y actualizar la pantalla.
func EliminarHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Extracción de parámetros de la URL
	// Captura el valor de 'id' de la cadena de consulta (ejemplo: localhost:8080/eliminar?id=3)
	idStr := r.URL.Query().Get("id")

	// Validación y parseo de datos
	// Convertimos el String a Entero. Si el usuario intenta modificar la URL con letras
	// (ej: ?id=abc), strconv.Atoi devolverá un error.
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// Mecanismo de seguridad: Si el ID es inválido, abortamos silenciosamente
		// y devolvemos al usuario al inicio.
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Conexión a la base de datos MySQL
	db, _ := database.Conectar()
	defer db.Close() // Garantiza que los recursos se liberen al terminar

	// Delegamos la lógica SQL a la capa de Modelo para mantener el controlador limpio
	err = models.EliminarLibro(db, id)

	err = models.EliminarLibro(db, id)
	if err != nil {
		log.Println("Error de ejecución o integridad al intentar eliminar libro:", err)
	}

	// Lanzamos la concurrencia (Goroutine)
	go utils.RegistrarAccion("ELIMINAR", "Se borró el libro con ID: "+strconv.Itoa(id))

	// 6. Redirección final...
	http.Redirect(w, r, "/", http.StatusSeeOther)
	if err != nil {
		//Manejo de errores del servidor
		log.Println("Error de ejecución o integridad al intentar eliminar libro:", err)
	}

	// Actualizamos la vista enviando al usuario de vuelta a la tabla principal
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
