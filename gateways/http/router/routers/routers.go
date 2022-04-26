package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rodolfoalvesg/api-banking/api/controllers"
	"github.com/rodolfoalvesg/api-banking/api/gateways/http/middlewares"
)

type Router struct {
	URI            string
	Method         string
	Function       func(w http.ResponseWriter, r *http.Request)
	Authentication bool
}

//Setup, configuração de rotas
func Setup(r *mux.Router, c *controllers.Controller) *mux.Router {

	routerAccounts := RouterAccounts(*c)
	routerLogin := RouterLogin(*c)
	routerTransfers := RouterTransfers(*c)

	routers := routerAccounts
	routers = append(routers, routerLogin)
	routers = append(routers, routerTransfers...)

	for _, router := range routers {

		if router.Authentication {
			r.HandleFunc(router.URI, middlewares.Auth(router.Function)).Methods(router.Method)
		} else {
			r.HandleFunc(router.URI, router.Function).Methods(router.Method)
		}
	}

	return r
}
