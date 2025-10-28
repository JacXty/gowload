# 🎶 Descargador de Música de YouTube en Go

## 📝 Descripción del Proyecto

Este proyecto es una herramienta simple y eficiente escrita en **Go** que permite la descarga de audio (música) desde enlaces de YouTube. Utiliza la poderosa librería `go-ytdlp`, que a su vez se basa en la funcionalidad del popular `yt-dlp`, para manejar la lógica de descarga y extracción de audio.

---

## 🚀 Requisitos para la Implementación

Para que este proyecto funcione correctamente, son necesarios los siguientes componentes en tu sistema:

1.  **Go (Golang):** El entorno de ejecución y compilador de Go.
    * **Versión Requerida:** `go 1.24.0` o superior.
2.  **yt-dlp:** El binario de `yt-dlp` (o `youtube-dl`), una herramienta de línea de comandos para descargar videos y audio. **La librería de Go se basa en que este binario esté disponible en el `PATH` de tu sistema.**

> **Instalación de yt-dlp:** Puedes encontrar instrucciones detalladas en el [repositorio oficial de yt-dlp](https://github.com/yt-dlp/yt-dlp).

---

## 🛠️ Librerías de Go Utilizadas

Este proyecto solo depende de una librería principal para la funcionalidad de descarga:

| Librería | Versión | Descripción |
| :--- | :--- | :--- |
| **`github.com/lrstanley/go-ytdlp`** | `v1.2.4` | La librería principal que proporciona una interfaz Go para el binario `yt-dlp`, permitiendo la descarga y la extracción de información del enlace. |
| *Dependencias Indirectas* | | Son librerías que `go-ytdlp` requiere internamente, principalmente para tareas de criptografía y manejo de archivos. |

---

## ⚙️ Cómo Ejecutar el Proyecto

Sigue estos pasos para clonar y ejecutar el proyecto:

### 1. Clonar el Repositorio

```bash
git clone URL_DEL_REPOSITORIO
```

### 2. Instalar las Dependencias

```bash
go mod tidy
```

### 3. Asegurar la Disponibilidad de yt-dlp

Verifica que el binario yt-dlp esté instalado y accesible desde la línea de comandos:

```bash
yt-dlp --version
```

### 4. Compilar y Ejecutar

Una vez que tengas tu archivo principal de Go (por ejemplo, downloader.go), puedes compilar y ejecutar tu aplicación.

```bash
go run .
```