package service

import (
	"fmt"
	"github.com/Hymiside/stubent-media-backend/pkg/models"
	"github.com/Hymiside/stubent-media-backend/pkg/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.School) (string, error) {
	var userId string

	user.Id = uuid.New().String()
	hash, err := hashPassword(user.Password)
	if err != nil {
		return "", fmt.Errorf("error to hashing password")
	}
	user.Password = hash

	userId, err = s.repo.CreateSchool(user)
	if err != nil {
		return "", err
	}

	return userId, nil
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (s *AuthService) ParseToken(token string) (int, error) {
	//TODO implement me
	panic("implement me")
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
