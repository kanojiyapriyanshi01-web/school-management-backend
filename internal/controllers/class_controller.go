package controllers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "school_backend/internal/config"
    "school_backend/internal/models"
)

// Education stage auto-assign
func getStage(className string) string {
    stages := map[string]string{
        "Nursery": "Foundational", "LKG": "Foundational", "UKG": "Foundational",
        "Class 1": "Foundational", "Class 2": "Foundational",
        "Class 3": "Preparatory", "Class 4": "Preparatory", "Class 5": "Preparatory",
        "Class 6": "Middle", "Class 7": "Middle", "Class 8": "Middle",
        "Class 9": "Secondary", "Class 10": "Secondary",
        "Class 11": "Secondary", "Class 12": "Secondary",
    }
    if stage, ok := stages[className]; ok {
        return stage
    }
    return "Foundational"
}

func GetClasses(c *gin.Context) {
    var classes []models.Class
    query := config.DB.Model(&models.Class{})
    if ay := c.Query("academic_year"); ay != "" {
        query = query.Where("academic_year = ?", ay)
    }
    query.Find(&classes)
    c.JSON(http.StatusOK, gin.H{"data": classes, "total": len(classes)})
}

func GetClassByID(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var class models.Class
    if err := config.DB.First(&class, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Class not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": class})
}

func CreateClass(c *gin.Context) {
    var class models.Class
    if err := c.ShouldBindJSON(&class); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    class.Stage = getStage(class.ClassName)
    if err := config.DB.Create(&class).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create class"})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"message": "Class created", "data": class})
}

func UpdateClass(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var class models.Class
    if err := config.DB.First(&class, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Class not found"})
        return
    }
    c.ShouldBindJSON(&class)
    class.Stage = getStage(class.ClassName)
    config.DB.Save(&class)
    c.JSON(http.StatusOK, gin.H{"message": "Class updated", "data": class})
}

func DeleteClass(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    config.DB.Delete(&models.Class{}, id)
    c.JSON(http.StatusOK, gin.H{"message": "Class deleted"})
}

func SeedClasses(c *gin.Context) {
    classes := []models.Class{
        {ClassName: "Nursery",  Section: "A", AcademicYear: "2025-26", Stage: "Foundational"},
        {ClassName: "LKG",     Section: "A", AcademicYear: "2025-26", Stage: "Foundational"},
        {ClassName: "UKG",     Section: "A", AcademicYear: "2025-26", Stage: "Foundational"},
        {ClassName: "Class 1", Section: "A", AcademicYear: "2025-26", Stage: "Foundational"},
        {ClassName: "Class 2", Section: "A", AcademicYear: "2025-26", Stage: "Foundational"},
        {ClassName: "Class 3", Section: "A", AcademicYear: "2025-26", Stage: "Preparatory"},
        {ClassName: "Class 4", Section: "A", AcademicYear: "2025-26", Stage: "Preparatory"},
        {ClassName: "Class 5", Section: "A", AcademicYear: "2025-26", Stage: "Preparatory"},
        {ClassName: "Class 6", Section: "A", AcademicYear: "2025-26", Stage: "Middle"},
        {ClassName: "Class 7", Section: "A", AcademicYear: "2025-26", Stage: "Middle"},
        {ClassName: "Class 8", Section: "A", AcademicYear: "2025-26", Stage: "Middle"},
        {ClassName: "Class 9",  Section: "A", AcademicYear: "2025-26", Stage: "Secondary"},
        {ClassName: "Class 10", Section: "A", AcademicYear: "2025-26", Stage: "Secondary"},
        {ClassName: "Class 11", Section: "A", AcademicYear: "2025-26", Stage: "Secondary"},
        {ClassName: "Class 12", Section: "A", AcademicYear: "2025-26", Stage: "Secondary"},
    }
    for _, cl := range classes {
        config.DB.FirstOrCreate(&cl, models.Class{ClassName: cl.ClassName, Section: cl.Section, AcademicYear: cl.AcademicYear})
    }
    c.JSON(http.StatusOK, gin.H{"message": "Classes seeded", "total": len(classes)})
}

// ── Admission ─────────────────────────────────────────────
func GetAdmissions(c *gin.Context) {
    var admissions []models.Admission
    query := config.DB.Model(&models.Admission{})
    if status := c.Query("status"); status != "" {
        query = query.Where("status = ?", status)
    }
    if search := c.Query("search"); search != "" {
        query = query.Where("student_name ILIKE ? OR email ILIKE ?", "%"+search+"%", "%"+search+"%")
    }
    query.Order("created_at desc").Find(&admissions)
    c.JSON(http.StatusOK, gin.H{"data": admissions, "total": len(admissions)})
}

func GetAdmissionByID(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var admission models.Admission
    if err := config.DB.First(&admission, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Admission not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": admission})
}

func CreateAdmission(c *gin.Context) {
    var admission models.Admission
    if err := c.ShouldBindJSON(&admission); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    admission.Status = "pending"
    if err := config.DB.Create(&admission).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to submit admission"})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"message": "Admission submitted successfully", "data": admission})
}

func UpdateAdmissionStatus(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var req struct {
        Status  string `json:"status"`
        Remarks string `json:"remarks"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    var admission models.Admission
    if err := config.DB.First(&admission, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Admission not found"})
        return
    }
    admission.Status = req.Status
    admission.Remarks = req.Remarks
    if req.Status == "approved" {
        admission.StudentID = generateStudentID()
    }
    config.DB.Save(&admission)
    c.JSON(http.StatusOK, gin.H{"message": "Status updated", "data": admission})
}

func generateStudentID() string {
    var count int64
    config.DB.Model(&models.Admission{}).Where("status = ?", "approved").Count(&count)
    return "STU" + strconv.Itoa(int(count+1001))
}

func DeleteAdmission(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    config.DB.Delete(&models.Admission{}, id)
    c.JSON(http.StatusOK, gin.H{"message": "Admission deleted"})
}
