package service

import (
	"github.com/Hymiside/stubent-media-backend/pkg/models"
	"github.com/Hymiside/stubent-media-backend/pkg/repository"
)

type LibService struct {
	repo repository.Library
}

func NewLibService(repo repository.Library) *LibService {
	return &LibService{repo: repo}
}

func (l *LibService) CreateClass(class models.Class) (int, error) {
	classId, err := l.repo.CreateClass(class)
	if err != nil {
		return 0, err
	}
	return classId, nil
}

func (l *LibService) CreateStudent(student models.Student) (int, error) {
	studentId, err := l.repo.CreateStudent(student)
	if err != nil {
		return 0, err
	}
	return studentId, nil
}

func (l *LibService) GetAllClasses(schoolId string) ([]models.Class, error) {
	classes, err := l.repo.GetAllClasses(schoolId)
	if err != nil {
		return nil, err
	}
	return classes, nil
}

func (l *LibService) GetAllStudents(schoolId string) ([]models.Student, error) {
	students, err := l.repo.GetAllStudents(schoolId)
	if err != nil {
		return nil, err
	}
	return students, nil
}
