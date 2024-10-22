package ports

import "back-end-inventory/internal/domains"

type BookService interface {
	InsertBook(book domains.Books) (domains.Books, error)
	GetBook() ([]domains.Books, error)
}

type BookRepositorie interface {
	InsertBook(book domains.Books) (domains.Books, error)
	GetBook() ([]domains.Books, error)
}
