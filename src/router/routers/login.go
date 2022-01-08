package routers

import (
	"net/http"

	"github.com/rodolfoalvesg/api-banking/api/src/controllers"
)

var loginRouter = Router{
	URI:            "/login",
	Method:         http.MethodPost,
	Function:       controllers.Login,
	Authentication: false,
}
