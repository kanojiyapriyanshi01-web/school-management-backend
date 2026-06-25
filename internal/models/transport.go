package models

import "time"

type Vehicle struct {
    ID                 uint      `gorm:"primaryKey" json:"id"`
    VehicleNumber      string    `json:"vehicle_number"`
    VehicleType        string    `json:"vehicle_type"`
    SeatingCapacity    int       `json:"seating_capacity"`
    DriverName         string    `json:"driver_name"`
    DriverPhone        string    `json:"driver_phone"`
    DriverLicense      string    `json:"driver_license"`
    ConductorName      string    `json:"conductor_name"`
    ConductorPhone     string    `json:"conductor_phone"`
    AssignedRoute      string    `json:"assigned_route"`
    FuelType           string    `json:"fuel_type"`
    InsuranceExpiry    string    `json:"insurance_expiry"`
    FitnessExpiry      string    `json:"fitness_expiry"`
    PollutionExpiry    string    `json:"pollution_expiry"`
    GPSEnabled         bool      `gorm:"default:true" json:"gps_enabled"`
    IsAC               bool      `json:"is_ac"`
    Status             string    `gorm:"default:active" json:"status"`
    CurrentLat         float64   `json:"current_lat"`
    CurrentLng         float64   `json:"current_lng"`
    CurrentSpeed       float64   `json:"current_speed"`
    Notes              string    `json:"notes"`
    CreatedAt          time.Time `json:"created_at"`
    UpdatedAt          time.Time `json:"updated_at"`
}

type TransportRoute struct {
    ID              uint      `gorm:"primaryKey" json:"id"`
    RouteName       string    `json:"route_name"`
    RouteCode       string    `json:"route_code"`
    StartPoint      string    `json:"start_point"`
    EndPoint        string    `json:"end_point"`
    TotalDistance   float64   `json:"total_distance"`
    Duration        string    `json:"duration"`
    MorningTime     string    `json:"morning_time"`
    EveningTime     string    `json:"evening_time"`
    AssignedVehicle uint      `json:"assigned_vehicle"`
    MonthlyFee      float64   `json:"monthly_fee"`
    Stops           string    `json:"stops"`
    Status          string    `gorm:"default:active" json:"status"`
    CreatedAt       time.Time `json:"created_at"`
}

type StudentTransport struct {
    ID              uint      `gorm:"primaryKey" json:"id"`
    StudentID       uint      `json:"student_id"`
    StudentName     string    `json:"student_name"`
    VehicleID       uint      `json:"vehicle_id"`
    RouteID         uint      `json:"route_id"`
    PickupStop      string    `json:"pickup_stop"`
    DropStop        string    `json:"drop_stop"`
    MonthlyFee      float64   `json:"monthly_fee"`
    Status          string    `gorm:"default:active" json:"status"`
    StartDate       string    `json:"start_date"`
    CreatedAt       time.Time `json:"created_at"`
}

type TransportFee struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    StudentID   uint      `json:"student_id"`
    StudentName string    `json:"student_name"`
    VehicleID   uint      `json:"vehicle_id"`
    RouteID     uint      `json:"route_id"`
    Amount      float64   `json:"amount"`
    Month       string    `json:"month"`
    Status      string    `gorm:"default:pending" json:"status"`
    PaidDate    string    `json:"paid_date"`
    CreatedAt   time.Time `json:"created_at"`
}
