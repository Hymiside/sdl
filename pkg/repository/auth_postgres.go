package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Hymiside/stubent-media-backend/pkg/models"
	"log"
	"time"
)

type AuthPostgres struct {
	db *sql.DB
}

func NewAuthPostgres(db *sql.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

// CreateSchool пишет в БД новую школу
func (r *AuthPostgres) CreateSchool(school models.School) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stmt, err := r.db.Prepare("insert into users(id, firstname, lastname, username, email, password_hash) values ($1, $2, $3, $4, $5, $6) returning id")
	if err != nil {
		return "", fmt.Errorf("error to prepare sql request: %v", err)
	}

	rows, err := stmt.QueryContext(ctx, school.Id, school.Name, school.PhoneNumber, school.Email, school.Password)
	if err != nil {
		return "", fmt.Errorf("error to create school: %v", err)
	}

	var schoolId string
	for rows.Next() {
		if err = rows.Scan(&schoolId); err != nil {
			return "", fmt.Errorf("error to parse schoolId: %v", err)
		}
	}

	// Тут бы какое-то логирование сделать
	if err = stmt.Close(); err != nil {
		log.Printf("error to close stmt: %v", err)
	}
	return schoolId, nil
}

func (r *AuthPostgres) GetSchool(username, password string) (models.School, error) {
	//TODO implement me
	panic("implement me")
}
