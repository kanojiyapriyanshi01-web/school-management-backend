package config
import (
    "fmt"
    "log"
    "os"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "school_backend/internal/models"
)
var DB *gorm.DB
func ConnectDB() {
    dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Kolkata",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
    )
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect database:", err)
    }
    DB = db
    log.Println("Database connected!")
}
func MigrateDB() {
    err := DB.AutoMigrate(
        &models.User{},
        &models.Student{},
        &models.Staff{},
        &models.Class{},
        &models.Subject{},
        &models.Attendance{},
        &models.Fee{},
        &models.Exam{},
        &models.Mark{},
        &models.Notice{},
        &models.Admission{},
        &models.HostelModel{},
        &models.Room{},
        &models.HostelStudent{},
        &models.HostelAttendance{},
        &models.HostelComplaint{},
        &models.Complaint{},
        &models.BookModel{},
        &models.BookIssue{},
        &models.Vehicle{},
        &models.TransportRoute{},
        &models.StudentTransport{},
        &models.TransportFee{},
    )
    if err != nil {
        log.Fatal("Migration failed:", err)
    }
    log.Println("Database migrated!")
}
