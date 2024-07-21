package teacher

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"school-management-portal/pkg/response"
	"strconv"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateTeacher(w http.ResponseWriter, r *http.Request) {
	var teacher Teacher
	err := json.NewDecoder(r.Body).Decode(&teacher)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		response.Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	log.Printf("Attempting to create teacher: %+v", teacher)

	id, err := h.service.CreateTeacher(&teacher)
	if err != nil {
		log.Printf("Error creating teacher: %v", err)

		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				response.Error(w, http.StatusConflict, "A teacher with this email already exists")
				return
			// Add more cases here for other MySQL error codes as needed
			default:
				response.Error(w, http.StatusInternalServerError, "Database error occurred")
				return
			}
		}

		if errors.Is(err, sql.ErrNoRows) {
			response.Error(w, http.StatusNotFound, "Resource not found")
			return
		}

		// For any other errors
		response.Error(w, http.StatusInternalServerError, "Failed to create teacher")
		return
	}

	log.Printf("Teacher created successfully with ID: %d", id)

	teacher.TeacherID = id
	response.JSON(w, http.StatusCreated, struct {
		Message string  `json:"message"`
		Teacher Teacher `json:"teacher"`
	}{

		Message: "Teacher created successfully",
		Teacher: teacher,
	})
}

func (h *Handler) GetTeacher(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid teacher ID")
		return
	}

	teacher, err := h.service.GetTeacher(id)
	if err != nil {
		response.Error(w, http.StatusNotFound, "Teacher not found")
		return
	}

	response.JSON(w, http.StatusOK, teacher)
}

func (h *Handler) UpdateTeacher(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teacherID, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid teacher ID")
		return
	}

	// Decode request body to get updated teacher data
	var updatedTeacher Teacher
	err = json.NewDecoder(r.Body).Decode(&updatedTeacher)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	updatedTeacher.TeacherID = teacherID // Ensure teacher ID matches the path parameter

	// Check if the teacher exists
	existingTeacher, err := h.service.GetTeacher(teacherID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to fetch teacher")
		return
	}
	if existingTeacher == nil {
		response.Error(w, http.StatusNotFound, "No teacher found with that ID")
		return
	}

	// Update the teacher
	err = h.service.UpdateTeacher(&updatedTeacher)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to update teacher")
		return
	}

	// Fetch the updated teacher data after the update
	updatedTeacherData, err := h.service.GetTeacher(teacherID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to fetch updated teacher data")
		return
	}

	response.JSON(w, http.StatusOK, updatedTeacherData)
}

func (h *Handler) DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid teacher ID")
		return
	}

	// Check if the teacher exists
	_, err = h.service.GetTeacher(id)
	if err != nil {
		if err == sql.ErrNoRows {
			response.Error(w, http.StatusNotFound, "Teacher not found")
		} else {
			response.Error(w, http.StatusInternalServerError, "Failed to fetch teacher")
		}
		return
	}

	// Perform the delete operation
	err = h.service.DeleteTeacher(id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to delete teacher")
		return
	}

	response.JSON(w, http.StatusOK, map[string]string{
		"message": "Teacher deleted successfully",
	})
}

func (h *Handler) ListTeachers(w http.ResponseWriter, r *http.Request) {
	teachers, err := h.service.ListTeachers()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to retrieve teachers")
		return
	}

	response.JSON(w, http.StatusOK, teachers)
}
