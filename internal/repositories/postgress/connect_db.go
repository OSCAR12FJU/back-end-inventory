package postgress

import (
	"database/sql"
	"fmt"
)

func CreateConnection(connect string) (*sql.DB, error) {
	// err := godo

	db, err := sql.Open("postgres", connect)
	if err != nil {
		return nil, fmt.Errorf("error al conectar a Postgres: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error al conectar desde Ping: %w", err)
	}

	fmt.Println("Conectado a Postgres!")
	return db, nil
}
