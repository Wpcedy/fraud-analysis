package config

import (
	"log"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis do sistema.")
	} else {
		log.Println("Arquivo .env carregado com sucesso.")
	}
}