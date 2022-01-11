package security

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/rodolfoalvesg/api-banking/api/src/config"
)

//CreateToken, retorna um token assinado com as permissões de usuários
func CreateToken(userId int64) (string, error) {

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
