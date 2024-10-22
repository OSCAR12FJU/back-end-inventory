package user

import (
	"back-end-inventory/internal/domains"
	"database/sql"
	"errors"
)

func (r Repositorie) FindUserByEmail(email string) (*domains.Users, error) {

	var objectUser domains.Users
	query := `SELECT id, name, last_name , email, password, age, nacionality, image FROM users WHERE email = $1`

	row := r.DB.QueryRow(query, email)
	err := row.Scan(&objectUser.ID, &objectUser.Name, &objectUser.LastName, &objectUser.Email, &objectUser.Password, &objectUser.Age, &objectUser.Nacionality, &objectUser.Image)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("usuario no encontrado")
		}
		return nil, err
	}

	return &objectUser, nil

}
