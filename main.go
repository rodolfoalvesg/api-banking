package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rodolfoalvesg/api-banking/api/common/config"
	"github.com/rodolfoalvesg/api-banking/api/controllers"
	account "github.com/rodolfoalvesg/api-banking/api/domain/usecases/accounts"
	"github.com/rodolfoalvesg/api-banking/api/gateways/db"
	"github.com/rodolfoalvesg/api-banking/api/gateways/http/router"
)

func main() {
	config.LoadEnv()

	repository := db.NewRepository()
	usecase := account.NewUsecase(repository)
	controller := controllers.NewController(usecase)
	r := router.CreateRouters(controller)

	fmt.Printf("Listening server %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
