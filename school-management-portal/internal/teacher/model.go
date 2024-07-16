package teacher

import (
	"time"
)

type Teacher struct {
	TeacherID      int       `json:"teacherId,omitempty"`
	FirstName      string    `json:"firstName"`
	LastName       string    `json:"lastName"`
	DateOfBirth    time.Time `json:"dateOfBirth"`
	Gender         string    `json:"gender"`
	Email          string    `json:"email"`
	PhoneNumber    string    `json:"phoneNumber"`
	Address        string    `json:"address"`
	City           string    `json:"city"`
	State          string    `json:"state"`
	ZipCode        string    `json:"zipCode"`
	Subject        string    `json:"subject"`
	HireDate       time.Time `json:"hireDate"`
	Qualifications string    `json:"qualifications"`
	Salary         float64   `json:"salary"`
}
