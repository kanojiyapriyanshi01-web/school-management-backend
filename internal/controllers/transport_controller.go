package controllers

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
    "school_backend/internal/config"
    "school_backend/internal/models"
)

// ── Vehicles ──────────────────────────────────────────────
func GetVehicles(c *gin.Context) {
    var vehicles []models.Vehicle
    config.DB.Find(&vehicles)
    c.JSON(http.StatusOK, gin.H{"data": vehicles})
}

func CreateVehicle(c *gin.Context) {
    var v models.Vehicle
    if err := c.ShouldBindJSON(&v); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Create(&v)
    c.JSON(http.StatusCreated, gin.H{"message": "Vehicle added", "data": v})
}

func UpdateVehicle(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var v models.Vehicle
    config.DB.First(&v, id)
    c.ShouldBindJSON(&v)
    config.DB.Save(&v)
    c.JSON(http.StatusOK, gin.H{"message": "Vehicle updated", "data": v})
}

func DeleteVehicle(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    config.DB.Delete(&models.Vehicle{}, id)
    c.JSON(http.StatusOK, gin.H{"message": "Vehicle deleted"})
}

func UpdateVehicleLocation(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var body struct {
        Lat   float64 `json:"lat"`
        Lng   float64 `json:"lng"`
        Speed float64 `json:"speed"`
    }
    c.ShouldBindJSON(&body)
    config.DB.Model(&models.Vehicle{}).Where("id = ?", id).Updates(map[string]interface{}{
        "current_lat": body.Lat,
        "current_lng": body.Lng,
        "current_speed": body.Speed,
    })
    c.JSON(http.StatusOK, gin.H{"message": "Location updated"})
}

// ── Routes ────────────────────────────────────────────────
func GetTransportRoutes(c *gin.Context) {
    var routes []models.TransportRoute
    config.DB.Find(&routes)
    c.JSON(http.StatusOK, gin.H{"data": routes})
}

func CreateTransportRoute(c *gin.Context) {
    var r models.TransportRoute
    c.ShouldBindJSON(&r)
    config.DB.Create(&r)
    c.JSON(http.StatusCreated, gin.H{"message": "Route added", "data": r})
}

func UpdateTransportRoute(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var r models.TransportRoute
    config.DB.First(&r, id)
    c.ShouldBindJSON(&r)
    config.DB.Save(&r)
    c.JSON(http.StatusOK, gin.H{"message": "Route updated", "data": r})
}

func DeleteTransportRoute(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    config.DB.Delete(&models.TransportRoute{}, id)
    c.JSON(http.StatusOK, gin.H{"message": "Route deleted"})
}

// ── Student Transport ─────────────────────────────────────
func GetStudentTransports(c *gin.Context) {
    var students []models.StudentTransport
    config.DB.Find(&students)
    c.JSON(http.StatusOK, gin.H{"data": students})
}

func AssignStudentTransport(c *gin.Context) {
    var st models.StudentTransport
    c.ShouldBindJSON(&st)
    config.DB.Create(&st)
    c.JSON(http.StatusCreated, gin.H{"message": "Transport assigned", "data": st})
}

func UpdateStudentTransport(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var st models.StudentTransport
    config.DB.First(&st, id)
    c.ShouldBindJSON(&st)
    config.DB.Save(&st)
    c.JSON(http.StatusOK, gin.H{"message": "Updated", "data": st})
}

// ── Transport Fees ────────────────────────────────────────
func GetTransportFees(c *gin.Context) {
    var fees []models.TransportFee
    query := config.DB.Model(&models.TransportFee{})
    if sid := c.Query("student_id"); sid != "" {
        query = query.Where("student_id = ?", sid)
    }
    query.Find(&fees)
    c.JSON(http.StatusOK, gin.H{"data": fees})
}

func CreateTransportFee(c *gin.Context) {
    var fee models.TransportFee
    c.ShouldBindJSON(&fee)
    config.DB.Create(&fee)
    c.JSON(http.StatusCreated, gin.H{"message": "Fee created", "data": fee})
}

func UpdateTransportFeeStatus(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var fee models.TransportFee
    config.DB.First(&fee, id)
    var body struct{ Status string `json:"status"`; PaidDate string `json:"paid_date"` }
    c.ShouldBindJSON(&body)
    fee.Status = body.Status
    fee.PaidDate = body.PaidDate
    config.DB.Save(&fee)
    c.JSON(http.StatusOK, gin.H{"message": "Fee updated", "data": fee})
}
