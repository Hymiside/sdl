package service

import (
	"fmt"
	"time"

	"github.com/Hymiside/stubent-media-backend/pkg/models"
	"github.com/Hymiside/stubent-media-backend/pkg/repository"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	signingKey = []byte("qrkjk#4#%35FSFJlja#4353KSFjH")
	tokenTTL   = 24 * time.Hour
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateSchool(school models.School) (string, error) {
	var schoolId string

	school.Id = uuid.New().String()
	hash, err := hashPassword(school.Password)
	if err != nil {
		return "", fmt.Errorf("error to hashing password")
	}
	school.Password = hash

	schoolId, err = s.repo.CreateSchool(school)
	if err != nil {
		return "", err
	}

	return schoolId, nil
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	school, err := s.repo.GetSchool(email)
	if err != nil {
		return "", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(school.Password), []byte(password)); err != nil {
		return "", fmt.Errorf("invalid password: %v", err)
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["ExpiresAt"] = time.Now().Add(10 * time.Second)
	claims["SchoolId"] = school.Id

	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", fmt.Errorf("error to create jwt-token: %v", err)
	}
	return tokenString, nil
}

func (s *AuthService) ParseToken(token string) (int, error) {
	//TODO implement me
	panic("implement me")
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
