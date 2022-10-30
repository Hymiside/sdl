package repository

import (
	"database/sql"
	"github.com/Hymiside/stubent-media-backend/pkg/models"
	_ "github.com/lib/pq"
)

type Authorization interface {
	CreateSchool(user models.School) (string, error)
	GetSchool(username, password string) (models.School, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
