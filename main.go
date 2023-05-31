package main

import (
	"consultacoins/coins"
	"consultacoins/login"
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Servidor ok!")
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login.Login)
	http.HandleFunc("/register", login.Register)
	http.HandleFunc("/saldo", coins.Saldo)

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
