package main

import (
	"context"
	"log"
	"os"

	"github.com/acaldo/rest-ws/handlers"
	"github.com/acaldo/rest-ws/middleware"
	"github.com/acaldo/rest-ws/server"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	s, err := server.NewServer(context.Background(), &server.Config{
		JWTSecret:   JWT_SECRET,
		Port:        PORT,
		DatabaseURL: DATABASE_URL,
	})

	if err != nil {
		log.Fatal(err)
	}

	s.Start(BindRoutes)
}

func BindRoutes(s server.Server, r *mux.Router) {
	r.Use(middleware.CheckAuthMiddleware(s))

	r.HandleFunc("/", handlers.HomeHandler(s)).Methods("GET")
	r.HandleFunc("/signup", handlers.SignUpHandler(s)).Methods("POST")
	r.HandleFunc("/login", handlers.LoginHandler(s)).Methods("POST")
	r.HandleFunc("/me", handlers.MeHandler(s)).Methods("GET")
	r.HandleFunc("/posts", handlers.InsertPostHandler(s)).Methods("POST")
	r.HandleFunc("/posts/{id}", handlers.GetPostByIdHandler(s)).Methods("GET")
	r.HandleFunc("/posts/{id}", handlers.UpdatePostHandler(s)).Methods("PUT")
	r.HandleFunc("/posts/{id}", handlers.DeletePostHandler(s)).Methods("DELETE")
	r.HandleFunc("/posts", handlers.ListPostHandler(s)).Methods("GET")

}
