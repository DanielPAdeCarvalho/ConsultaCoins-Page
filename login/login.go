package login

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println(os.Getenv("API_LOGIN"))
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

func Register(w http.ResponseWriter, r *http.Request) {
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
		fmt.Println("nome:", r.Form["fnome"])
		fmt.Println("sobrenome:", r.Form["flast"])
		fmt.Println("password:", r.Form["fpass"])
	}
}
