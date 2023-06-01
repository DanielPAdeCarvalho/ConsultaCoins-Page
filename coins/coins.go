package coins

import (
	"consultacoins/env"
	"consultacoins/models"
	"fmt"
	"html/template"
	"net/http"
)

func Saldo(w http.ResponseWriter, r *http.Request, email string) {
	t, err := template.ParseFiles("html/saldo.html")
	if err != nil {
		fmt.Println("Error parsing template do index.html:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	url := env.API_LOGON + "/"

	clientData := models.Client{
		Nome:  "Algodao Carvalho",
		Saldo: "P¢: 278",
	}
	t.Execute(w, clientData)
}
