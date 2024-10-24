package main

import (
	bookHandler "back-end-inventory/cmd/api/handlers/book"
	"back-end-inventory/internal/repositories/postgress"
	bookRepo "back-end-inventory/internal/repositories/postgress/book"
	bookService "back-end-inventory/internal/services/book"
	"net/http"

	userHandler "back-end-inventory/cmd/api/handlers/user"
	userRepo "back-end-inventory/internal/repositories/postgress/user"
	userService "back-end-inventory/internal/services/user"

	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No archivo .env")
	}
	routers := mux.NewRouter()

	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	portAsNumber, _ := strconv.Atoi(dbPort)

	connectKey := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", portAsNumber, dbUser, dbPassword, dbName)
	db, err := postgress.CreateConnection(connectKey)
	if err != nil {
		log.Fatal(err.Error())
	}

	bookRepo := bookRepo.Repositorie{
		DB: db,
	}
	bookSrv := bookService.Services{
		Repo: bookRepo,
	}
	bookHandler := bookHandler.Handler{
		BookService: bookSrv,
	}

	// tokenGenerator := util.NewTokenUtil("my_secret_key")

	userRepo := userRepo.Repositorie{
		DB: db,
	}
	userSrv := userService.Services{
		Repo: userRepo,
		// TokenGenerator: tokenGenerator,
	}
	userHandler := userHandler.Handler{
		UserService: userSrv,
	}

	routers.HandleFunc("/create-user", userHandler.CreateUSer)
	routers.HandleFunc("/login-user", userHandler.LoginUser)

	routers.HandleFunc("/create-book", bookHandler.CreateBook)
	routers.HandleFunc("/get-book", bookHandler.GetAllBook)

	routers.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads/"))))

	handler := cors.Default().Handler(routers)
	log.Println("Srv corriendo en el puerto 8081")
	log.Fatal(http.ListenAndServe(":8081", handler))

}
