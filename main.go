package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rodolfoalvesg/api-banking/api/common/config"
	"github.com/rodolfoalvesg/api-banking/api/gateways/http/router"
)

func main() {
	config.LoadEnv()
	r := router.CreateRouters()
	fmt.Printf("Escutando servidor %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
