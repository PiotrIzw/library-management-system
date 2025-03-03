package model

import (
	"gorm.io/gorm"
	"library-management-system/internal/user/model"
	"time"
)

type Book struct {
	gorm.Model
	Title      string `gorm:"not null"`
	Author     string `gorm:"not null"`
	Status     string `gorm:"not null"`
	LoanedAt   time.Time
	DueDate    time.Time
	ReturnedAt time.Time
	BorrowerID *uint
	Borrower   model.User `gorm:"foreignKey:BorrowerID"`
}
