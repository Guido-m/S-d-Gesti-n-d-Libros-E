# Sistema de Gestión de Libros Electrónicos

> Proyecto académico desarrollado en Go (Golang) para la gestión eficiente de una biblioteca digital.

Este sistema permite administrar un inventario de libros electrónicos a través de una interfaz de línea de comandos (CLI).

## 🚀 Avance Actual (Estado del Proyecto)

El proyecto se encuentra en una fase de refactorización y mejora arquitectónica.

- [x] **Arquitectura Base:** Implementación de Structs e Interfaces.
- [x] **Persistencia en Memoria:** Uso de Slices dinámicos.
- [x] **Módulo de Registro:** Validación y guardado de datos.
- [x] **Módulo de Consulta:** Listado general de libros.
- [ ] **Módulo de Actualización:** Edición de metadatos (En desarrollo).
- [ ] **Módulo de Eliminación:** Borrado de registros (En desarrollo).

## 🛠️ Tecnologías y Conceptos Aplicados

El código ha evolucionado para incluir conceptos avanzados de la Unidad 2 y 3:

* **Lenguaje:** Go (Golang) 1.20+
* **Paradigma:** Programación Estructurada y Orientada a Objetos (Simulada).
* **Estructuras de Datos:** `Structs` para modelar Libros y `Slices` para el inventario.
* **Abstracción:** Uso de `Interfaces` para definir el contrato `IGestionLibros`.
* **Control de Errores:** Implementación nativa de `error`.
* **Encapsulamiento:** Protección de campos mediante getters y constructores.

## 📋 Funcionalidades

### 1. Registrar Libro
Permite ingresar un nuevo libro validando que los datos sean correctos (ej. año numérico, título no vacío).

### 2. Consultar Libros
Muestra un listado formateado de todos los libros cargados en la sesión actual.

### 3. Gestión (Próximamente)
Las funciones de actualizar y eliminar se encuentran definidas en la interfaz y listas para su implementación lógica.

## 💻 Instalación y Ejecución

Para correr este proyecto localmente:

1. Clonar el repositorio:
   ```bash
   git clone [https://github.com/Guido-m/S-d-Gesti-n-d-Libros-E.git](https://github.com/Guido-m/S-d-Gesti-n-d-Libros-E.git)
