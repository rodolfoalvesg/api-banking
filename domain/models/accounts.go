package models

import (
	"time"
)

type AccountId string

type Account struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Cpf       string    `json:"cpf"`
	Secret    string    `json:"secret"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
