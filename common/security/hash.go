package security

import "golang.org/x/crypto/bcrypt"

// SecurityHash recebe uma string e coloca um hash nela
func SecurityHash(passwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
}

// VerifyPasswd faz a comparação do hash com string
func VerifyPasswd(passwdHash, passwdString string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwdHash), []byte(passwdString))
}
