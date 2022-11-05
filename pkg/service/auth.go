package service

import (
	"errors"
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

type Claims struct {
	jwt.StandardClaims
	SchoolId string
}

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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		school.Id,
	})

	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", fmt.Errorf("error to create jwt-token: %v", err)
	}
	return tokenString, nil
}

func (s *AuthService) ParseToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return signingKey, nil
	})
	if err != nil {
		return "", fmt.Errorf("error to parse jwt-token: %v", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}
	return claims.SchoolId, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
