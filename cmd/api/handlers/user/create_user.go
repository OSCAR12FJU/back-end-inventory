package user

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

	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) CreateUSer(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Obtener los valores del formulario
	name := r.FormValue("name")
	lastName := r.FormValue("last_name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	age := r.FormValue("age")
	nacionality := r.FormValue("nacionality")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error al encriptar la contraseña", http.StatusInternalServerError)
		return
	}

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
	ageInt, err := strconv.Atoi(age)
	if err != nil {
		http.Error(w, "Invalid number of pages", http.StatusBadRequest)
		return
	}
	serverURL := "http://localhost:8081"

	imageURL := fmt.Sprintf("%s/%s", serverURL, filePath)

	user := &domains.Users{
		Name:        name,
		LastName:    lastName,
		Email:       email,
		Password:    string(hashedPassword),
		Age:         ageInt,
		Nacionality: nacionality,
		Image:       imageURL,
	}

	log.Printf("Usuarios: %+v", user)

	newBook, err := h.UserService.InsertUser(*user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(newBook)

}
