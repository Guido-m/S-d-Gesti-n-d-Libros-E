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

// ConsultarLibros muestra el listado actual o permite buscar por coincidencias.
// Implementa un sub-menú para manejar diferentes tipos de consulta.
func (s *SistemaGestion) ConsultarLibros() {
	if len(s.inventario) == 0 {
		fmt.Println("\nNo hay libros registrados en el sistema.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n--- Módulo de Consulta ---")
	fmt.Println("1. Listar todos los libros")
	fmt.Println("2. Buscar por título o autor")
	fmt.Print("Seleccione una opción: ")

	// 1. Captura de la opción de consulta
	opcion, _ := reader.ReadString('\n')
	opcion = strings.TrimSpace(opcion)

	// 2. Control de flujo según el tipo de consulta
	if opcion == "1" {
		fmt.Println("\nListado completo de Libros:")
		for _, libro := range s.inventario {
			fmt.Println(libro.ObtenerDetalles())
		}
	} else if opcion == "2" {
		fmt.Print("Ingrese el término de búsqueda (título o autor): ")
		termino, _ := reader.ReadString('\n')
		termino = strings.ToLower(strings.TrimSpace(termino)) // Normalizamos a minúsculas para mejorar la búsqueda

		encontrados := 0
		fmt.Println("\nResultados de la búsqueda:")

		// 3. Recorrido del slice filtrando por coincidencias
		for _, libro := range s.inventario {
			tituloLower := strings.ToLower(libro.titulo)
			autorLower := strings.ToLower(libro.autor)

			if strings.Contains(tituloLower, termino) || strings.Contains(autorLower, termino) {
				fmt.Println(libro.ObtenerDetalles())
				encontrados++
			}
		}

		// 4. Validación si no hubo resultados
		if encontrados == 0 {
			fmt.Println("No se encontraron coincidencias.")
		}
	} else {
		fmt.Println("Opción de consulta inválida.")
	}
}

// ActualizarLibro permite modificar la información de un libro existente mediante su ID.
// Valida la existencia del registro y actualiza solo los campos proporcionados.
func (s *SistemaGestion) ActualizarLibro() error {
	if len(s.inventario) == 0 {
		return errors.New("el inventario está vacío, no hay nada que actualizar")
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n--- Módulo de Actualización ---")

	// 1. Captura y validación del ID
	fmt.Print("Ingrese el ID del libro a modificar: ")
	idStr, _ := reader.ReadString('\n')
	idBuscado, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		return errors.New("el ID debe ser un valor numérico")
	}

	// 2. Búsqueda del libro en el Slice de memoria
	var libroEncontrado *Libro
	for _, libro := range s.inventario {
		if libro.id == idBuscado {
			libroEncontrado = libro
			break
		}
	}

	// 3. Validación de existencia
	if libroEncontrado == nil {
		return errors.New("no se encontró ningún libro con ese ID")
	}

	// 4. Modificación de campos (Permite dejar en blanco para mantener el actual)
	fmt.Println("Deje el campo en blanco y presione Enter si no desea modificarlo.")

	fmt.Printf("Nuevo Título (Actual: %s): ", libroEncontrado.titulo)
	nuevoTitulo, _ := reader.ReadString('\n')
	nuevoTitulo = strings.TrimSpace(nuevoTitulo)
	if nuevoTitulo != "" {
		libroEncontrado.titulo = nuevoTitulo
	}

	fmt.Printf("Nuevo Autor (Actual: %s): ", libroEncontrado.autor)
	nuevoAutor, _ := reader.ReadString('\n')
	nuevoAutor = strings.TrimSpace(nuevoAutor)
	if nuevoAutor != "" {
		libroEncontrado.autor = nuevoAutor
	}

	fmt.Println(">> Libro actualizado correctamente.")
	return nil
}

// EliminarLibro permite borrar un libro del sistema mediante su ID.
// Utiliza la manipulación de Slices nativa de Go para remover el registro.
func (s *SistemaGestion) EliminarLibro() error {
	if len(s.inventario) == 0 {
		return errors.New("el inventario está vacío, no hay nada que eliminar")
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n--- Módulo de Eliminación ---")

	// 1. Captura y validación del ID a eliminar
	fmt.Print("Ingrese el ID del libro a eliminar: ")
	idStr, _ := reader.ReadString('\n')
	idBuscado, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		return errors.New("el ID debe ser un valor numérico")
	}

	// 2. Búsqueda del índice exacto del libro en el Slice
	indiceAEliminar := -1
	for i, libro := range s.inventario {
		if libro.id == idBuscado {
			indiceAEliminar = i
			break
		}
	}

	// 3. Validación de existencia
	if indiceAEliminar == -1 {
		return errors.New("no se encontró ningún libro con ese ID")
	}

	// 4. Reestructuración del Slice excluyendo el elemento encontrado
	s.inventario = append(s.inventario[:indiceAEliminar], s.inventario[indiceAEliminar+1:]...)
	fmt.Println(">> Libro eliminado del sistema exitosamente.")
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
