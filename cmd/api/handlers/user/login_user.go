package user

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	// var loginUser struct {
	// 	Email    string `json:"email"`
	// 	Password string `json:"password"`
	// }

	email := r.FormValue("email")
	password := r.FormValue("password")
	if email == "" || password == "" {
		http.Error(w, "El correo y la contraseña son requeridos", http.StatusBadRequest)
		return
	}

	// err := json.NewDecoder(r.Body).Decode(&loginUser)
	// if err != nil {
	// 	http.Error(w, "Datos invalidos", http.StatusBadRequest)
	// 	return
	// }

	user, err := h.UserService.FindUserByEmail(email)
	if err != nil {
		http.Error(w, "Correo o contraseña incorrectos", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		http.Error(w, "Correo o contraseña incorrectos", http.StatusUnauthorized)
		return
	}

	// loginResponse := &domains.LoginResponse{
	// 	Message: "Inicio de sesion exitoso",
	// 	Token:   user.Token,
	// }
	// fmt.Println(loginResponse)

	response := map[string]interface{}{
		"message": "Inicio de sesión exitoso",
		"user":    user,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error al generar la respuesta", http.StatusInternalServerError)
	}

	// if err := json.NewEncoder(w).Encode(user); err != nil {
	// 	http.Error(w, "Error al generar la respuesta", http.StatusInternalServerError)
	// }
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(user)

}
