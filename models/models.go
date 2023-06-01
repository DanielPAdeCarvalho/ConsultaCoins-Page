package models

type Client struct {
	Nome        string `json:"nome"`
	Sobrenome   string `json:"sobrenome"`
	Email       string `json:"email"`
	Senha       string `json:"senha"`
	Saldo       string `json:"saldo"`
	DataCriacao string `json:"data-criacao"`
}
