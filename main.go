package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rodolfoalvesg/api-banking/api/common/config"
	"github.com/rodolfoalvesg/api-banking/api/controllers"
	account "github.com/rodolfoalvesg/api-banking/api/domain/usecases/accounts"
	transfer "github.com/rodolfoalvesg/api-banking/api/domain/usecases/transfers"
	account_db "github.com/rodolfoalvesg/api-banking/api/gateways/db/account"
	transfer_db "github.com/rodolfoalvesg/api-banking/api/gateways/db/transfers"
	"github.com/rodolfoalvesg/api-banking/api/gateways/http/router"
)

func main() {
	config.LoadEnv()

	//Repositories
	repositoryAccount := account_db.NewRepository()
	repositoryTransfers := transfer_db.NewRepositoryTransfer()

	//Usecases
	usecaseAccount := account.NewUsecase(repositoryAccount)
	usecaseTransfer := transfer.NewUsecaseTransfers(repositoryTransfers)

	//Controllers
	controller := controllers.NewController(usecaseAccount, usecaseTransfer)

	//Router
	r := router.CreateRouters(controller)

	fmt.Printf("Listening server %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
