package coins

import (
	"consultacoins/env"
	"consultacoins/models"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strings"
)

func Saldo(w http.ResponseWriter, r *http.Request, email string) {
	template, err := template.ParseFiles("html/saldo.html")
	if err != nil {
		fmt.Println("Error parsing template do index.html:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	url := env.API_COINS + "/mail/" + email
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error getting saldo:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	parts := strings.Split(string(body), " ")
	if len(parts) < 3 {
		http.Error(w, "expected at least 3 parts, but got: ", http.StatusInternalServerError)
		return
	}

	clientData := models.Client{
		Nome:  parts[0][1:] + " " + parts[1],
		Saldo: "P¢: " + parts[2][:len(parts[2])-1],
	}
	template.Execute(w, clientData)
}
