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

type StudNCls interface {
	CreateClass(class models.Class) (int, error)
	CreateStudent(student models.Student) (int, error)
	GetAllClasses(schoolId string) ([]models.Class, error)
	GetAllStudents(schoolId string) ([]models.Student, error)
	DeleteStudent(email string) error

	// TODO
}

type Repository struct {
	Authorization
	StudNCls
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		StudNCls:      NewStudNClsPostgres(db),
	}
}
