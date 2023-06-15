package models

type Client struct {
	Nome        string  `json:"nome"`
	Email       string  `json:"email"`
	Senha       string  `json:"senha"`
	Saldo       float64 `json:"saldo"`
	DataCriacao string  `json:"data-criacao"`
}
