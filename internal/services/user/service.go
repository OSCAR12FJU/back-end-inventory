package user

import "back-end-inventory/internal/ports"

type Services struct {
	Repo ports.UserRepositorie
	// TokenGenerator ports.TokenGenerator
}
