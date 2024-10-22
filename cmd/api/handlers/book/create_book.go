package book

import (
	"back-end-inventory/internal/domains"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func (h *Handler) CreateBook(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Obtener los valores del formulario
	name := r.FormValue("name")
	author := r.FormValue("author")
	pages := r.FormValue("pages")
	description := r.FormValue("description")
	published := r.FormValue("published")

	// Manejar el archivo de imagen
	file, fileHandler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	uploadDir := "uploads"

	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		err := os.Mkdir(uploadDir, os.ModePerm)
		if err != nil {
			http.Error(w, "Error creando directorio de uploads", http.StatusInternalServerError)
			return
		}
	}

	fileName := fmt.Sprintf("%d-%s", time.Now().Unix(), fileHandler.Filename)
	filePath := filepath.Join(uploadDir, fileName)

	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convertir las páginas a int
	pagesInt, err := strconv.Atoi(pages)
	if err != nil {
		http.Error(w, "Invalid number of pages", http.StatusBadRequest)
		return
	}
	serverURL := "http://localhost:8081"

	imageURL := fmt.Sprintf("%s/%s", serverURL, filePath)

	book := &domains.Books{
		Name:        name,
		Author:      author,
		Pages:       pagesInt,
		Description: description,
		Published:   published,
		Image:       imageURL,
	}
	log.Printf("Book: %+v", book)

	newBook, err := h.BookService.InsertBook(*book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(newBook)

}
