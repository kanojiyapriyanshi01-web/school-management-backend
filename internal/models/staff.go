package models

import "time"

type Staff struct {
	ID               uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID           uint      `json:"user_id"`
	EmployeeID       string    `gorm:"unique;not null" json:"employee_id"`
	Name             string    `gorm:"not null" json:"name"`
	Designation      string    `json:"designation"`
	Department       string    `json:"department"`
	Phone            string    `json:"phone"`
	Email            string    `json:"email"`
	Gender           string    `json:"gender"`
	DOB              string    `json:"dob"`
	JoiningDate      string    `json:"joining_date"`
	Qualification    string    `json:"qualification"`
	Experience       string    `json:"experience"`
	Address          string    `json:"address"`
	Role             string    `json:"role"`
	EmergencyContact string    `json:"emergency_contact"`
	PhotoURL         string    `json:"photo_url"`
	Status           string    `gorm:"default:active" json:"status"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

