package student

type Service interface {
	CreateStudent(student *Student) (int, error)
	GetStudentByID(id int) (*Student, error)
	GetAllStudents() ([]Student, error)
	UpdateStudent(id int, student *Student) error
	DeleteStudent(id int) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateStudent(student *Student) (int, error) {
	return s.repo.CreateStudent(student)
}

func (s *service) GetStudentByID(id int) (*Student, error) {
	return s.repo.GetStudentByID(id)
}

func (s *service) GetAllStudents() ([]Student, error) {
	return s.repo.GetAllStudents()
}

func (s *service) UpdateStudent(id int, student *Student) error {
	return s.repo.UpdateStudent(id, student)
}

func (s *service) DeleteStudent(id int) error {
	return s.repo.DeleteStudent(id)
}
