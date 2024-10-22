package user

import (
	"back-end-inventory/internal/ports"
)

type Handler struct {
	UserService ports.UserService
}
