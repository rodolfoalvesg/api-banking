package routers

import (
	"net/http"
)

var loginRouter = Router{
	URI:            "/login",
	Method:         http.MethodPost,
	Authentication: false,
}
