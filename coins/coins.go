package coins

import (
	"fmt"
	"html/template"
	"net/http"
)

type ClientData struct {
	Name  string
	Saldo string
}

func Saldo(w http.ResponseWriter, r *http.Request, email string) {
	t, err := template.ParseFiles("html/saldo.html")
	if err != nil {
		fmt.Println("Error parsing template do index.html:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	clientData := ClientData{
		Name:  "Algodao Carvalho",
		Saldo: "P¢: 278",
	}

	t.Execute(w, clientData)
}
