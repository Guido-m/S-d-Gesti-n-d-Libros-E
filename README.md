# 📚 Sistema de Gestión de Libros Electrónicos (Web MVC)

> Proyecto académico desarrollado en Go (Golang) para la gestión eficiente de una biblioteca digital.

Este sistema ha evolucionado de una interfaz de línea de comandos (CLI) a una **Aplicación Web completa** utilizando el patrón de diseño MVC (Modelo-Vista-Controlador) y persistencia de datos real.

## 🚀 Estado del Proyecto (Versión 2.0)

El proyecto ha sido refactorizado exitosamente para operar como un servidor web con conexión a base de datos.

- [x] **Arquitectura MVC:** Separación limpia de responsabilidades (`handlers`, `models`, `templates`, `database`).
- [x] **Persistencia Real:** Integración con motor de base de datos MySQL.
- [x] **Operación Create:** Formulario web para registro de nuevos libros.
- [x] **Operación Read:** Tabla dinámica y responsiva que lista el inventario actual.
- [x] **Operación Update:** Vista de edición con precarga de datos existentes.
- [x] **Operación Delete:** Eliminación de registros con validación de seguridad (UX).

## 🛠️ Tecnologías y Conceptos Aplicados

El código demuestra el dominio de conceptos avanzados de desarrollo backend en Go:

* **Lenguaje:** Go (Golang) 1.20+
* **Arquitectura:** Patrón MVC (Model-View-Controller).
* **Base de Datos:** MySQL / MariaDB utilizando el paquete nativo `database/sql`.
* **Servidor Web:** Implementación de rutas y servidor con `net/http`.
* **Frontend (Vistas):** HTML5 y CSS3 moderno (Variables CSS, Flexbox, UI Cards) renderizado desde el backend con `html/template`.
* **Seguridad:** Uso de sentencias SQL preparadas (`?`) para evitar inyecciones SQL.

## 📂 Estructura del Proyecto

```text
/
├── database/      # Lógica de conexión y configuración del pool de MySQL
├── handlers/      # Controladores HTTP que orquestan peticiones y respuestas
├── models/        # Estructuras de datos (Structs) y consultas SQL (Lógica de negocio)
├── templates/     # Archivos HTML con sintaxis de Go Templates para la interfaz de usuario
└── main.go        # Entry point: Inicializa la BD, define el enrutador y arranca el servidor

## 💻 Instalación y Ejecución

Para correr este proyecto localmente, necesitas tener instalado **Go (Golang)** y un servidor **MySQL** (puedes usar herramientas como XAMPP, WAMP, o Docker).

### Paso 1: Clonar el repositorio
Abre tu terminal y ejecuta los siguientes comandos:
```bash
git clone [https://github.com/Guido-m/S-d-Gesti-n-d-Libros-E.git](https://github.com/Guido-m/S-d-Gesti-n-d-Libros-E.git)
cd S-d-Gesti-n-d-Libros-E
