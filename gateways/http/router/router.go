package router

import (
	"github.com/rodolfoalvesg/api-banking/api/gateways/http/router/routers"

	"github.com/gorilla/mux"
)

func CreateRouters() *mux.Router {
	r := mux.NewRouter()
	return routers.Setup(r)
}
