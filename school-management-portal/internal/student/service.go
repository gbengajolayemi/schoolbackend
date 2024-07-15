package student

type Service interface {
	CreateStudent(student *Student) error
	GetAllStudents() ([]Student, error)
	UpdateStudent(student *Student) error
	GetStudentByID(studentID int) (*Student, error)
	DeleteStudent(studentID int) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateStudent(student *Student) error {
	return s.repo.CreateStudent(student)
}

func (s *service) GetAllStudents() ([]Student, error) {
	return s.repo.GetAllStudents()
}

func (s *service) UpdateStudent(student *Student) error {
	return s.repo.UpdateStudent(student)
}

func (s *service) GetStudentByID(studentID int) (*Student, error) {
	return s.repo.GetStudentByID(studentID)
}

func (s *service) DeleteStudent(studentID int) error {
	return s.repo.DeleteStudent(studentID)
}
