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
	//TODO implement me
	panic("implement me")
}
