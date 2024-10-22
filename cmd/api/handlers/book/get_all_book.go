package book

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) GetAllBook(w http.ResponseWriter, r *http.Request) {
	getBook, err := h.BookService.GetBook()
	if err != nil {
		http.Error(w, "Error al obtener los libros",
			http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getBook)

}
