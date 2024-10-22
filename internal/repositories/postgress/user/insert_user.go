package user

import "back-end-inventory/internal/domains"

func (r Repositorie) InsertUser(user domains.Users) (domains.Users, error) {
	query := `INSERT INTO users (name, last_name, email, password, age, nacionality, image) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id`

	var objectUser domains.Users

	err := r.DB.QueryRow(query, user.Name, user.LastName, user.Email, user.Password, user.Age, user.Nacionality, user.Image).Scan(objectUser)

	if err != nil {
		return domains.Users{}, nil
	}

	return objectUser, nil
}
