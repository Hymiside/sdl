package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Hymiside/stubent-media-backend/pkg/models"
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

	stmt, err := r.db.Prepare("insert into schools(id, name, phone_number, email, password_hash) values ($1, $2, $3, $4, $5) returning id")
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

func (r *AuthPostgres) GetSchool(email string) (models.School, error) {
	var school models.School

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	request := fmt.Sprintf("select id, password_hash from schools where email = $1")
	row := r.db.QueryRowContext(ctx, request, email)
	if row.Err() != nil {
		return models.School{}, fmt.Errorf("error to get schoolId and passwordHash: %v", row.Err())
	}
	if err := row.Scan(&school.Id, &school.Password); err != nil {
		return models.School{}, fmt.Errorf("error to parse schoolId and passwordHash: %v", err)
	}

	return school, nil
}
