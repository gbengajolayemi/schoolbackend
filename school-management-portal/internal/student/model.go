package student

import (
	"time"
)

type Student struct {
	StudentID                   int       `json:"studentId,omitempty"`
	FirstName                   string    `json:"firstName"`
	LastName                    string    `json:"lastName"`
	DateOfBirth                 time.Time `json:"dateOfBirth"`
	Gender                      string    `json:"gender"`
	Address                     string    `json:"address"`
	City                        string    `json:"city"`
	State                       string    `json:"state"`
	ZipCode                     string    `json:"zipCode"`
	PhoneNumber                 string    `json:"phoneNumber"`
	Email                       string    `json:"email,omitempty"`
	ParentName                  string    `json:"parentName"`
	ParentPhoneNumber           string    `json:"parentPhoneNumber"`
	ParentEmail                 string    `json:"parentEmail,omitempty"`
	ClassID                     int       `json:"classId"`
	Section                     string    `json:"section,omitempty"`
	EnrollmentDate              time.Time `json:"enrollmentDate"`
	EmergencyContactName        string    `json:"emergencyContactName,omitempty"`
	EmergencyContactPhoneNumber string    `json:"emergencyContactPhoneNumber,omitempty"`
	Allergies                   string    `json:"allergies,omitempty"`
	MedicalConditions           string    `json:"medicalConditions,omitempty"`
	PhotoURL                    string    `json:"photoUrl,omitempty"`
}
