package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.CreateRouters()

	fmt.Println("Escutando servidor")

	log.Fatal(http.ListenAndServe(":5000", r))
}
