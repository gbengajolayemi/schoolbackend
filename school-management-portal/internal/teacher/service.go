package teacher

type Service interface {
    CreateTeacher(teacher *Teacher) (int, error)
    GetTeacher(id int) (*Teacher, error)
    UpdateTeacher(teacher *Teacher) error
    DeleteTeacher(id int) error
    ListTeachers() ([]*Teacher, error)
}

type service struct {
    repo Repository
}

func NewService(repo Repository) Service {
    return &service{repo: repo}
}

func (s *service) CreateTeacher(teacher *Teacher) (int, error) {
    return s.repo.CreateTeacher(teacher)
}

func (s *service) GetTeacher(id int) (*Teacher, error) {
    return s.repo.GetTeacher(id)
}

func (s *service) UpdateTeacher(teacher *Teacher) error {
    return s.repo.UpdateTeacher(teacher)
}

func (s *service) DeleteTeacher(id int) error {
    return s.repo.DeleteTeacher(id)
}

func (s *service) ListTeachers() ([]*Teacher, error) {
    return s.repo.ListTeachers()
}