package service

import (
	"github.com/Hymiside/stubent-media-backend/pkg/models"
	"github.com/Hymiside/stubent-media-backend/pkg/repository"
)

type Service struct {
	Authorization
	Library
}

type Authorization interface {
	CreateSchool(school models.School) (string, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (string, error)
}

type Library interface {
	CreateClass(class models.Class) (int, error)
	CreateStudent(student models.Student) (int, error)

	// TODO
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Library:       NewLibService(repos.Library),
	}
}
