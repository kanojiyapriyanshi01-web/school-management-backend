package controllers

import (
    "fmt"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "school_backend/internal/config"
    "school_backend/internal/models"
)

func GetStaff(c *gin.Context) {
    var staff []models.Staff
    config.DB.Find(&staff)
    c.JSON(http.StatusOK, gin.H{"data": staff, "total": len(staff)})
}

func CreateStaff(c *gin.Context) {
    var staff models.Staff
    if err := c.ShouldBindJSON(&staff); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    staff.Status = "active"

    // Auto generate unique Employee ID
    if staff.EmployeeID == "" {
        var count int64
        config.DB.Model(&models.Staff{}).Count(&count)
        for {
            staff.EmployeeID = fmt.Sprintf("EMP%04d", count+1)
            var existing models.Staff
            if err := config.DB.Where("employee_id = ?", staff.EmployeeID).First(&existing).Error; err != nil {
                break
            }
            count++
        }
    }

    config.DB.Create(&staff)
    c.JSON(http.StatusCreated, gin.H{"message": "Staff created", "data": staff})
}

func UpdateStaff(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var staff models.Staff
    config.DB.First(&staff, id)
    c.ShouldBindJSON(&staff)
    config.DB.Save(&staff)
    c.JSON(http.StatusOK, gin.H{"message": "Staff updated", "data": staff})
}

func DeleteStaff(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    config.DB.Delete(&models.Staff{}, id)
    c.JSON(http.StatusOK, gin.H{"message": "Staff deleted"})
}

func GetAttendance(c *gin.Context) {
    var records []models.Attendance
    query := config.DB.Model(&models.Attendance{})
    if date := c.Query("date"); date != "" {
        query = query.Where("date = ?", date)
    }
    if classID := c.Query("class_id"); classID != "" {
        query = query.Where("class_id = ?", classID)
    }
    if studentID := c.Query("student_id"); studentID != "" {
        query = query.Where("student_id = ?", studentID)
    }
    query.Find(&records)
    c.JSON(http.StatusOK, gin.H{"data": records})
}

func MarkAttendance(c *gin.Context) {
    var records []models.Attendance
    if err := c.ShouldBindJSON(&records); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    for i := range records {
        config.DB.Where(models.Attendance{
            StudentID: records[i].StudentID,
            Date:      records[i].Date,
        }).Assign(records[i]).FirstOrCreate(&records[i])
    }
    c.JSON(http.StatusOK, gin.H{"message": "Attendance marked"})
}

func GetFees(c *gin.Context) {
    var fees []models.Fee
    query := config.DB.Model(&models.Fee{})
    if sid := c.Query("student_id"); sid != "" {
        query = query.Where("student_id = ?", sid)
    }
    if status := c.Query("status"); status != "" {
        query = query.Where("status = ?", status)
    }
    query.Find(&fees)

    // Add student names
    type FeeWithStudent struct {
        models.Fee
        StudentName string `json:"student_name"`
    }
    var result []FeeWithStudent
    for _, fee := range fees {
        var student models.Student
        name := "Unknown"
        if err := config.DB.First(&student, fee.StudentID).Error; err == nil {
            name = student.Name
        }
        result = append(result, FeeWithStudent{Fee: fee, StudentName: name})
    }
    if result == nil {
        result = []FeeWithStudent{}
    }
    c.JSON(http.StatusOK, gin.H{"data": result, "total": len(result)})
}

func CreateFee(c *gin.Context) {
    var fee models.Fee
    if err := c.ShouldBindJSON(&fee); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Create(&fee)
    c.JSON(http.StatusCreated, gin.H{"message": "Fee created", "data": fee})
}

func UpdateFee(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var fee models.Fee
    config.DB.First(&fee, id)
    c.ShouldBindJSON(&fee)
    config.DB.Save(&fee)
    c.JSON(http.StatusOK, gin.H{"message": "Fee updated", "data": fee})
}

func GetNotices(c *gin.Context) {
    var notices []models.Notice
    config.DB.Order("created_at desc").Find(&notices)
    c.JSON(http.StatusOK, gin.H{"data": notices})
}

