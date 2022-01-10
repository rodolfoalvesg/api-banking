package models

import (
	"regexp"
	"time"
)

type Account struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Cpf       string    `json:"cpf"`
	Secret    string    `json:"secret"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// FormatDocumentNumber remove espações e caracteres não alphanuméricos
func (a *Account) FormatDocumentNumber() {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	a.Cpf = reg.ReplaceAllString(a.Cpf, "")
}
