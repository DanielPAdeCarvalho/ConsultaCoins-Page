package coins

import (
	"consultacoins/env"
	"consultacoins/models"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Saldo(w http.ResponseWriter, r *http.Request, email string) {
	template, err := template.ParseFiles("html/saldo.html")
	if err != nil {
		fmt.Println("Error parsing template do index.html:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	url := env.API_COINS + "/mail/" + email
	HttpClient := CertifyCoins()
	resp, err := HttpClient.Get(url)
	if err != nil {
		fmt.Println("Error getting saldo:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", body)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	parts := strings.Split(string(body), " ")
	if len(parts) < 3 {
		fmt.Println("Error parsing response body parts: ", parts)
		http.Error(w, "expected at least 3 parts, but got: ", http.StatusInternalServerError)
		return
	}

	saldinho, err := strconv.ParseFloat(parts[2][:len(parts[2])-1], 32)
	saldinho = math.Trunc(saldinho*100) / 100
	if err != nil {
		fmt.Println("Error parsing saldo to float:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	clientData := models.Client{
		Nome:  parts[0][1:] + " " + parts[1],
		Saldo: saldinho,
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
	HttpClient := CertifyCoins()
	response, err := HttpClient.Do(req)
	if err != nil {
		fmt.Println("Failed to do request iniciar carteira:", err)
		return
	}
	defer response.Body.Close()
}

func CertifyCoins() http.Client {
	//Certificate
	caPool := x509.NewCertPool()
	if ok := caPool.AppendCertsFromPEM([]byte(env.COINS_CERTIFICATE)); !ok {
		log.Fatalf("Failed to append certificate")
	}
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:    caPool,
			MinVersion: tls.VersionTLS13,
		},
	}
	client := &http.Client{
		Transport: transport,
		Timeout:   time.Second * 10, // Timeout after 10 seconds
	}
	return *client
}
