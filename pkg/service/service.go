package service

import (
	"github.com/Hymiside/stubent-media-backend/pkg/models"
	"github.com/Hymiside/stubent-media-backend/pkg/repository"
)

type Service struct {
	Authorization
}

type Authorization interface {
	CreateSchool(school models.School) (string, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (string, error)
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
