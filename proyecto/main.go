package main

import (
	"fmt"
	"log"
	"net/http"

	// Importamos los paquetes
	"gestion-libros/database"
	"gestion-libros/handlers"
)

func main() {
	// INICIALIZAR LA DB
	db, err := database.Conectar()
	if err != nil {
		log.Fatalf("ERROR CRÍTICO: No se pudo conectar a MySQL: %v\n", err)
	}
	defer db.Close()

	// CONFIGURAR EL ENRUTADOR

	http.HandleFunc("/", handlers.HomeHandler)                 // Tabla principal
	http.HandleFunc("/crear", handlers.CrearHandler)           // Mostrar crear
	http.HandleFunc("/guardar", handlers.GuardarHandler)       // Guardar nuevo
	http.HandleFunc("/eliminar", handlers.EliminarHandler)     // Borrar
	http.HandleFunc("/editar", handlers.MostrarEditarHandler)  // Mostrar editar
	http.HandleFunc("/actualizar", handlers.ActualizarHandler) // Guardar cambios

	// RUTAS DE LOS SERVICIOS WEB (JSON API) - ¡Lo Nuevo!
	http.HandleFunc("/api/libros", handlers.APIObtenerTodos)       // Servicio 1
	http.HandleFunc("/api/libro", handlers.APIObtenerPorID)        // Servicio 2
	http.HandleFunc("/api/eliminar", handlers.APIEliminarLibro)    // Servicio 3
	http.HandleFunc("/api/estadistica", handlers.APITotalLibros)   // Servicio 4
	http.HandleFunc("/api/estado", handlers.APIEstadoServidor)     // Servicio 5
	http.HandleFunc("/api/info", handlers.APIInfoProyecto)         // Servicio 6
	http.HandleFunc("/api/filtrar", handlers.APIFiltrarPorFormato) // Servicio 7
	http.HandleFunc("/api/crear", handlers.APICrearLibro)          // Servicio 8

	fmt.Println("======================================================")
	fmt.Println("SISTEMA DE GESTIÓN DE LIBROS ELECTRÓNICOS (WEB) ")
	fmt.Println("======================================================")
	fmt.Println("El servidor está corriendo.")
	fmt.Println("Abre tu navegador y visita: http://localhost:8080")
	fmt.Println("======================================================")

	// Bloquea el programa y se queda escuchando el puerto 8080
	errServidor := http.ListenAndServe(":8080", nil)
	if errServidor != nil {
		log.Fatal("Error al iniciar el servidor: ", errServidor)
	}
}
