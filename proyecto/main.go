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

	// ARRANCAR EL SERVIDOR WEB
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
