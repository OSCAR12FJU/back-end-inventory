package book

import (
	"back-end-inventory/internal/domains"
	"fmt"
)

func (r Repositorie) GetBook() ([]domains.Books, error) {
	query := `SELECT id ,name ,author ,image ,status, pages, published, description FROM books`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []domains.Books

	for rows.Next() {
		var book domains.Books
		err := rows.Scan(&book.ID, &book.Name, &book.Author, &book.Image, &book.Status, &book.Pages, &book.Published, &book.Description)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	fmt.Println(books)
	return books, nil
}
