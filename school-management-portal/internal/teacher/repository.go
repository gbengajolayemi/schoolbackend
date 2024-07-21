package teacher

import (
	"database/sql"
	"errors"
)

type Repository interface {
	CreateTeacher(teacher *Teacher) (int, error)
	GetTeacher(id int) (*Teacher, error)
	UpdateTeacher(teacher *Teacher) error
	DeleteTeacher(id int) error
	ListTeachers() ([]*Teacher, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateTeacher(teacher *Teacher) (int, error) {
	query := `INSERT INTO Teachers (FirstName, LastName, DateOfBirth, Gender, Email, PhoneNumber, 
              Address, City, State, ZipCode, Subject, HireDate, Qualifications, Salary) 
              VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := r.db.Exec(query, teacher.FirstName, teacher.LastName, teacher.DateOfBirth, teacher.Gender,
		teacher.Email, teacher.PhoneNumber, teacher.Address, teacher.City, teacher.State, teacher.ZipCode,
		teacher.Subject, teacher.HireDate, teacher.Qualifications, teacher.Salary)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) GetTeacher(id int) (*Teacher, error) {
	query := `SELECT * FROM Teachers WHERE TeacherID = ?`
	var teacher Teacher
	err := r.db.QueryRow(query, id).Scan(
		&teacher.TeacherID, &teacher.FirstName, &teacher.LastName, &teacher.DateOfBirth,
		&teacher.Gender, &teacher.Email, &teacher.PhoneNumber, &teacher.Address,
		&teacher.City, &teacher.State, &teacher.ZipCode, &teacher.Subject,
		&teacher.HireDate, &teacher.Qualifications, &teacher.Salary,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("teacher not found")
		}
		return nil, err
	}
	return &teacher, nil
}

func (r *repository) UpdateTeacher(teacher *Teacher) error {
	query := `UPDATE Teachers SET FirstName=?, LastName=?, DateOfBirth=?, Gender=?, Email=?, 
              PhoneNumber=?, Address=?, City=?, State=?, ZipCode=?, Subject=?, HireDate=?, 
              Qualifications=?, Salary=? WHERE TeacherID=?`

	_, err := r.db.Exec(query, teacher.FirstName, teacher.LastName, teacher.DateOfBirth, teacher.Gender,
		teacher.Email, teacher.PhoneNumber, teacher.Address, teacher.City, teacher.State, teacher.ZipCode,
		teacher.Subject, teacher.HireDate, teacher.Qualifications, teacher.Salary, teacher.TeacherID)

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteTeacher(id int) error {
	query := `DELETE FROM Teachers WHERE TeacherID = ?`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *repository) ListTeachers() ([]*Teacher, error) {
	query := `SELECT * FROM Teachers`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teachers []*Teacher
	for rows.Next() {
		var t Teacher
		err := rows.Scan(
			&t.TeacherID, &t.FirstName, &t.LastName, &t.DateOfBirth,
			&t.Gender, &t.Email, &t.PhoneNumber, &t.Address,
			&t.City, &t.State, &t.ZipCode, &t.Subject,
			&t.HireDate, &t.Qualifications, &t.Salary,
		)
		if err != nil {
			return nil, err
		}
		teachers = append(teachers, &t)
	}
	return teachers, nil
}
