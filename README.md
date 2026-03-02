# 📚 Sistema de Gestión de Libros Electrónicos (Web MVC + API REST)

> Proyecto académico desarrollado en Go (Golang) para la gestión eficiente de una biblioteca digital.

Este sistema ha evolucionado de una interfaz de línea de comandos (CLI) a una **Aplicación Web completa** utilizando el patrón de diseño MVC (Modelo-Vista-Controlador), persistencia de datos real, procesamiento concurrente y servicios web.

## 🚀 Estado del Proyecto (Versión 3.0)

El proyecto ha sido refactorizado exitosamente para operar como un servidor web robusto, cumpliendo con altos estándares de ingeniería de software.

- [x] **Arquitectura MVC:** Separación limpia de responsabilidades (`handlers`, `models`, `templates`, `database`, `utils`).
- [x] **Persistencia Real:** Integración con motor de base de datos MySQL.
- [x] **CRUD Completo:** Interfaz gráfica web para registrar, listar, editar y eliminar libros.
- [x] **Concurrencia (Goroutines):** Sistema de auditoría en segundo plano asíncrono que no bloquea el hilo principal HTTP.
- [x] **API REST (Servicios Web):** 8 endpoints JSON para la futura integración con sistemas externos y aplicaciones móviles.

## 🛠️ Tecnologías y Conceptos Aplicados

El código demuestra el dominio de conceptos avanzados de desarrollo backend en Go:

* **Lenguaje:** Go (Golang) 1.20+
* **Arquitectura:** Patrón MVC (Model-View-Controller) y API REST.
* **Base de Datos:** MySQL / MariaDB utilizando el paquete nativo `database/sql`.
* **Concurrencia:** Uso de `goroutines` para el procesamiento asíncrono de logs.
* **Servicios Web:** Serialización de estructuras a formato universal mediante `encoding/json`.
* **Servidor Web:** Implementación de rutas y servidor HTTP con `net/http`.
* **Frontend (Vistas):** HTML5 y CSS3 renderizado dinámicamente desde el backend con `html/template`.

## 📂 Estructura del Proyecto

```text
/
├── database/      # Lógica de conexión y configuración del pool de MySQL
├── handlers/      # Controladores HTTP (Vistas web) y Controladores API (JSON)
├── models/        # Estructuras de datos (Structs) y consultas SQL (Lógica de negocio)
├── templates/     # Archivos HTML con sintaxis de Go Templates para la interfaz visual
├── utils/         # Funciones auxiliares y sistema de auditoría concurrente
└── main.go        # Entry point: Inicializa la BD, define el enrutador y arranca el servidor
## 💻 Instalación y Ejecución

Para correr este proyecto localmente, necesitas tener instalado **Go (Golang)** y un servidor **MySQL** (puedes usar herramientas como XAMPP, WAMP, o Docker).

### Paso 1: Clonar el repositorio
Abre tu terminal y ejecuta los siguientes comandos:
```bash
git clone [https://github.com/Guido-m/S-d-Gesti-n-d-Libros-E.git](https://github.com/Guido-m/S-d-Gesti-n-d-Libros-E.git)
cd S-d-Gesti-n-d-Libros-E
```
-- 1. Crear la base de datos
```bash
CREATE DATABASE IF NOT EXISTS gestion_libros;
USE gestion_libros;

-- 2. Crear la tabla principal
CREATE TABLE IF NOT EXISTS libros (
    id INT AUTO_INCREMENT PRIMARY KEY,
    titulo VARCHAR(255) NOT NULL,
    autor VARCHAR(255) NOT NULL,
    genero VARCHAR(100),
    anio INT,
    formato VARCHAR(50)
);
```
