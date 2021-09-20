package server

import (
	"log"

	"github.com/wildanadt/go-rest-clean/config/db"
	// "github.com/wildanadt/go-rest-clean/config/db"
	"github.com/wildanadt/go-rest-clean/config/env"
	"github.com/wildanadt/go-rest-clean/controllers"
	"github.com/wildanadt/go-rest-clean/middlewares"
	"github.com/wildanadt/go-rest-clean/repository"
	r "github.com/wildanadt/go-rest-clean/router"
	usecases "github.com/wildanadt/go-rest-clean/usecase"
)

func init() {
	InitServer()
}

func InitServer() {
	router := r.NewRouter()
	auth := router.Group("/")
	protected_api := router.Group("/api/v1")
	protected_api.Use(middlewares.NewAuthMiddleware())
	db := db.GetDBInstance()

	userRepo := repository.NewUserRepo(db)
	userUsecase := usecases.NewUserHandler(userRepo)
	authUsecase := usecases.NewAuthHanadler(userRepo)

	controllers.NewUserHandler(protected_api, userUsecase)
	controllers.NewAuthandler(auth, userUsecase, authUsecase)

	if err := router.Run(env.Config.Host + ":" + env.Config.Port); err != nil {
		log.Fatal(err, "main : run router")
	}

}
