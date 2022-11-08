package repository

import (
	"database/sql"
	"github.com/Hymiside/stubent-media-backend/pkg/models"
	_ "github.com/lib/pq"
)

type Authorization interface {
	CreateSchool(school models.School) (string, error)
	GetSchool(email string) (models.School, error)
}

type Library interface {
	CreateClass(class models.Class) (int, error)
	CreateStudent(student models.Student) (int, error)
	GetAllClasses(schoolId string) ([]models.Class, error)

	// TODO
}

type Repository struct {
	Authorization
	Library
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Library:       NewLibPostgres(db),
	}
}
