package models

import "time"

type Class struct {
    ID           uint      `gorm:"primaryKey" json:"id"`
    ClassName    string    `gorm:"not null" json:"class_name"`
    Section      string    `json:"section"`
    AcademicYear string    `json:"academic_year"`
    Stage        string    `json:"stage"`
    TeacherID    uint      `json:"teacher_id"`
    CreatedAt    time.Time `json:"created_at"`
}

type Subject struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    SubjectName string    `gorm:"not null" json:"subject_name"`
    SubjectCode string    `json:"subject_code"`
    ClassID     uint      `json:"class_id"`
    CreatedAt   time.Time `json:"created_at"`
}

type Attendance struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    StudentID uint      `gorm:"not null" json:"student_id"`
    ClassID   uint      `json:"class_id"`
    Date      string    `gorm:"not null" json:"date"`
    Status    string    `gorm:"not null" json:"status"`
    Period    string    `json:"period"`
    CreatedBy uint      `json:"created_by"`
    CreatedAt time.Time `json:"created_at"`
}

type Fee struct {
    ID         uint      `gorm:"primaryKey" json:"id"`
    StudentID  uint      `gorm:"not null" json:"student_id"`
    FeeType    string    `gorm:"not null" json:"fee_type"`
    Amount     float64   `gorm:"not null" json:"amount"`
    PaidAmount float64   `gorm:"default:0" json:"paid_amount"`
    DueDate    string    `json:"due_date"`
    PaidDate   string    `json:"paid_date"`
    Status     string    `gorm:"default:pending" json:"status"`
    CreatedAt  time.Time `json:"created_at"`
}

type Exam struct {
    ID           uint      `gorm:"primaryKey" json:"id"`
    ExamName     string    `gorm:"not null" json:"exam_name"`
    ClassID      uint      `json:"class_id"`
    ClassName    string    `json:"class_name"`
    Section      string    `gorm:"default:A" json:"section"`
    ExamType     string    `gorm:"default:'Unit Test'" json:"exam_type"`
    AcademicYear string    `gorm:"default:'2025-26'" json:"academic_year"`
    StartDate    string    `json:"start_date"`
    EndDate      string    `json:"end_date"`
    Description  string    `json:"description"`
    Status       string    `gorm:"default:draft" json:"status"`
    CreatedAt    time.Time `json:"created_at"`
}

type Mark struct {
    ID            uint      `gorm:"primaryKey" json:"id"`
    StudentID     uint      `gorm:"not null" json:"student_id"`
    ExamID        uint      `gorm:"not null" json:"exam_id"`
    SubjectID     uint      `json:"subject_id"`
    MarksObtained float64   `json:"marks_obtained"`
    MaxMarks      float64   `json:"max_marks"`
    Grade         string    `json:"grade"`
    CreatedAt     time.Time `json:"created_at"`
}

type Notice struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    Title       string    `gorm:"not null" json:"title"`
    Description string    `json:"description"`
    Date        string    `json:"date"`
    CreatedBy   uint      `json:"created_by"`
    TargetRole  string    `gorm:"default:all" json:"target_role"`
    CreatedAt   time.Time `json:"created_at"`
}

type Admission struct {
    ID            uint      `gorm:"primaryKey" json:"id"`
    StudentName   string    `gorm:"not null" json:"student_name"`
    DOB           string    `json:"dob"`
    Gender        string    `json:"gender"`
    ApplyingClass string    `json:"applying_class"`
    AcademicYear  string    `json:"academic_year"`
    FatherName    string    `json:"father_name"`
    MotherName    string    `json:"mother_name"`
    ParentPhone   string    `json:"parent_phone"`
    Email         string    `json:"email"`
    Address       string    `json:"address"`
    PreviousSchool string   `json:"previous_school"`
    Documents     string    `json:"documents"`
    Status        string    `gorm:"default:pending" json:"status"`
    StudentID     string    `json:"student_id"`
    Remarks       string    `json:"remarks"`
    CreatedAt     time.Time `json:"created_at"`
    UpdatedAt     time.Time `json:"updated_at"`
}