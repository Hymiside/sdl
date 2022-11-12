package service

import (
	"github.com/Hymiside/stubent-media-backend/pkg/models"
	"github.com/Hymiside/stubent-media-backend/pkg/repository"
)

type StudNClsService struct {
	repo repository.StudNCls
}

func NewStudNClsService(repo repository.StudNCls) *StudNClsService {
	return &StudNClsService{repo: repo}
}

func (s *StudNClsService) CreateClass(class models.Class) (int, error) {
	classId, err := s.repo.CreateClass(class)
	if err != nil {
		return 0, err
	}
	return classId, nil
}

func (s *StudNClsService) CreateStudent(student models.Student) (int, error) {
	studentId, err := s.repo.CreateStudent(student)
	if err != nil {
		return 0, err
	}
	return studentId, nil
}

func (s *StudNClsService) GetAllClasses(schoolId string) ([]models.Class, error) {
	classes, err := s.repo.GetAllClasses(schoolId)
	if err != nil {
		return nil, err
	}
	return classes, nil
}

func (s *StudNClsService) GetAllStudents(schoolId string) ([]models.Student, error) {
	students, err := s.repo.GetAllStudents(schoolId)
	if err != nil {
		return nil, err
	}
	return students, nil
}

func (s *StudNClsService) DeleteStudent(email string) error {
	//TODO implement me
	panic("implement me")
}
