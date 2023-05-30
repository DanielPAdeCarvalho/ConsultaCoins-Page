package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Servidor ok!")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login method:", r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("html/index.html")
		if err != nil {
			fmt.Println("Error parsing template:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("register method:", r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("html/register.html")
		if err != nil {
			fmt.Println("Error parsing template:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("username:", r.Form["fuser"])
		fmt.Println("email:", r.Form["femail"])
		fmt.Println("password:", r.Form["fpass"])
	}
}

func main() {
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)

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
