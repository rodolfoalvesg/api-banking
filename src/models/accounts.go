package models

import (
	"encoding/json"
	"time"
)

type Account struct {
	Id        string      `json:"id,omitempty"`
	Name      string      `json:"name,omitempty"`
	Cpf       string      `json:"cpf,omitempty"`
	Secret    string      `json:"secret,omitempty"`
	Balance   json.Number `json:"balance,omitempty"`
	CreatedAt time.Time   `json:"created_at,omitempty"`
}
