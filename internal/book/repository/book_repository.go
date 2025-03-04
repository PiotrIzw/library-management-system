package repository

import (
	"gorm.io/gorm"
	"library-management-system/internal/book/model"
)

type BookRepository struct {
	DB *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{DB: db}
}

func (r *BookRepository) CreateBook(book *model.Book) error {
	return r.DB.Create(book).Error
}

func (r *BookRepository) FindBookByID(id uint) (*model.Book, error) {
	var book model.Book
	err := r.DB.Where("id = ?", id).Preload("Borrower", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name, email")
	}).First(&book).Error
	return &book, err
}

func (r *BookRepository) FindAll() ([]model.Book, error) {
	var books []model.Book
	err := r.DB.Preload("Borrower", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name, email")
	}).Find(&books).Error
	return books, err
}

func (r *BookRepository) UpdateBook(book *model.Book) error {
	return r.DB.Save(book).Error
}
