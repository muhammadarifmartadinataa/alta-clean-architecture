package main

import (
	"log"
	"rest/config"
	authController "rest/controller/auth"
	"rest/middleware"
	authRepo "rest/repo/auth"
	"rest/route"
	authService "rest/service/auth"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	loadEnv()
	db, _ := config.ConnectDatabase()
	config.MigrateDB(db)
	e := echo.New()

	authJwt := middleware.JwtAlta{}

	authRepo := authRepo.NewAuthRepo(db)
	authService := authService.NewAuthService(authRepo, authJwt)
	authController := authController.NewAuthController(authService)

	routeController := route.RouteController{
		AuthController: authController,
	}
	routeController.InitRoute(e)

	e.Start(":8000")
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		panic("failed lod env")
	}
}
