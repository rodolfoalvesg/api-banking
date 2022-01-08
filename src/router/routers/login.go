package routers

import (
	"api/src/controllers"
	"net/http"
)

var loginRouter = Router{
	URI:            "/login",
	Method:         http.MethodPost,
	Function:       controllers.Login,
	Authentication: false,
}
