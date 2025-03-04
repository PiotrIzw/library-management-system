package main

import (
	"github.com/gin-gonic/gin"
	"library-management-system/internal/book/handler"
	"library-management-system/internal/book/repository"
	"library-management-system/internal/book/service"
	"library-management-system/internal/shared/database"
	"library-management-system/internal/user/middleware"
	"library-management-system/pkg/logging"
)

func main() {
	logger := logging.Logger
	database.ConnectDB(logger)

	bookRepo := repository.NewBookRepository(database.DB)
	bookService := service.NewBookService(bookRepo)
	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()

	router.Use(middleware.AuthMiddleware(logger))
	router.POST("/books", middleware.RoleBasedAuthMiddleware("librarian", logger), bookHandler.CreateBook)
	router.GET("/books/:id", bookHandler.GetBook)
	router.GET("/books", bookHandler.GetAllBooks)
	router.POST("/books/loan", bookHandler.LoanBook)
	router.PUT("/books/:id/return", bookHandler.ReturnBook)

	router.Run(":8081")

}
