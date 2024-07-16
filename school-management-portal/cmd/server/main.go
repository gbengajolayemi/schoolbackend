package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"school-management-portal/internal/db"
	"school-management-portal/internal/router"
	"school-management-portal/internal/student"
	"school-management-portal/internal/teacher"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
    // Load environment variables from .env file
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Get database configuration
    dbUsername := os.Getenv("DB_USERNAME")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPortStr := os.Getenv("DB_PORT")
    dbPort, err := strconv.Atoi(dbPortStr)
    if err != nil {
        log.Fatalf("Invalid DB_PORT: %v", err)
    }
    dbName := os.Getenv("DB_NAME")

    // Initialize database connection
    db.InitializeDB(dbUsername, dbPassword, dbHost, dbPort, dbName)
    defer db.CloseDB()

    // Test database connection
    err = db.DB().Ping()
    if err != nil {
        log.Fatalf("Failed to ping database: %v", err)
    }

    // Initialize student repository, service, and handler
    studentRepository := student.NewRepository(db.DB())
    studentService := student.NewService(studentRepository)
    studentHandler := student.NewHandler(studentService)

    // Initialize teacher repository, service, and handler
    teacherRepository := teacher.NewRepository(db.DB())
    teacherService := teacher.NewService(teacherRepository)
    teacherHandler := teacher.NewHandler(teacherService)

    // Initialize router
    r := router.NewRouter(studentHandler, teacherHandler)

    // Start server
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Default port if not specified
    }
    log.Printf("Starting server on port %s", port)
    if err := http.ListenAndServe(":"+port, r); err != nil {
        log.Fatalf("Could not start server: %s", err.Error())
    }
}
