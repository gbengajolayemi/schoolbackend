package student

import (
	"encoding/json"
	"net/http"
	"strconv"

	"school-management-portal/pkg/response"

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

	id, err := h.service.CreateStudent(&student)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to create student")
		return
	}

	student.StudentID = id
	response.JSON(w, http.StatusCreated, struct {
		Message string  `json:"message"`
		Student Student `json:"student"`
	}{
		Message: "Student created successfully",
		Student: student,
	})
}

func (h *Handler) GetStudentByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["studentID"])
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid student ID")
		return
	}

	student, err := h.service.GetStudentByID(id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	if student == nil {
		response.Error(w, http.StatusNotFound, "Student not found")
		return
	}

	response.JSON(w, http.StatusOK, student)
}

func (h *Handler) GetAllStudents(w http.ResponseWriter, r *http.Request) {
	students, err := h.service.GetAllStudents()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to fetch students")
		return
	}

	response.JSON(w, http.StatusOK, students)
}

func (h *Handler) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID, err := strconv.Atoi(vars["studentID"])
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid student ID")
		return
	}

	// Decode request body to get updated student data
	var updatedStudent Student
	err = json.NewDecoder(r.Body).Decode(&updatedStudent)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	updatedStudent.StudentID = studentID // Ensure student ID matches the path parameter

	// Check if the student exists
	existingStudent, err := h.service.GetStudentByID(studentID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to fetch student")
		return
	}
	if existingStudent == nil {
		response.Error(w, http.StatusNotFound, "No student found with that ID")
		return
	}

	// Update the student
	err = h.service.UpdateStudent(studentID, &updatedStudent)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to update student")
		return
	}

	// Fetch the updated student data after the update
	updatedStudentData, err := h.service.GetStudentByID(studentID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to fetch updated student data")
		return
	}

	response.JSON(w, http.StatusOK, updatedStudentData)
}

func (h *Handler) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID, err := strconv.Atoi(vars["studentID"])
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid student ID")
		return
	}

	// Check if the student exists
	existingStudent, err := h.service.GetStudentByID(studentID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to fetch student")
		return
	}
	if existingStudent == nil {
		response.Error(w, http.StatusNotFound, "No student found with that ID")
		return
	}

	err = h.service.DeleteStudent(studentID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to delete student")
		return
	}

	response.JSON(w, http.StatusOK, map[string]string{"message": "Student deleted successfully", "studentID": strconv.Itoa(studentID)})
}
