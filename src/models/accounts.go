package models

import (
	"encoding/json"
	"regexp"
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

// FormatDocumentNumber remove espações e caracteres não alphanuméricos
func (a *Account) FormatDocumentNumber() {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	a.Cpf = reg.ReplaceAllString(a.Cpf, "")
}
