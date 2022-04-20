package security

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/rodolfoalvesg/api-banking/api/common/config"
)

//CreateToken, retorna um token assinado com as permissões de usuários
func CreateToken(userId string) (string, error) {

	permission := jwt.MapClaims{}
	permission["authorized"] = true
	permission["id"] = userId
	permission["exp"] = time.Now().Add(time.Minute * 5).Unix()

	tokenMethod := jwt.GetSigningMethod("HS256")
	newTokenClaims := jwt.NewWithClaims(tokenMethod, permission)
	token, err := newTokenClaims.SignedString(config.SecretKey)

	if err != nil {
		return "", err
	}

	return token, nil
}

// ValidateToken, verifica se o token passado na requisição é válido
func ValidateToken(r *http.Request) error {

	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, returnVerificationKey)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Invalid token")
}

//extractToken, desmenbra o token
func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

//returnVerificationKey, verificação do metódo de assinatura
func returnVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected subscription method! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}
