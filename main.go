// main.go
package main

import (
	"auth-api/database"
	"auth-api/handlers"
	"auth-api/middleware"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	database.InitDB()
	defer database.DB.Close()

	r := mux.NewRouter()

	r.HandleFunc("/signup", handlers.SignUp).Methods("POST")
	r.HandleFunc("/signin", handlers.SignIn).Methods("POST")
	r.HandleFunc("/refresh", handlers.RefreshToken).Methods("POST")
	r.HandleFunc("/logout", handlers.Logout).Methods("POST")

	r.Handle("/protected", middleware.AuthMiddleware(http.HandlerFunc(handlers.ProtectedEndpoint))).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Printf("Server started on :%s\n", port)
	http.ListenAndServe(":"+port, r)
}
