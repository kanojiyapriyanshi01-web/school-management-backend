package models

import "time"

type Student struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	UserID           uint      `json:"user_id"`
	AdmissionNo      string    `gorm:"unique;not null" json:"admission_no"`
	Name             string    `gorm:"not null" json:"name"`
	ClassID          uint      `json:"class_id"`
	ClassName        string    `json:"class_name"`
	Section          string    `json:"section"`
	RollNo           string    `json:"roll_no"`
	DOB              string    `json:"dob"`
	Gender           string    `json:"gender"`
	BloodGroup       string    `json:"blood_group"`
	Phone            string    `json:"phone"`
	Email            string    `json:"email"`
	Address          string    `json:"address"`
	FatherName       string    `json:"father_name"`
	MotherName       string    `json:"mother_name"`
	ParentPhone      string    `json:"parent_phone"`
	ParentEmail      string    `json:"parent_email"`
	ParentOccupation string    `json:"parent_occupation"`
	EmergencyContact string    `json:"emergency_contact"`
	MedicalInfo      string    `json:"medical_info"`
	Transport        string    `json:"transport"`
	BusRoute         string    `json:"bus_route"`
	AdmissionDate    string    `json:"admission_date"`
	Status           string    `gorm:"default:active" json:"status"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}