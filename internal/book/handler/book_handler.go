package handler

import (
	"github.com/gin-gonic/gin"
	"library-management-system/internal/book/model"
	"library-management-system/internal/book/service"
	"net/http"
	"strconv"
)

type BookHandler struct {
	service *service.BookService
}

func NewBookHandler(service *service.BookService) *BookHandler {
	return &BookHandler{service: service}
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateBook(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": book})
}

func (h *BookHandler) GetBook(c *gin.Context) {

	bookId := c.Param("id")
	bookIdInt, err := strconv.Atoi(bookId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	book, err := h.service.GetBookByID(uint(bookIdInt))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func (h *BookHandler) GetAllBooks(c *gin.Context) {
	books, err := h.service.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve books"})
		return
	}

	if len(books) == 0 {
		c.JSON(http.StatusOK, gin.H{"data": []model.Book{}})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func (h *BookHandler) LoanBook(c *gin.Context) {

	var loanRequest struct {
		BookID  uint   `json:"book_id"`
		UserID  uint   `json:"user_id"`
		DueDate string `json:"due_date"`
	}

	if err := c.ShouldBindJSON(&loanRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := h.service.LoanBook(loanRequest.BookID, loanRequest.UserID, loanRequest.DueDate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": "Book Loaned"})
}

func (h *BookHandler) ReturnBook(c *gin.Context) {

	bookIdToReturn := c.Param("id")
	bookIdToReturnInt, err := strconv.Atoi(bookIdToReturn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = h.service.ReturnBook(uint(bookIdToReturnInt)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to return book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book returned successfully"})

}