func CreateNotice(c *gin.Context) {
    var notice models.Notice
    if err := c.ShouldBindJSON(&notice); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    notice.CreatedBy = c.GetUint("user_id")
    config.DB.Create(&notice)
    c.JSON(http.StatusCreated, gin.H{"message": "Notice created", "data": notice})
}

func DeleteNotice(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    config.DB.Delete(&models.Notice{}, id)
    c.JSON(http.StatusOK, gin.H{"message": "Notice deleted"})
}

func GetExams(c *gin.Context) {
    var exams []models.Exam
    config.DB.Find(&exams)
    c.JSON(http.StatusOK, gin.H{"data": exams})
}

func UpdateExam(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var exam models.Exam
    if err := config.DB.First(&exam, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Exam not found"})
        return
    }
    c.ShouldBindJSON(&exam)
    config.DB.Save(&exam)
    c.JSON(http.StatusOK, gin.H{"message": "Exam updated", "data": exam})
}
func CreateExam(c *gin.Context) {
    var exam models.Exam
    c.ShouldBindJSON(&exam)
    config.DB.Create(&exam)
    c.JSON(http.StatusCreated, gin.H{"message": "Exam created", "data": exam})
}

func GetMarks(c *gin.Context) {
    var marks []models.Mark
    query := config.DB.Model(&models.Mark{})
    if sid := c.Query("student_id"); sid != "" {
        query = query.Where("student_id = ?", sid)
    }
    if eid := c.Query("exam_id"); eid != "" {
        query = query.Where("exam_id = ?", eid)
    }
    query.Find(&marks)
    c.JSON(http.StatusOK, gin.H{"data": marks})
}

func SaveMarks(c *gin.Context) {
    var marks []models.Mark
    c.ShouldBindJSON(&marks)
    for i := range marks {
        config.DB.Where(models.Mark{
            StudentID: marks[i].StudentID,
            ExamID:    marks[i].ExamID,
            SubjectID: marks[i].SubjectID,
        }).Assign(marks[i]).FirstOrCreate(&marks[i])
    }
    c.JSON(http.StatusOK, gin.H{"message": "Marks saved"})
}

func GetBooks(c *gin.Context) {
    var books []models.BookModel
    query := config.DB.Model(&models.BookModel{})
    if s := c.Query("search"); s != "" {
        query = query.Where("title ILIKE ? OR author ILIKE ?", "%"+s+"%", "%"+s+"%")
    }
    query.Find(&books)
    c.JSON(http.StatusOK, gin.H{"data": books, "total": len(books)})
}

func CreateBook(c *gin.Context) {
    var book models.BookModel
    c.ShouldBindJSON(&book)
    config.DB.Create(&book)
    c.JSON(http.StatusCreated, gin.H{"message": "Book added", "data": book})
}

func IssueBook(c *gin.Context) {
    var issue models.BookIssue
    if err := c.ShouldBindJSON(&issue); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    issue.Status = "issued"
    config.DB.Create(&issue)
    config.DB.Model(&models.BookModel{}).Where("id = ?", issue.BookID).
        UpdateColumn("available_copies", gorm.Expr("available_copies - 1"))
    c.JSON(http.StatusCreated, gin.H{"message": "Book issued", "data": issue})
}

func ReturnBook(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var issue models.BookIssue
    config.DB.First(&issue, id)
    c.ShouldBindJSON(&issue)
    issue.Status = "returned"
    config.DB.Save(&issue)
    config.DB.Model(&models.BookModel{}).Where("id = ?", issue.BookID).
        UpdateColumn("available_copies", gorm.Expr("available_copies + 1"))
    c.JSON(http.StatusOK, gin.H{"message": "Book returned", "data": issue})
}

func GetMyBooks(c *gin.Context) {
    userID := c.GetUint("user_id")
    role := c.GetString("role")
    var issues []models.BookIssue
    if role == "admin" || role == "staff" {
        config.DB.Find(&issues)
    } else {
        config.DB.Where("user_id = ?", userID).Find(&issues)
    }
    c.JSON(http.StatusOK, gin.H{"data": issues})
}


