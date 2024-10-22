package book

import "back-end-inventory/internal/domains"

func (s Services) GetBook() ([]domains.Books, error) {
	return s.Repo.GetBook()
}
