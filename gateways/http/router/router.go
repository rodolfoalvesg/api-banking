package router

import (
	"github.com/rodolfoalvesg/api-banking/api/controllers"
	"github.com/rodolfoalvesg/api-banking/api/gateways/http/router/routers"

	"github.com/gorilla/mux"
)

//CreateRouters, cria as rotas
func CreateRouters(c *controllers.Controller) *mux.Router {
	r := mux.NewRouter()
	return routers.Setup(r, c)
}
