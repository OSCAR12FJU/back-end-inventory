package book

import "back-end-inventory/internal/domains"

func (s Services) InsertBook(book domains.Books) (domains.Books, error) {
	return s.Repo.InsertBook(book)
}
