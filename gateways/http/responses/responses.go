package responses

import (
	"encoding/json"
	"net/http"
)

// RespondJSON, retorna uma resposta em JSON para a requisição
func RespondJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	response, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte(response))

}

// RespondError, retorna um erro em formato JSON
func RespondError(w http.ResponseWriter, statusCode int, err error) {
	RespondJSON(w, statusCode, map[string]string{
		"Error": err.Error(),
	})
}
