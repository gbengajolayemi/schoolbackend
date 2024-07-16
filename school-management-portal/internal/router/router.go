package router

import (
	"school-management-portal/internal/student"
	"school-management-portal/internal/teacher"

	"github.com/gorilla/mux"
)

func NewRouter(studentHandler *student.Handler, teacherHandler *teacher.Handler) *mux.Router {
	r := mux.NewRouter()

	// Student routes
	r.HandleFunc("/students", studentHandler.CreateStudent).Methods("POST")
	r.HandleFunc("/students", studentHandler.GetAllStudents).Methods("GET")
	r.HandleFunc("/students/{studentID}", studentHandler.GetStudentByID).Methods("GET")
	r.HandleFunc("/students/{studentID}", studentHandler.UpdateStudent).Methods("PUT")
	r.HandleFunc("/students/{studentID}", studentHandler.DeleteStudent).Methods("DELETE")

	// Teacher routes
	r.HandleFunc("/teachers", teacherHandler.CreateTeacher).Methods("POST")
	r.HandleFunc("/teachers", teacherHandler.ListTeachers).Methods("GET")
	r.HandleFunc("/teachers/{id}", teacherHandler.GetTeacher).Methods("GET")
	r.HandleFunc("/teachers/{id}", teacherHandler.UpdateTeacher).Methods("PUT")
	r.HandleFunc("/teachers/{id}", teacherHandler.DeleteTeacher).Methods("DELETE")
	return r
}
