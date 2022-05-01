package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// Porta de funcionamento da APi
	Port = 0

	//SecretKey chave de assinatura de token
	SecretKey []byte
)

//LoadEnv, carregamento das variáveis de ambiente
func LoadEnv() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 5050
	}

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
