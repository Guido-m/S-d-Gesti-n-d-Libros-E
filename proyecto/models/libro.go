package models

import (
	"database/sql"
	"fmt"
)

// Libro representa la estructura de datos en la base de datos.
type Libro struct {
	ID      int
	Titulo  string
	Autor   string
	Genero  string
	Anio    int
	Formato string
}

// ObtenerTodos ejecuta el SELECT y devuelve un Slice de libros.
func ObtenerTodos(db *sql.DB) ([]Libro, error) {
	query := "SELECT id, titulo, autor, genero, anio, formato FROM libros"
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al consultar BD: %v", err)
	}
	defer rows.Close()

	var listaLibros []Libro

	for rows.Next() {
		var l Libro
		err := rows.Scan(&l.ID, &l.Titulo, &l.Autor, &l.Genero, &l.Anio, &l.Formato)
		if err == nil {
			listaLibros = append(listaLibros, l)
		}
	}
	return listaLibros, nil
}

// InsertarLibro guarda un nuevo registro en la base de datos MySQL.
func InsertarLibro(db *sql.DB, l Libro) error {
	query := `INSERT INTO libros (titulo, autor, genero, anio, formato) VALUES (?, ?, ?, ?, ?)`
	_, err := db.Exec(query, l.Titulo, l.Autor, l.Genero, l.Anio, l.Formato)
	return err
}

// Agregar al final de models/libro.go

// ObtenerLibroPorID busca un solo libro en la BD para poder editarlo
func ObtenerLibroPorID(db *sql.DB, id int) (Libro, error) {
	var l Libro
	query := "SELECT id, titulo, autor, genero, anio, formato FROM libros WHERE id = ?"
	err := db.QueryRow(query, id).Scan(&l.ID, &l.Titulo, &l.Autor, &l.Genero, &l.Anio, &l.Formato)
	return l, err
}

// ActualizarLibro guarda los cambios de un libro editado
func ActualizarLibro(db *sql.DB, l Libro) error {
	query := "UPDATE libros SET titulo=?, autor=?, genero=?, anio=?, formato=? WHERE id=?"
	_, err := db.Exec(query, l.Titulo, l.Autor, l.Genero, l.Anio, l.Formato, l.ID)
	return err
}

// EliminarLibro borra un libro de la base de datos
func EliminarLibro(db *sql.DB, id int) error {
	query := "DELETE FROM libros WHERE id = ?"
	_, err := db.Exec(query, id)
	return err
}
