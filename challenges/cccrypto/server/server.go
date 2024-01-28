package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed web/*
var webContent embed.FS

func main() {
	// Use the subdirectory web as the root of the file system.
	contentStatic, err := fs.Sub(webContent, "web")
	if err != nil {
		log.Fatal(err)
	}

	// Create a file server handler that serves the embedded files.
	http.Handle("/", http.FileServer(http.FS(contentStatic)))

	// Start the server on port 8080.
	log.Println("Listening on :3100...")
	err = http.ListenAndServe(":3100", nil)
	if err != nil {
		log.Fatal(err)
	}
}
