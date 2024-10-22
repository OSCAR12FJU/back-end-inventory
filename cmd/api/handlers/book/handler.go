package book

import (
	"back-end-inventory/internal/ports"
)

type Handler struct {
	BookService ports.BookService
}
