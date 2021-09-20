package env

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/wildanadt/go-rest-clean/models"
)

var Config models.ServerConfig

func init() {
	err := LoadEnv()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func LoadEnv() (err error) {
	err = godotenv.Load(".env")
	if err != nil {
		fmt.Println(err.Error())
	}

	err = env.Parse(&Config)
	if err != nil {
		return err
	}
	err = env.Parse(&Config.PostgresConfig)
	if err != nil {
		return err
	}
	err = env.Parse(&Config.JWTConfig)
	if err != nil {
		return err
	}

	return err

}
