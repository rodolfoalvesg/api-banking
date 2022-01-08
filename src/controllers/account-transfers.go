package controllers

import "net/http"

func AmountTransfers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Transferencia entre contas"))
}

func ShowTransfers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listando transferências"))
}
