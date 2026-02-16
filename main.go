package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// =================================================================================
// ESTRUCTURAS DE DATOS Y ENCAPSULAMIENTO
// =================================================================================

// Libro define la estructura de datos para un libro electrónico.
// Los campos inician con minúscula para que sean privados (encapsulados) dentro del paquete.
// Esto obliga a utilizar métodos o constructores para acceder o modificar la información.
type Libro struct {
	id      int
	titulo  string
	autor   string
	genero  string
	anio    int
	formato string // Ejemplo: PDF, EPUB, MOBI
}

// NuevoLibro es el constructor que permite crear instancias de 'Libro' de forma segura,
// inicializando todos sus campos correctamente.
func NuevoLibro(id int, titulo, autor, genero string, anio int, formato string) *Libro {
	return &Libro{
		id:      id,
		titulo:  titulo,
		autor:   autor,
		genero:  genero,
		anio:    anio,
		formato: formato,
	}
}

// ObtenerDetalles es un método (Getter) que permite leer la información del libro
// formateada en una cadena de texto, sin exponer los campos directamente.
func (l *Libro) ObtenerDetalles() string {
	return fmt.Sprintf("ID: %d | Título: %s | Autor: %s | Año: %d | Formato: %s",
		l.id, l.titulo, l.autor, l.anio, l.formato)
}

// =================================================================================
// INTERFAZ Y GESTIÓN
// =================================================================================

// IGestionLibros define el contrato de métodos que debe tener el sistema.
// Esto permite abstracción: al programa principal no le importa cómo se hacen las cosas,
// solo sabe que el sistema "sabe" hacer estas 4 acciones.
type IGestionLibros interface {
	RegistrarLibro() error
	ConsultarLibros()
	ActualizarLibro() error
	EliminarLibro() error
}

// SistemaGestion implementa la interfaz IGestionLibros.
// Contiene un Slice (lista dinámica) para almacenar los libros en memoria.
type SistemaGestion struct {
	inventario []*Libro
	ultimoID   int // Contador interno para asignar IDs únicos automáticamente
}

// =================================================================================
// LÓGICA DE LOS MÓDULOS
// =================================================================================

// RegistrarLibro solicita los datos al usuario, valida la entrada y guarda el libro.
// Retorna un error si algo falla durante el proceso.
func (s *SistemaGestion) RegistrarLibro() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n--- Módulo de Registro de Libros ---")

	// 1. Captura y limpieza de datos
	fmt.Print("Ingrese Título: ")
	titulo, _ := reader.ReadString('\n')
	titulo = strings.TrimSpace(titulo)

	// Validación: El título es obligatorio
	if titulo == "" {
		return errors.New("el título no puede estar vacío")
	}

	fmt.Print("Ingrese Autor: ")
	autor, _ := reader.ReadString('\n')
	autor = strings.TrimSpace(autor)

	fmt.Print("Ingrese Género: ")
	genero, _ := reader.ReadString('\n')
	genero = strings.TrimSpace(genero)

	fmt.Print("Ingrese Año de publicación: ")
	anioStr, _ := reader.ReadString('\n')
	anio, err := strconv.Atoi(strings.TrimSpace(anioStr))
	if err != nil {
		return errors.New("el año debe ser un valor numérico válido")
	}

	fmt.Print("Ingrese Formato (PDF, EPUB, MOBI): ")
	formato, _ := reader.ReadString('\n')
	formato = strings.TrimSpace(formato)

	// 2. Creación del objeto y actualización del estado
	s.ultimoID++
	nuevoLibro := NuevoLibro(s.ultimoID, titulo, autor, genero, anio, formato)

	// 3. Almacenamiento en la lista dinámica (Slice)
	s.inventario = append(s.inventario, nuevoLibro)

	fmt.Printf(">> ¡Éxito! Libro '%s' registrado con ID %d.\n", titulo, s.ultimoID)
	return nil
}

// ConsultarLibros muestra el listado actual de libros registrados.
// (Actualmente en fase de construcción básica para verificar funcionalidad).
func (s *SistemaGestion) ConsultarLibros() {
	fmt.Println("\n--- Módulo de Consulta ---")
	if len(s.inventario) == 0 {
		fmt.Println("No hay libros registrados en el sistema.")
		return
	}

	fmt.Println("Listado de Libros:")
	for _, libro := range s.inventario {
		fmt.Println(libro.ObtenerDetalles())
	}
}

// ActualizarLibro permite modificar la información de un libro existente.
// (Estructura definida, lógica pendiente de implementación detallada).
func (s *SistemaGestion) ActualizarLibro() error {
	fmt.Println("\n--- Módulo de Actualización (En construcción) ---")
	return nil
}

// EliminarLibro permite borrar un libro del sistema mediante su ID.
// (Estructura definida, lógica pendiente de implementación detallada).
func (s *SistemaGestion) EliminarLibro() error {
	fmt.Println("\n--- Módulo de Eliminación (En construcción) ---")
	return nil
}

// =================================================================================
// FUNCIÓN PRINCIPAL
// =================================================================================

func main() {
	// Inicialización del sistema con una lista vacía
	var app IGestionLibros = &SistemaGestion{
		inventario: []*Libro{},
		ultimoID:   0,
	}

	reader := bufio.NewReader(os.Stdin)
	ejecutando := true

	// Bucle principal para mantener el menú activo
	for ejecutando {
		fmt.Println("\n========================================")
		fmt.Println(" SISTEMA DE GESTIÓN DE LIBROS ELECTRÓNICOS ")
		fmt.Println("========================================")
		fmt.Println("1. Registrar Libro")
		fmt.Println("2. Consultar Libros")
		fmt.Println("3. Actualizar Libro")
		fmt.Println("4. Eliminar Libro")
		fmt.Println("5. Salir")
		fmt.Print("Seleccione una opción: ")

		input, _ := reader.ReadString('\n')
		opcion := strings.TrimSpace(input)

		var err error

		// Control de flujo según la opción seleccionada
		switch opcion {
		case "1":
			err = app.RegistrarLibro()
		case "2":
			app.ConsultarLibros()
		case "3":
			err = app.ActualizarLibro()
		case "4":
			err = app.EliminarLibro()
		case "5":
			fmt.Println("Cerrando el sistema...")
			ejecutando = false
		default:
			fmt.Println("Opción no válida, intente nuevamente.")
		}

		// Manejo centralizado de errores
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}