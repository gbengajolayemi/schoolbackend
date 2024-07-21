package student

import (
	"database/sql"
	"errors"
)

type Repository interface {
	CreateStudent(student *Student) (int, error)
	GetStudentByID(id int) (*Student, error)
	GetAllStudents() ([]Student, error)
	UpdateStudent(id int, student *Student) error
	DeleteStudent(id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateStudent(student *Student) (int, error) {
	query := `INSERT INTO Students (FirstName, LastName, DateOfBirth, Gender, Address, City, State, ZipCode, 
              PhoneNumber, Email, ParentName, ParentPhoneNumber, ParentEmail, ClassID, Section, EnrollmentDate, 
              EmergencyContactName, EmergencyContactPhoneNumber, Allergies, MedicalConditions, PhotoURL) 
              VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := r.db.Exec(query, student.FirstName, student.LastName, student.DateOfBirth, student.Gender,
		student.Address, student.City, student.State, student.ZipCode, student.PhoneNumber, student.Email,
		student.ParentName, student.ParentPhoneNumber, student.ParentEmail, student.ClassID, student.Section,
		student.EnrollmentDate, student.EmergencyContactName, student.EmergencyContactPhoneNumber,
		student.Allergies, student.MedicalConditions, student.PhotoURL)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) GetStudentByID(id int) (*Student, error) {
	var student Student
	query := `SELECT StudentID, FirstName, LastName, DateOfBirth, Gender, Address, City, State, ZipCode, 
              PhoneNumber, Email, ParentName, ParentPhoneNumber, ParentEmail, ClassID, Section, EnrollmentDate, 
              EmergencyContactName, EmergencyContactPhoneNumber, Allergies, MedicalConditions, PhotoURL 
              FROM Students WHERE StudentID = ?`

	err := r.db.QueryRow(query, id).Scan(&student.StudentID, &student.FirstName, &student.LastName, &student.DateOfBirth,
		&student.Gender, &student.Address, &student.City, &student.State, &student.ZipCode, &student.PhoneNumber,
		&student.Email, &student.ParentName, &student.ParentPhoneNumber, &student.ParentEmail, &student.ClassID,
		&student.Section, &student.EnrollmentDate, &student.EmergencyContactName, &student.EmergencyContactPhoneNumber,
		&student.Allergies, &student.MedicalConditions, &student.PhotoURL)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &student, nil
}

func (r *repository) GetAllStudents() ([]Student, error) {
	query := `SELECT StudentID, FirstName, LastName, DateOfBirth, Gender, Address, City, State, ZipCode, 
              PhoneNumber, Email, ParentName, ParentPhoneNumber, ParentEmail, ClassID, Section, EnrollmentDate, 
              EmergencyContactName, EmergencyContactPhoneNumber, Allergies, MedicalConditions, PhotoURL 
              FROM Students`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var student Student
		err := rows.Scan(&student.StudentID, &student.FirstName, &student.LastName, &student.DateOfBirth, &student.Gender,
			&student.Address, &student.City, &student.State, &student.ZipCode, &student.PhoneNumber, &student.Email,
			&student.ParentName, &student.ParentPhoneNumber, &student.ParentEmail, &student.ClassID, &student.Section,
			&student.EnrollmentDate, &student.EmergencyContactName, &student.EmergencyContactPhoneNumber, &student.Allergies,
			&student.MedicalConditions, &student.PhotoURL)

		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return students, nil
}

func (r *repository) UpdateStudent(studentID int, student *Student) error {
	query := `UPDATE Students SET FirstName=?, LastName=?, DateOfBirth=?, Gender=?, Address=?, City=?, 
              State=?, ZipCode=?, PhoneNumber=?, Email=?, ParentName=?, ParentPhoneNumber=?, ParentEmail=?, 
              ClassID=?, Section=?, EnrollmentDate=?, EmergencyContactName=?, EmergencyContactPhoneNumber=?, 
              Allergies=?, MedicalConditions=?, PhotoURL=? WHERE StudentID=?`

	_, err := r.db.Exec(query, student.FirstName, student.LastName, student.DateOfBirth, student.Gender,
		student.Address, student.City, student.State, student.ZipCode, student.PhoneNumber, student.Email,
		student.ParentName, student.ParentPhoneNumber, student.ParentEmail, student.ClassID, student.Section,
		student.EnrollmentDate, student.EmergencyContactName, student.EmergencyContactPhoneNumber,
		student.Allergies, student.MedicalConditions, student.PhotoURL, studentID)

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteStudent(id int) error {
	query := `DELETE FROM Students WHERE StudentID = ?`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("student not found")
	}

	return nil
}
