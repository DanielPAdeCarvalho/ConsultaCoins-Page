package login

import (
	"bytes"
	"consultacoins/env"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

func Login(w http.ResponseWriter, r *http.Request) {
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
		data := map[string]string{
			"email":    r.Form["email"][0],
			"password": r.Form["password"][0],
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Failed to marshal JSON de logon na funcao login:", err)
			return
		}

		url := env.API_LOGON + "/logonclient"
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			log.Fatalf("Failed to create request: %s", err)
		}
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		response, err := client.Do(req)
		if err != nil {
			log.Fatalf("Failed to do request: %s", err)
		}
		defer response.Body.Close()
		// Read the response body
		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal("Error reading response body:", err)
		}
		resposta := string(body)
		if strings.Contains(resposta, `"Authorized"`) {
			fmt.Println("Logado com sucesso!")
		} else {
			fmt.Println(resposta)
		}
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
