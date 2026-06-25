package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"school_backend/internal/config"
	"school_backend/internal/models"
)

func GetStudents(c *gin.Context) {
	var students []models.Student
	query := config.DB.Model(&models.Student{})

	// Search by name or admission number
	if s := c.Query("search"); s != "" {
		query = query.Where("name ILIKE ? OR admission_no ILIKE ?", "%"+s+"%", "%"+s+"%")
	}

	// Filter by status (active / inactive)
	if status := c.Query("status"); status != "" && status != "All" {
		query = query.Where("status = ?", status)
	}

	// ✅ FIXED: Filter by class name (Nursery, LKG, UKG, Class 1 … Class 12)
	if class := c.Query("class"); class != "" && class != "All" {
		query = query.Where("class_name = ?", class)
	}

	query.Find(&students)
	c.JSON(http.StatusOK, gin.H{"data": students, "total": len(students)})
}

func GetStudentByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var student models.Student
	if err := config.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": student})
}

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	student.Status = "active"
	if err := config.DB.Create(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create student"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Student created", "data": student})
}

func UpdateStudent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var student models.Student
	if err := config.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.ShouldBindJSON(&student)
	config.DB.Save(&student)
	c.JSON(http.StatusOK, gin.H{"message": "Student updated", "data": student})
}

func DeleteStudent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	config.DB.Delete(&models.Student{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Student deleted"})
}