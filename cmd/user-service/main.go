package main

import (
	"github.com/gin-gonic/gin"
	"library-management-system/internal/shared/database"
	"library-management-system/internal/user/handler"
	"library-management-system/internal/user/repository"
	"library-management-system/internal/user/service"
	"library-management-system/pkg/logging"
)

func main() {
	logger := logging.Logger
	database.ConnectDB(logger)

	userRepo := repository.NewUserRepository(database.DB)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	router.POST("/register", userHandler.Register)
	router.POST("/login", userHandler.Login)

	router.Run(":8080")

}
