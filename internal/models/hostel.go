package models

import "time"

type HostelModel struct {
	ID               uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name             string    `gorm:"not null" json:"name"`
	Code             string    `gorm:"unique" json:"code"`
	Type             string    `json:"type"` // boys,girls,coed
	Address          string    `json:"address"`
	Floors           int       `json:"floors"`
	WardenName       string    `json:"warden_name"`
	Phone            string    `json:"phone"`
	Email            string    `json:"email"`
	EmergencyContact string    `json:"emergency_contact"`
	Status           string    `gorm:"default:active" json:"status"`
	CreatedAt        time.Time `json:"created_at"`
}

type Room struct {
	ID              uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	HostelID        uint      `gorm:"not null" json:"hostel_id"`
	RoomNumber      string    `gorm:"not null" json:"room_number"`
	Floor           int       `json:"floor"`
	Block           string    `json:"block"`
	RoomType        string    `json:"room_type"`
	Capacity        int       `json:"capacity"`
	Occupied        int       `gorm:"default:0" json:"occupied"`
	AttachedBath    bool      `json:"attached_bathroom"`
	IsAC            bool      `json:"is_ac"`
	IsFurnished     bool      `json:"is_furnished"`
	MonthlyRent     float64   `json:"monthly_rent"`
	SecurityDeposit float64   `json:"security_deposit"`
	Status          string    `gorm:"default:available" json:"status"`
	CreatedAt       time.Time `json:"created_at"`
}

type HostelStudent struct {
	ID              uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	StudentID       uint      `gorm:"not null" json:"student_id"`
	HostelID        uint      `json:"hostel_id"`
	RoomID          uint      `json:"room_id"`
	BedNumber       string    `json:"bed_number"`
	JoiningDate     string    `json:"joining_date"`
	ExpectedLeaving string    `json:"expected_leaving"`
	MonthlyFee      float64   `json:"monthly_fee"`
	Deposit         float64   `json:"deposit"`
	FeeStatus       string    `gorm:"default:pending" json:"fee_status"`
	Status          string    `gorm:"default:active" json:"status"`
	CreatedAt       time.Time `json:"created_at"`
}
type HostelAttendance struct {
	ID              uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	HostelStudentID uint      `gorm:"not null" json:"hostel_student_id"`
	Date            string    `gorm:"not null" json:"date"` // format: "2026-06-20"
	Status          string    `gorm:"not null" json:"status"` // present | absent | leave
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
type HostelComplaint struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	StudentID   uint      `gorm:"not null" json:"student_id"`
	HostelID    uint      `json:"hostel_id"`
	RoomNumber  string    `json:"room_number"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `json:"description"`
	Priority    string    `gorm:"default:medium" json:"priority"` // low | medium | high
	Status      string    `gorm:"default:pending" json:"status"`  // pending | assigned | resolved
	AssignedTo  string    `json:"assigned_to"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}


