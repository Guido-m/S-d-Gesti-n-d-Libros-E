package database

import (
	"database/sql"
	"fmt"

	// El driver se importa aquí porque esta es la capa de datos
	_ "github.com/go-sql-driver/mysql"
)

func Conectar() (*sql.DB, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/biblioteca"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("fallo al inicializar el driver: %v", err)
	}

	// Verificamos conectividad real
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("no hay comunicación con MySQL. Verifique XAMPP: %v", err)
	}

	fmt.Println("--- Conexión a BD MySQL establecida desde el paquete 'database' ---")
	return db, nil
}