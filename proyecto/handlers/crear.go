package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"gestion-libros/database"
	"gestion-libros/models"
	"gestion-libros/utils"
)

// CrearHandler responde a peticiones GET renderizando la vista del formulario HTML.
func CrearHandler(w http.ResponseWriter, r *http.Request) {
	plantilla, err := template.ParseFiles("templates/crear.html")
	if err != nil {
		http.Error(w, "Error interno al cargar la interfaz visual", http.StatusInternalServerError)
		return
	}

	// Ejecuta y envía la plantilla vacía al navegador
	plantilla.Execute(w, nil)
}

// GuardarHandler procesa la petición POST del formulario e inserta el registro en MySQL.
func GuardarHandler(w http.ResponseWriter, r *http.Request) {
	// Validación de seguridad: Asegura que la ruta solo acepte el envío del formulario
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Extracción de datos del request (coinciden con el atributo 'name' del HTML)
	titulo := r.FormValue("titulo")
	autor := r.FormValue("autor")
	genero := r.FormValue("genero")
	anioStr := r.FormValue("anio")
	formato := r.FormValue("formato")

	// Parseo de tipos de datos (De string HTTP a int nativo de Go)
	anio, _ := strconv.Atoi(anioStr)

	// Instanciación del modelo Libro
	nuevoLibro := models.Libro{
		Titulo:  titulo,
		Autor:   autor,
		Genero:  genero,
		Anio:    anio,
		Formato: formato,
	}

	// Conexión a la base de datos
	db, _ := database.Conectar()
	defer db.Close()

	// Ejecución de la operación Create (C del CRUD)
	err := models.InsertarLibro(db, nuevoLibro)
	if err != nil {
		log.Println("Fallo de integridad o conexión al guardar en BD:", err)
		http.Error(w, "Error crítico al guardar el registro", http.StatusInternalServerError)
		return
	}
	// Lanzamos la concurrencia (Goroutine)
	go utils.RegistrarAccion("CREAR", "Se registró un nuevo libro: "+titulo)

	// 5. Redirigir al usuario...
	http.Redirect(w, r, "/", http.StatusSeeOther)
	// 6. Patrón PRG (Post/Redirect/Get): Redirige al inicio para evitar reenvíos duplicados
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
