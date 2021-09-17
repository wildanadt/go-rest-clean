package server

import (
	"log"

	"github.com/wildanadt/go-rest-clean/config/db"
	// "github.com/wildanadt/go-rest-clean/config/db"
	"github.com/wildanadt/go-rest-clean/config/env"
	"github.com/wildanadt/go-rest-clean/controllers"
	"github.com/wildanadt/go-rest-clean/repository"
	r "github.com/wildanadt/go-rest-clean/router"
	usecases "github.com/wildanadt/go-rest-clean/usecase"
)

func init() {
	InitServer()
}

func InitServer() {
	router := r.NewRouter()
	api := router.Group("api/v1")
	db := db.GetDBInstance()

	userRepo := repository.NewUserRepo(db)
	userUsecase := usecases.NewUserHandler(userRepo)

	// controllers.NewUserHandler(api, userUsecase)
	controllers.NewAuthhandler(api, userUsecase)

	if err := router.Run(env.Config.Host + ":" + env.Config.Port); err != nil {
		log.Fatal(err, "main : run router")
	}

}
