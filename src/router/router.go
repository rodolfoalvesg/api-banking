package router

import (
	"api/src/router/routers"

	"github.com/gorilla/mux"
)

func CreateRouters() *mux.Router {
	r := mux.NewRouter()
	return routers.Setup(r)
}
