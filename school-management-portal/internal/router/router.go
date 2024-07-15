package router

import (
	"school-management-portal/internal/student"

	"github.com/gorilla/mux"
)

func NewRouter(studentHandler *student.Handler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/students", studentHandler.CreateStudent).Methods("POST")
	r.HandleFunc("/students", studentHandler.GetAllStudents).Methods("GET")
	r.HandleFunc("/students/{studentID}", studentHandler.GetStudentByID).Methods("GET")

	r.HandleFunc("/students/{studentID}", studentHandler.UpdateStudent).Methods("PUT")
	r.HandleFunc("/students/{studentID}", studentHandler.DeleteStudent).Methods("DELETE")

	return r
}
