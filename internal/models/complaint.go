package models

import "time"

type Complaint struct {
    ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
    StudentID   uint      `json:"student_id"`
    StudentName string    `json:"student_name"`
    RoomNumber  string    `json:"room_number"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Category    string    `json:"category"`
    Priority    string    `gorm:"default:medium" json:"priority"`
    Status      string    `gorm:"default:pending" json:"status"`
    AssignedTo  string    `json:"assigned_to"`
    AdminRemark string    `json:"admin_remark"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
