package main

import (
	"consultacoins/login"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Servidor ok!")
}

func LoadEnv() {
	err := godotenv.Load("env/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}

func main() {
	LoadEnv()
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login.Login)
	http.HandleFunc("/register", login.Register)

	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Start server
	fmt.Println("Servidor iniciado na porta 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erro ao iniciar servidor: ", err)
	}
}
