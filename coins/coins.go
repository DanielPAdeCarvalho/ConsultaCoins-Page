package coins

import (
	"consultacoins/env"
	"consultacoins/models"
	"encoding/json"
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
		Saldo: 42,
	}
	template.Execute(w, clientData)
}

func StartWallet(client models.Client) {

	client.Saldo = 0
	jsonData, err := json.Marshal(client)
	if err != nil {
		fmt.Println("Failed to marshal JSON de logon na funcao startWallet:", err)
		return
	}
	url := env.API_COINS + "/newclient"
	req, err := http.NewRequest("POST", url, strings.NewReader(string(jsonData)))
	if err != nil {
		fmt.Println("Failed to create request iniciar carteira:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	clientAPI := &http.Client{}
	response, err := clientAPI.Do(req)
	if err != nil {
		fmt.Println("Failed to do request iniciar carteira:", err)
		return
	}
	defer response.Body.Close()
}
