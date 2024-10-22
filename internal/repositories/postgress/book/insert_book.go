package book

import "back-end-inventory/internal/domains"

func (r Repositorie) InsertBook(book domains.Books) (domains.Books, error) {
	query := `INSERT INTO books (name, author, pages, description, published, image) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id`

	var objectBook domains.Books

	err := r.DB.QueryRow(query, book.Name, book.Author, book.Pages, book.Description, book.Published, book.Image).Scan(objectBook)
	if err != nil {
		return domains.Books{}, nil
	}

	return objectBook, nil
}
