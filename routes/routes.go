// routes.go
package routes

import (
	"net/http"
	"tiny_files/handlers"

	"github.com/minio/minio-go/v7"
)

func SetupRoutes(minioClient *minio.Client) {
	http.HandleFunc("/get-upload-url", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleGetUploadURL(w, r, minioClient)
	})

	http.HandleFunc("/get-download-url", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleGetDownloadURL(w, r, minioClient)
	})

	// Serve static files
	http.Handle("/", http.FileServer(http.Dir("pages")))
}
