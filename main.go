package main

import (
	"consultacoins/login"
	"crypto/tls"
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

	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Start server
	fmt.Println("Servidor iniciado na porta 8080")

	server := &http.Server{
		Addr:    ":8080",
		Handler: nil,
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS13,
		},
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Erro ao iniciar servidor: ", err)
	}
}
