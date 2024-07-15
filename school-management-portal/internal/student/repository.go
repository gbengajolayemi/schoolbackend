package student

import (
	"database/sql"
)

type Repository interface {
	CreateStudent(student *Student) error
	GetAllStudents() ([]Student, error)
	UpdateStudent(student *Student) error
	GetStudentByID(studentID int) (*Student, error)
	DeleteStudent(studentID int) error

}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateStudent(student *Student) error {
	query := `INSERT INTO Students (FirstName, LastName, DateOfBirth, Gender, Address, City, State, ZipCode, 
              PhoneNumber, Email, ParentName, ParentPhoneNumber, ParentEmail, ClassID, Section, EnrollmentDate, 
              EmergencyContactName, EmergencyContactPhoneNumber, Allergies, MedicalConditions, PhotoURL) 
              VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := r.db.Exec(query, student.FirstName, student.LastName, student.DateOfBirth, student.Gender,
		student.Address, student.City, student.State, student.ZipCode, student.PhoneNumber, student.Email,
		student.ParentName, student.ParentPhoneNumber, student.ParentEmail, student.ClassID, student.Section,
		student.EnrollmentDate, student.EmergencyContactName, student.EmergencyContactPhoneNumber,
		student.Allergies, student.MedicalConditions, student.PhotoURL)

	return err
}

func (r *repository) GetAllStudents() ([]Student, error) {
	query := `SELECT StudentID, FirstName, LastName, DateOfBirth, Gender, Address, City, State, ZipCode, 
              PhoneNumber, Email, ParentName, ParentPhoneNumber, ParentEmail, ClassID, Section, EnrollmentDate, 
              EmergencyContactName, EmergencyContactPhoneNumber, Allergies, MedicalConditions, PhotoURL FROM Students`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var student Student
		if err := rows.Scan(&student.StudentID, &student.FirstName, &student.LastName, &student.DateOfBirth, &student.Gender,
			&student.Address, &student.City, &student.State, &student.ZipCode, &student.PhoneNumber, &student.Email,
			&student.ParentName, &student.ParentPhoneNumber, &student.ParentEmail, &student.ClassID, &student.Section,
			&student.EnrollmentDate, &student.EmergencyContactName, &student.EmergencyContactPhoneNumber,
			&student.Allergies, &student.MedicalConditions, &student.PhotoURL); err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}

func (r *repository) UpdateStudent(student *Student) error {
	query := `UPDATE Students SET FirstName=?, LastName=?, DateOfBirth=?, Gender=?, Address=?, City=?, State=?, 
              ZipCode=?, PhoneNumber=?, Email=?, ParentName=?, ParentPhoneNumber=?, ParentEmail=?, ClassID=?, 
              Section=?, EnrollmentDate=?, EmergencyContactName=?, EmergencyContactPhoneNumber=?, Allergies=?, 
              MedicalConditions=?, PhotoURL=? WHERE StudentID=?`

	_, err := r.db.Exec(query, student.FirstName, student.LastName, student.DateOfBirth, student.Gender,
		student.Address, student.City, student.State, student.ZipCode, student.PhoneNumber, student.Email,
		student.ParentName, student.ParentPhoneNumber, student.ParentEmail, student.ClassID, student.Section,
		student.EnrollmentDate, student.EmergencyContactName, student.EmergencyContactPhoneNumber,
		student.Allergies, student.MedicalConditions, student.PhotoURL, student.StudentID)

	return err
}

func (r *repository) GetStudentByID(studentID int) (*Student, error) {
	query := `SELECT StudentID, FirstName, LastName, DateOfBirth, Gender, Address, City, State,
                    ZipCode, PhoneNumber, Email, ParentName, ParentPhoneNumber, ParentEmail,
                    ClassID, Section, EnrollmentDate, EmergencyContactName, EmergencyContactPhoneNumber,
                    Allergies, MedicalConditions, PhotoURL
              FROM Students
              WHERE StudentID = ?`

	var student Student
	err := r.db.QueryRow(query, studentID).Scan(
		&student.StudentID, &student.FirstName, &student.LastName, &student.DateOfBirth, &student.Gender,
		&student.Address, &student.City, &student.State, &student.ZipCode, &student.PhoneNumber, &student.Email,
		&student.ParentName, &student.ParentPhoneNumber, &student.ParentEmail, &student.ClassID, &student.Section,
		&student.EnrollmentDate, &student.EmergencyContactName, &student.EmergencyContactPhoneNumber,
		&student.Allergies, &student.MedicalConditions, &student.PhotoURL,
	)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (r *repository) DeleteStudent(studentID int) error {
    query := `DELETE FROM Students WHERE StudentID=?`

    _, err := r.db.Exec(query, studentID)
    return err
}
