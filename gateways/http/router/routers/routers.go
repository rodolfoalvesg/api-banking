package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	URI            string
	Method         string
	Function       func(w http.ResponseWriter, r *http.Request)
	Authentication bool
}

func Setup(r *mux.Router) *mux.Router {

	routers := accountRouters
	routers = append(routers, loginRouter)
	routers = append(routers, transfersRouters...)

	for _, router := range routers {
		r.HandleFunc(router.URI, router.Function).Methods(router.Method)
	}

	return r
}