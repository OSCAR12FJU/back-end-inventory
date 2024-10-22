package ports

import "back-end-inventory/internal/domains"

type UserService interface {
	InsertUser(user domains.Users) (domains.Users, error)
	// GetUser() ([]domains.Users, error)
	FindUserByEmail(email string) (*domains.Users, error)
}

type UserRepositorie interface {
	InsertUser(user domains.Users) (domains.Users, error)
	// GetUser() ([]domains.Users, error)
	FindUserByEmail(email string) (*domains.Users, error)
}
