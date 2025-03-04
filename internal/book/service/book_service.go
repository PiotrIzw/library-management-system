package service

import (
	"errors"
	"library-management-system/internal/book/model"
	"library-management-system/internal/book/repository"
	"time"
)

type BookService struct {
	repo *repository.BookRepository
}

func NewBookService(repo *repository.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) CreateBook(book *model.Book) error {
	return s.repo.CreateBook(book)
}

func (s *BookService) GetBookByID(id uint) (*model.Book, error) {
	return s.repo.FindBookByID(id)
}

func (s *BookService) GetAllBooks() ([]model.Book, error) {
	return s.repo.FindAll()
}

func (s *BookService) LoanBook(bookID, userID uint, dueDate string) error {
	book, err := s.repo.FindBookByID(bookID)
	if err != nil {
		return err
	}

	if book.Status != "available" {
		return errors.New("book is not available for loan")
	}

	parsedDueDate, err := time.Parse("2006-01-02", dueDate)
	if err != nil {
		return errors.New("invalid due date format")
	}

	book.Status = "loaned"
	book.BorrowerID = &userID
	book.LoanedAt = time.Now()
	book.DueDate = parsedDueDate

	return s.repo.UpdateBook(book)
}

func (s *BookService) ReturnBook(bookID uint) error {
	book, err := s.repo.FindBookByID(bookID)
	if err != nil {
		return err
	}
	if book.Status != "loaned" {
		return errors.New("book is not currently loaned")
	}

	book.Status = "available"
	book.ReturnedAt = time.Now()
	book.BorrowerID = nil

	return s.repo.UpdateBook(book)
}
