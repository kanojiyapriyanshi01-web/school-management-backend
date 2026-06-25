package routes

import (
    "github.com/gin-gonic/gin"
    "school_backend/internal/auth"
    "school_backend/internal/controllers"
)

func SetupRoutes(r *gin.Engine) {
    api := r.Group("/api/v1")

    api.POST("/auth/login",    controllers.Login)
    api.POST("/auth/register", controllers.Register)
    api.GET("/auth/seed",      controllers.SeedAdmin)
    api.GET("/auth/seed-classes", controllers.SeedClasses)

    // Public admission form
    api.POST("/admissions", controllers.CreateAdmission)

    protected := api.Group("/")
    protected.Use(auth.AuthMiddleware())
    {
        protected.GET("auth/me", controllers.GetMe)

        // Classes
        protected.GET("classes",        controllers.GetClasses)
        protected.POST("classes",       controllers.CreateClass)
        protected.GET("classes/:id",    controllers.GetClassByID)
        protected.PUT("classes/:id",    controllers.UpdateClass)
        protected.DELETE("classes/:id", controllers.DeleteClass)

        // Admissions
        protected.GET("admissions",            controllers.GetAdmissions)
        protected.GET("admissions/:id",        controllers.GetAdmissionByID)
        protected.PUT("admissions/:id/status", controllers.UpdateAdmissionStatus)
        protected.DELETE("admissions/:id",     controllers.DeleteAdmission)

        // Students
        protected.GET("students",        controllers.GetStudents)
        protected.POST("students",       controllers.CreateStudent)
        protected.GET("students/:id",    controllers.GetStudentByID)
        protected.PUT("students/:id",    controllers.UpdateStudent)
        protected.DELETE("students/:id", controllers.DeleteStudent)

        // Staff
        protected.GET("staff",        controllers.GetStaff)
        protected.POST("staff",       controllers.CreateStaff)
        protected.PUT("staff/:id",    controllers.UpdateStaff)
        protected.DELETE("staff/:id", controllers.DeleteStaff)

        // Attendance
        protected.GET("attendance",  controllers.GetAttendance)
        protected.POST("attendance", controllers.MarkAttendance)

        // Fees
        protected.GET("fees",      controllers.GetFees)
        protected.POST("fees",     controllers.CreateFee)
        protected.PUT("fees/:id",  controllers.UpdateFee)

        // Notices
        protected.GET("notices",        controllers.GetNotices)
        protected.POST("notices",       controllers.CreateNotice)
        protected.DELETE("notices/:id", controllers.DeleteNotice)

        // Exams & Marks
        protected.GET("exams",         controllers.GetExams)
        protected.POST("exams",        controllers.CreateExam)
        protected.PUT("exams/:id",     controllers.UpdateExam)
        protected.GET("marks",         controllers.GetMarks)
        protected.POST("marks",        controllers.SaveMarks)

        // Library
        protected.GET("books",            controllers.GetBooks)
        protected.POST("books",           controllers.CreateBook)
        protected.POST("books/issue",     controllers.IssueBook)
        protected.PUT("books/return/:id", controllers.ReturnBook)
        protected.GET("books/my",         controllers.GetMyBooks)

        // Hostel
        protected.GET("hostel",        controllers.GetHostels)
        protected.POST("hostel",       controllers.CreateHostel)
        protected.GET("hostel/:id",    controllers.GetHostelByID)
        protected.PUT("hostel/:id",    controllers.UpdateHostel)
        protected.DELETE("hostel/:id", controllers.DeleteHostel)

        // Hostel Rooms
        protected.GET("hostel/rooms",        controllers.GetRooms)
        protected.POST("hostel/rooms",       controllers.CreateRoom)
        protected.GET("hostel/rooms/:id",    controllers.GetRoomByID)
        protected.PUT("hostel/rooms/:id",    controllers.UpdateRoom)
        protected.DELETE("hostel/rooms/:id", controllers.DeleteRoom)

        // Hostel Students
        protected.GET("hostel/students",              controllers.GetHostelStudents)
        protected.GET("hostel/students/:id",          controllers.GetHostelStudentByID)
        protected.POST("hostel/allocate",             controllers.AllocateRoom)
        protected.PUT("hostel/students/:id/transfer", controllers.TransferRoom)
        protected.PUT("hostel/students/:id/checkout", controllers.CheckoutHostelStudent)

        // Hostel Fees
        protected.GET("hostel/fees",     controllers.GetHostelFees)
        protected.PUT("hostel/fees/:id", controllers.UpdateHostelFeeStatus)

        // Transport
             protected.GET("transport/vehicles", controllers.GetVehicles)
             protected.POST("transport/vehicles", controllers.CreateVehicle)
             protected.PUT("transport/vehicles/:id", controllers.UpdateVehicle)
             protected.DELETE("transport/vehicles/:id", controllers.DeleteVehicle)
             protected.PUT("transport/vehicles/:id/location", controllers.UpdateVehicleLocation)
             protected.GET("transport/routes", controllers.GetTransportRoutes)
             protected.POST("transport/routes", controllers.CreateTransportRoute)
             protected.PUT("transport/routes/:id", controllers.UpdateTransportRoute)
             protected.DELETE("transport/routes/:id", controllers.DeleteTransportRoute)
             protected.GET("transport/students", controllers.GetStudentTransports)
             protected.POST("transport/students", controllers.AssignStudentTransport)
             protected.PUT("transport/students/:id", controllers.UpdateStudentTransport)
             protected.GET("transport/fees", controllers.GetTransportFees)
             protected.POST("transport/fees", controllers.CreateTransportFee)
             protected.PUT("transport/fees/:id/status", controllers.UpdateTransportFeeStatus)

             // Complaints
             protected.GET("complaints", controllers.GetComplaints)
             protected.POST("complaints", controllers.CreateComplaint)
             protected.PUT("complaints/:id/accept", controllers.AcceptComplaint)
             protected.PUT("complaints/:id/reject", controllers.RejectComplaint)

             // Hostel Attendance
        protected.GET("hostel/attendance",  controllers.GetHostelAttendance)
        protected.POST("hostel/attendance", controllers.MarkHostelAttendance)

        // Hostel Complaints
        // ⚠️  IMPORTANT: specific routes pehle likhni hain (:id se pehle)
        protected.GET("hostel/complaints",             controllers.GetComplaints)     // Admin: saari
        protected.GET("hostel/complaints/my",          controllers.GetMyComplaints)   // Student: apni
        protected.POST("hostel/complaints",            controllers.CreateComplaint)   // Student: naya
        protected.PUT("hostel/complaints/:id/assign",  controllers.AssignComplaint)   // Admin
        protected.PUT("hostel/complaints/:id/resolve", controllers.ResolveComplaint)  // Admin
        protected.PUT("hostel/complaints/:id/accept",  controllers.AcceptComplaint)   // ✅ Admin NEW
        protected.PUT("hostel/complaints/:id/reject",  controllers.RejectComplaint)   // ✅ Admin NEW
    }
}


