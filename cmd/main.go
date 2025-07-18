package main

import (
	"log"
	"net/http"

	handlers "github.com/t-saturn/auth-service-client/internal/hanlders"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handlers.LoginPage)
	// http.HandleFunc("/auth/google", handlers.GoogleLoginHandler)
	// http.HandleFunc("/auth/callback", handlers.GoogleCallbackHandler)

	log.Println("Servidor corriendo en http://localhost:5000")
	log.Fatal(http.ListenAndServe(":5080", nil))
}
