package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rodolfoalvesg/api-banking/api/src/router"
)

func main() {
	r := router.CreateRouters()

	fmt.Println("Escutando servidor")

	log.Fatal(http.ListenAndServe(":5000", r))
}
