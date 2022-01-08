package controllers

import "net/http"

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando conta"))
}

func ShowBallance(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Mostrando o saldo"))
}

func ShowAccounts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listando contas"))
}
