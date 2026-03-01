package utils

import (
	"fmt"
	"os"
	"time"
)

// RegistrarAccion guarda un historial en un archivo de texto plano.
// Está diseñada para ejecutarse como Goroutine (concurrencia).
func RegistrarAccion(accion string, detalle string) {
	// TRUCO ACADÉMICO: Pausamos este hilo por 2 segundos.
	// Esto demuestra que el hilo principal (la página web) NO se bloquea esperando esto.
	time.Sleep(2 * time.Second)

	// Nombre del archivo de log
	nombreArchivo := "historial_auditoria.txt"

	// Abrimos el archivo en modo "Append" (añadir al final). Si no existe, lo crea.
	archivo, err := os.OpenFile(nombreArchivo, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error del sistema de auditoría:", err)
		return
	}
	defer archivo.Close()

	// Formateamos la fecha y hora actual
	horaActual := time.Now().Format("2006-01-02 15:04:05")

	// Preparamos el mensaje que se guardará en el TXT
	mensaje := fmt.Sprintf("[%s] Acción: %s | %s\n", horaActual, accion, detalle)

	// Escribimos en el archivo
	_, err = archivo.WriteString(mensaje)
	if err != nil {
		fmt.Println("Error al escribir el log:", err)
	}
}