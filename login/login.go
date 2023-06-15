package login

import (
	"bytes"
	"consultacoins/coins"
	"consultacoins/env"
	"consultacoins/models"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("html/index.html")
		if err != nil {
			fmt.Println("Error parsing template do index.html:", err)
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
			coins.Saldo(w, r, data["email"])
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
		nomecliente := r.Form["fnome"][0] + " " + r.Form["flast"][0]
		cliente := models.Client{
			Nome:  nomecliente,
			Email: r.Form["femail"][0],
			Senha: r.Form["fpass"][0],
		}
		jsonData, err := json.Marshal(cliente)
		if err != nil {
			fmt.Println("Failed to marshal JSON cadastro:", err)
			return
		}
		uri := env.API_LOGON + "/signclient"
		req, err := http.NewRequest("POST", uri, bytes.NewBuffer(jsonData))
		if err != nil {
			log.Fatalf("Failed to create cadastro: %s\n", err)
		}
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		response, err := client.Do(req)
		if err != nil {
			log.Fatalf("Failed to do request cadastro: %s\n", err)
		}
		coins.StartWallet(cliente)
		time.Sleep(2 * time.Second)
		defer response.Body.Close()
		//fando o login da pessoa na pagina
		loginReq, err := http.NewRequest("POST", "", nil)
		if err != nil {
			log.Fatalf("Failed to create login request: %s\n", err)
		}

		loginReq.PostForm = url.Values{
			"email":    {cliente.Email},
			"password": {cliente.Senha},
		}

		Login(w, loginReq)
	}
}
