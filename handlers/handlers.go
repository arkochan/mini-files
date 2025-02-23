package handlers

import (
	"fmt"
	"net/http"
	"tiny_files/config"
	"tiny_files/services"

	"github.com/minio/minio-go/v7"
)

func HandleGetUploadURL(w http.ResponseWriter, r *http.Request, minioClient *minio.Client) {
	filename := r.URL.Query().Get("filename")
	if filename == "" {
		http.Error(w, "filename is required", http.StatusBadRequest)
		return
	}

	url, err := services.GenerateUploadURL(minioClient, config.BucketName, filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, url)
}

func HandleGetDownloadURL(w http.ResponseWriter, r *http.Request, minioClient *minio.Client) {
	filename := r.URL.Query().Get("filename")
	if filename == "" {
		http.Error(w, "filename is required", http.StatusBadRequest)
		return
	}

	url, err := services.GenerateDownloadURL(minioClient, config.BucketName, filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, url)
}

func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/favicon.ico")
}
