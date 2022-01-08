package controllers

import "net/http"

func Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Rota de login"))
}
