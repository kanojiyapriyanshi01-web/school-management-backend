package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"school_backend/internal/config"
	"school_backend/internal/models"
)

// ============================================================
// HOSTEL (Hostel Master)
// ============================================================

func GetHostels(c *gin.Context) {
	var hostels []models.HostelModel
	config.DB.Find(&hostels)
	c.JSON(http.StatusOK, gin.H{"data": hostels, "total": len(hostels)})
}

func GetHostelByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var hostel models.HostelModel
	if err := config.DB.First(&hostel, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hostel not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": hostel})
}

func CreateHostel(c *gin.Context) {
	var hostel models.HostelModel
	if err := c.ShouldBindJSON(&hostel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if hostel.Status == "" {
		hostel.Status = "active"
	}
	if err := config.DB.Create(&hostel).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create hostel"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Hostel created", "data": hostel})
}

func UpdateHostel(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var hostel models.HostelModel
	if err := config.DB.First(&hostel, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hostel not found"})
		return
	}
	c.ShouldBindJSON(&hostel)
	config.DB.Save(&hostel)
	c.JSON(http.StatusOK, gin.H{"message": "Hostel updated", "data": hostel})
}

func DeleteHostel(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	config.DB.Delete(&models.HostelModel{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Hostel deleted"})
}

// ============================================================
// ROOMS
// ============================================================

func GetRooms(c *gin.Context) {
	var rooms []models.Room
	query := config.DB.Model(&models.Room{})
	if hostelID := c.Query("hostel_id"); hostelID != "" {
		query = query.Where("hostel_id = ?", hostelID)
	}
	if status := c.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}
	query.Find(&rooms)
	c.JSON(http.StatusOK, gin.H{"data": rooms, "total": len(rooms)})
}

func GetRoomByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var room models.Room
	if err := config.DB.First(&room, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": room})
}

func CreateRoom(c *gin.Context) {
	var room models.Room
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if room.Status == "" {
		room.Status = "available"
	}
	if err := config.DB.Create(&room).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create room"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Room created", "data": room})
}

func UpdateRoom(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var room models.Room
	if err := config.DB.First(&room, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
		return
	}
	c.ShouldBindJSON(&room)
	config.DB.Save(&room)
	c.JSON(http.StatusOK, gin.H{"message": "Room updated", "data": room})
}

