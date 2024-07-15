package student

import (
	"encoding/json"
	"net/http"
	"school-management-portal/pkg/response"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = h.service.CreateStudent(&student)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to create student")
		return
	}

	response.JSON(w, http.StatusCreated, map[string]string{"message": "Student created successfully"})
}

func (h *Handler) GetAllStudents(w http.ResponseWriter, r *http.Request) {
	students, err := h.service.GetAllStudents()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to get students")
		return
	}

	response.JSON(w, http.StatusOK, students)
}

func (h *Handler) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	var student Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	vars := mux.Vars(r)
	studentID, err := strconv.Atoi(vars["studentID"])
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid student ID")
		return
	}
	student.StudentID = studentID

	err = h.service.UpdateStudent(&student)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to update student")
		return
	}

	response.JSON(w, http.StatusOK, map[string]string{"message": "Student updated successfully"})
}

func (h *Handler) GetStudentByID(w http.ResponseWriter, r *http.Request) {
	// Extract studentID from URL parameters
	vars := mux.Vars(r)
	studentIDStr := vars["studentID"]

	// Convert studentIDStr to an integer
	studentID, err := strconv.Atoi(studentIDStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid student ID")
		return
	}

	// Call service to retrieve student by ID
	student, err := h.service.GetStudentByID(studentID)
	if err != nil {
		response.Error(w, http.StatusNotFound, "Student not found")
		return
	}

	// Return the student as JSON response
	response.JSON(w, http.StatusOK, student)
}

func (h *Handler) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentIDStr := vars["studentID"]

	studentID, err := strconv.Atoi(studentIDStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid student ID")
		return
	}

	err = h.service.DeleteStudent(studentID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to delete student")
		return
	}

	response.JSON(w, http.StatusOK, map[string]string{"message": "Student deleted successfully"})
}
