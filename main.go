package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	// ServerPort defines the port on which the server will listen
	ServerPort = ":8080"
	// UploadPath defines the directory where uploaded files will be stored
	UploadPath = "./files"
	// MaxUploadSize defines maximum upload size (100 MB)
	MaxUploadSize = 100 << 20
)

// ServerConfig holds the server configuration
type ServerConfig struct {
	logger *log.Logger
}

// FileHandler handles all file-related operations
type FileHandler struct {
	config *ServerConfig
}

// uploadFormTemplate contains the HTML template for the upload form

// NewFileHandler creates a new FileHandler instance
func NewFileHandler(config *ServerConfig) *FileHandler {
	return &FileHandler{
		config: config,
	}
}

// ServeHTTP handles all HTTP requests
func (h *FileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && r.URL.Path == "/":
		h.handleUploadForm(w, r)
	case r.Method == http.MethodPost && r.URL.Path == "/":
		h.handleFileUpload(w, r)
	default:
		h.handleFileDownload(w, r)
	}
}

// handleUploadForm displays the file upload form
func (h *FileHandler) handleUploadForm(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./pages/index.html")
}

// handleFileUpload processes file uploads
func (h *FileHandler) handleFileUpload(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(MaxUploadSize); err != nil {
		h.config.logger.Printf("Error parsing multipart form: %v", err)
		http.Error(w, "File too large", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		h.config.logger.Printf("Error retrieving file: %v", err)
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	filename := r.FormValue("filename")
	if filename == "" {
		filename = header.Filename
	}

	if err := os.MkdirAll(UploadPath, os.ModePerm); err != nil {
		h.config.logger.Printf("Error creating upload directory: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	filepath := filepath.Join(UploadPath, filename)
	dst, err := os.Create(filepath)
	if err != nil {
		h.config.logger.Printf("Error creating file: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		h.config.logger.Printf("Error saving file: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	h.config.logger.Printf("File uploaded successfully: %s", filename)
	fmt.Fprintf(w, "File uploaded successfully: %s", filename)
}

// handleFileDownload serves downloaded files
func (h *FileHandler) handleFileDownload(w http.ResponseWriter, r *http.Request) {
	filepath := filepath.Join(UploadPath, strings.TrimPrefix(r.URL.Path, "/"))

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", GetBaseFileName(filepath)))
	http.ServeFile(w, r, filepath)
}

func main() {
	logger := log.New(os.Stdout, "[FileServer] ", log.LstdFlags|log.Lshortfile)

	config := &ServerConfig{
		logger: logger,
	}

	fileHandler := NewFileHandler(config)

	server := &http.Server{
		Addr:    ServerPort,
		Handler: fileHandler,
	}

	logger.Printf("Server starting on %s", ServerPort)
	if err := server.ListenAndServe(); err != nil {
		logger.Fatalf("Server failed to start: %v", err)
	}
}