func DeleteRoom(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	config.DB.Delete(&models.Room{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Room deleted"})
}

// ============================================================
// HOSTEL STUDENTS (Allocation / Students tab / Checkout / Transfer)
// ============================================================

// GetHostelStudents -> powers the "Students" tab (Image 1) and is the
// base data source the Fees and Attendance tabs also read from.
func GetHostelStudents(c *gin.Context) {
	var hostelStudents []models.HostelStudent
	query := config.DB.Model(&models.HostelStudent{})
	if hostelID := c.Query("hostel_id"); hostelID != "" {
		query = query.Where("hostel_id = ?", hostelID)
	}
	if status := c.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}
	if feeStatus := c.Query("fee_status"); feeStatus != "" {
		query = query.Where("fee_status = ?", feeStatus)
	}
	query.Find(&hostelStudents)
	c.JSON(http.StatusOK, gin.H{"data": hostelStudents, "total": len(hostelStudents)})
}

func GetHostelStudentByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var hs models.HostelStudent
	if err := config.DB.First(&hs, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hostel student record not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": hs})
}

// AllocateRoom -> powers the "Allocate Room" button (Image 1).
// Creates a HostelStudent record linking an existing Student to a Room,
// and bumps the room's Occupied count.
func AllocateRoom(c *gin.Context) {
	var hs models.HostelStudent
	if err := c.ShouldBindJSON(&hs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if hs.Status == "" {
		hs.Status = "active"
	}
	if hs.FeeStatus == "" {
		hs.FeeStatus = "pending"
	}

	if err := config.DB.Create(&hs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to allocate room"})
		return
	}

	// increment room occupied count
	config.DB.Model(&models.Room{}).Where("id = ?", hs.RoomID).
		UpdateColumn("occupied", gorm.Expr("occupied + 1"))

	c.JSON(http.StatusCreated, gin.H{"message": "Room allocated", "data": hs})
}

// TransferRoom -> powers the "Transfer" button. Moves a hostel student to a
// different room (and optionally a different hostel).
func TransferRoom(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var hs models.HostelStudent
	if err := config.DB.First(&hs, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hostel student record not found"})
		return
	}

	var body struct {
		HostelID uint   `json:"hostel_id"`
		RoomID   uint   `json:"room_id"`
		BedNumber string `json:"bed_number"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	oldRoomID := hs.RoomID

	hs.HostelID = body.HostelID
	hs.RoomID = body.RoomID
	hs.BedNumber = body.BedNumber
	config.DB.Save(&hs)

	// free up old room, occupy new room
	config.DB.Model(&models.Room{}).Where("id = ?", oldRoomID).
		UpdateColumn("occupied", gorm.Expr("GREATEST(occupied - 1, 0)"))
	config.DB.Model(&models.Room{}).Where("id = ?", body.RoomID).
		UpdateColumn("occupied", gorm.Expr("occupied + 1"))

	c.JSON(http.StatusOK, gin.H{"message": "Student transferred", "data": hs})
}

// CheckoutHostelStudent -> powers the "Checkout" button.
func CheckoutHostelStudent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var hs models.HostelStudent
	if err := config.DB.First(&hs, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hostel student record not found"})
		return
	}

	hs.Status = "checked_out"
	config.DB.Save(&hs)

	config.DB.Model(&models.Room{}).Where("id = ?", hs.RoomID).
		UpdateColumn("occupied", gorm.Expr("GREATEST(occupied - 1, 0)"))

	c.JSON(http.StatusOK, gin.H{"message": "Student checked out", "data": hs})
}

// ============================================================
// HOSTEL FEES (Fees tab - Image 2)
// ============================================================

// GetHostelFees -> reuses HostelStudent records (same data Students tab
// uses) since fee info already lives on HostelStudent. Adds collected /
// pending totals so the summary cards at the top of the Fees tab work.
func GetHostelFees(c *gin.Context) {
	var hostelStudents []models.HostelStudent
	query := config.DB.Model(&models.HostelStudent{}).Where("status = ?", "active")
	if hostelID := c.Query("hostel_id"); hostelID != "" {
		query = query.Where("hostel_id = ?", hostelID)
	}
	query.Find(&hostelStudents)

	var collected, pending float64
	for _, hs := range hostelStudents {
		if hs.FeeStatus == "paid" {
			collected += hs.MonthlyFee
		} else {
			pending += hs.MonthlyFee
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data":      hostelStudents,
		"total":     len(hostelStudents),
		"collected": collected,
		"pending":   pending,
	})
}

// UpdateHostelFeeStatus -> mark a hostel student's fee as paid/pending/overdue
func UpdateHostelFeeStatus(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var hs models.HostelStudent
	if err := config.DB.First(&hs, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hostel student record not found"})
		return
	}

	var body struct {
		FeeStatus string `json:"fee_status"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hs.FeeStatus = body.FeeStatus
	config.DB.Save(&hs)
	c.JSON(http.StatusOK, gin.H{"message": "Fee status updated", "data": hs})
}

// ============================================================
// HOSTEL ATTENDANCE (Attendance tab - Image 3)
// ============================================================

// GetHostelAttendance -> returns attendance for a given date, joined with
// active hostel students so anyone with no record yet still shows up
// (frontend can default their buttons to unselected).
func GetHostelAttendance(c *gin.Context) {
	date := c.Query("date")
	if date == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "date query param is required"})
		return
	}

	var hostelStudents []models.HostelStudent
	query := config.DB.Model(&models.HostelStudent{}).Where("status = ?", "active")
	if hostelID := c.Query("hostel_id"); hostelID != "" {
		query = query.Where("hostel_id = ?", hostelID)
	}
	query.Find(&hostelStudents)

	var records []models.HostelAttendance
	config.DB.Where("date = ?", date).Find(&records)

	statusByStudent := map[uint]string{}
	for _, r := range records {
		statusByStudent[r.HostelStudentID] = r.Status
	}

	type AttendanceRow struct {
		HostelStudentID uint   `json:"hostel_student_id"`
		StudentID       uint   `json:"student_id"`
		RoomID          uint   `json:"room_id"`
		Status          string `json:"status"` // present | absent | leave | "" (not marked)
	}

	rows := make([]AttendanceRow, 0, len(hostelStudents))
	present, absent, leave := 0, 0, 0
	for _, hs := range hostelStudents {
		status := statusByStudent[hs.ID]
		switch status {
		case "present":
			present++
		case "absent":
			absent++
		case "leave":
			leave++
		}
		rows = append(rows, AttendanceRow{
			HostelStudentID: hs.ID,
			StudentID:       hs.StudentID,
			RoomID:          hs.RoomID,
			Status:          status,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    rows,
		"date":    date,
		"present": present,
		"absent":  absent,
		"leave":   leave,
		"total":   len(hostelStudents),
	})
}

// MarkHostelAttendance -> powers the "Save Attendance" button (Image 3).
// Accepts a bulk array of {hostel_student_id, status} for one date and
// upserts each record.
func MarkHostelAttendance(c *gin.Context) {
	var body struct {
		Date    string `json:"date" binding:"required"`
		Records []struct {
			HostelStudentID uint   `json:"hostel_student_id"`
			Status          string `json:"status"`
		} `json:"records" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, rec := range body.Records {
		var existing models.HostelAttendance
		err := config.DB.Where("hostel_student_id = ? AND date = ?", rec.HostelStudentID, body.Date).
			First(&existing).Error

		if err != nil {
			// no existing record -> create
			config.DB.Create(&models.HostelAttendance{
				HostelStudentID: rec.HostelStudentID,
				Date:            body.Date,
				Status:          rec.Status,
			})
		} else {
			// existing record -> update status
			existing.Status = rec.Status
			config.DB.Save(&existing)
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Attendance saved"})
}
func GetComplaints(c *gin.Context) {
	var complaints []models.HostelComplaint
	query := config.DB.Model(&models.HostelComplaint{})
 
	if hostelID := c.Query("hostel_id"); hostelID != "" {
		query = query.Where("hostel_id = ?", hostelID)
	}
	if status := c.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}
 
	query.Order("created_at desc").Find(&complaints)
	c.JSON(http.StatusOK, gin.H{"data": complaints, "total": len(complaints)})
}
 
// GetMyComplaints -> Student ke liye sirf apni complaints
func GetMyComplaints(c *gin.Context) {
	studentID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
 
	var complaints []models.HostelComplaint
	config.DB.Where("student_id = ?", studentID).
		Order("created_at desc").
		Find(&complaints)
 
	c.JSON(http.StatusOK, gin.H{"data": complaints, "total": len(complaints)})
}
 
// CreateComplaint -> Student complaint submit karta hai
func CreateComplaint(c *gin.Context) {
	studentID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
 
	var complaint models.HostelComplaint
	if err := c.ShouldBindJSON(&complaint); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
 
	complaint.StudentID = studentID.(uint)
	if complaint.Status == "" {
		complaint.Status = "pending"
	}
	if complaint.Priority == "" {
		complaint.Priority = "medium"
	}
 
	if err := config.DB.Create(&complaint).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to submit complaint"})
		return
	}
 
	c.JSON(http.StatusCreated, gin.H{"message": "Complaint submitted", "data": complaint})
}
 
// AssignComplaint -> Admin complaint assign karta hai
func AssignComplaint(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var complaint models.HostelComplaint
	if err := config.DB.First(&complaint, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Complaint not found"})
		return
	}
 
	var body struct {
		AssignedTo string `json:"assigned_to" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
 
	complaint.AssignedTo = body.AssignedTo
	complaint.Status = "assigned"
	config.DB.Save(&complaint)
 
	c.JSON(http.StatusOK, gin.H{"message": "Complaint assigned", "data": complaint})
}
 
// ResolveComplaint -> Admin complaint resolve karta hai
func ResolveComplaint(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var complaint models.HostelComplaint
	if err := config.DB.First(&complaint, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Complaint not found"})
		return
	}
 
	complaint.Status = "resolved"
	config.DB.Save(&complaint)
 
	c.JSON(http.StatusOK, gin.H{"message": "Complaint resolved", "data": complaint})
}
func AcceptComplaint(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var complaint models.HostelComplaint
    if err := config.DB.First(&complaint, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Complaint not found"})
        return
    }
    complaint.Status = "accepted"
    config.DB.Save(&complaint)
    c.JSON(http.StatusOK, gin.H{"message": "Complaint accepted", "data": complaint})
}

func RejectComplaint(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var complaint models.HostelComplaint
    if err := config.DB.First(&complaint, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Complaint not found"})
        return
    }
    complaint.Status = "rejected"
    config.DB.Save(&complaint)
    c.JSON(http.StatusOK, gin.H{"message": "Complaint rejected", "data": complaint})
}