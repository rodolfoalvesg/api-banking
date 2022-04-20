package middlewares

import (
	"net/http"

	"github.com/rodolfoalvesg/api-banking/api/common/security"
	"github.com/rodolfoalvesg/api-banking/api/gateways/http/responses"
)

// Auth, verifica se o usuário ao fazer a requisição está autenticado
func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := security.ValidateToken(r); err != nil {
			responses.RespondError(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}
