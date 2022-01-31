package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rodolfoalvesg/api-banking/api/src/config"
	"github.com/rodolfoalvesg/api-banking/api/src/db"
	"github.com/rodolfoalvesg/api-banking/api/src/router"
)

func main() {
	config.LoadEnv()
	r := router.CreateRouters()

	CreateAccount := db.FieldsToMethodsDB{}
	CreateAccount.AddedAccount()

	fmt.Printf("Escutando servidor %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
