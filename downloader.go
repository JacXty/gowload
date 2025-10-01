// package main

// import (
// 	"fmt"
// 	"os"
// 	"os/exec"
// 	// "github.com/lrstanley/go-ytdlp"
// )

// func main() {
// 	fmt.Println("YouTube Downloader starting...")

// 	// Create a new YouTube client
// 	client := ytdlp.New()

// 	// YouTube video URL
// 	videoURL := "https://www.youtube.com/watch?v=nAksuGJosms"

// 	// Get video info
// 	video, err := client.GetInfo(videoURL)
// 	if err != nil {
// 		fmt.Printf("Error getting video info: %v\n", err)
// 		return
// 	}

// 	// Get the best quality format
// 	stream, _, err := client.GetStream(video, &video.Formats[0])
// 	if err != nil {
// 		fmt.Printf("Error getting video stream: %v\n", err)
// 		return
// 	}
// 	defer stream.Close()

// 	// Create output file
// 	file, err := os.Create("video.mp4")
// 	if err != nil {
// 		fmt.Printf("Error creating file: %v\n", err)
// 		return
// 	}
// 	defer file.Close()

// 	// Download the video
// 	ctx := context.Background()
// 	_, err = client.Download(ctx, video, &video.Formats[0], file)
// 	if err != nil {
// 		fmt.Printf("Error downloading video: %v\n", err)
// 		return
// 	}

//		fmt.Println("Download completed successfully!")
//	}
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func ensureYtDlp() (string, error) {
	// Buscar en el PATH
	if path, err := exec.LookPath("yt-dlp"); err == nil {
		return path, nil
	}

	// Si no está, descargarlo en la carpeta actual
	url := "https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp.exe"
	target := filepath.Join(".", "yt-dlp.exe")

	fmt.Println("⏬ yt-dlp no encontrado, descargando...")

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	out, err := os.Create(target)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}

	// Ruta absoluta del binario descargado
	absPath, err := filepath.Abs(target)
	if err != nil {
		return "", err
	}

	// Hacerlo ejecutable
	if err := os.Chmod(absPath, 0755); err != nil {
		return "", err
	}

	return absPath, nil
}

func main() {
	ytDlpPath, err := ensureYtDlp()
	if err != nil {
		panic(err)
	}

	// URL a descargar
	url := "https://youtu.be/tujjT4ZhwuI?si=_0aEBhTgBcHDkpRK"

	// Crear carpeta music si no existe
	if _, err := os.Stat("music"); os.IsNotExist(err) {
		os.Mkdir("music", 0755)
	}

	// Ejecutar yt-dlp con la ruta absoluta y descargar foto del albun y artista
	cmd := exec.Command(ytDlpPath,
		"-x", "--audio-format", "mp3",
		"-o", "music/%(title)s.%(ext)s",
		"--embed-thumbnail", // Agrega la foto del albun y artista
		url,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		panic(err)
	}

	fmt.Println("✅ Descarga completada en formato MP3 en la carpeta 'music'")
}
