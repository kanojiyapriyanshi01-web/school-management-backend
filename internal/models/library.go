package models

import "time"

type BookModel struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title          string    `gorm:"not null" json:"title"`
	Author         string    `json:"author"`
	ISBN           string    `gorm:"unique" json:"isbn"`
	Category       string    `json:"category"`
	Publisher      string    `json:"publisher"`
	PublishYear    int       `json:"publish_year"`
	TotalCopies    int       `gorm:"default:1" json:"total_copies"`
	AvailableCopies int      `gorm:"default:1" json:"available_copies"`
	Location       string    `json:"location"`
	Status         string    `gorm:"default:active" json:"status"`
	CreatedAt      time.Time `json:"created_at"`
}

type BookIssue struct {
	ID               uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	BookID           uint      `gorm:"not null" json:"book_id"`
	UserID           uint      `gorm:"not null" json:"user_id"`
	UserType         string    `json:"user_type"` // student,staff
	IssueDate        string    `json:"issue_date"`
	DueDate          string    `json:"due_date"`
	ReturnDate       string    `json:"return_date"`
	Status           string    `gorm:"default:issued" json:"status"`
	Fine             float64   `gorm:"default:0" json:"fine"`
	FinePaid         bool      `gorm:"default:false" json:"fine_paid"`
	CreatedAt        time.Time `json:"created_at"`
}

