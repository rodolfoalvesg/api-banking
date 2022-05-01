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
	permission["user_id"] = userId
	permission["exp"] = time.Now().Add(time.Minute * 30).Unix()

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

// ExtractUserID, retona o usuário que está salvo no token
func ExtractUserID(r *http.Request) (string, error) {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, returnVerificationKey)
	if err != nil {
		return "", err
	}

	if permission, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := permission["user_id"].(string)
		return userID, nil
	}

	return "", errors.New("Invalid Token")
}

//extractToken, desmenbra o token em caso de Bearer: Bearer token..........
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
