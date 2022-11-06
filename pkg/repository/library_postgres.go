package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Hymiside/stubent-media-backend/pkg/models"
)

type LibPostgres struct {
	db *sql.DB
}

func NewLibPostgres(db *sql.DB) *LibPostgres {
	return &LibPostgres{db: db}
}

func (l *LibPostgres) CreateClass(class models.Class) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	request := fmt.Sprintf("insert into classes (school_id, number, letter) values ($1, $2, $3) returning id")
	row := l.db.QueryRowContext(ctx, request, class.SchoolId, class.NumClass, class.LetClass)
	if row.Err() != nil {
		return 0, fmt.Errorf("error to create class: %v", row.Err())
	}

	var classId int
	if err := row.Scan(&classId); err != nil {
		return 0, fmt.Errorf("error to scan classId: %v", err)
	}
	return classId, nil
}

func (l *LibPostgres) CreateStudent(student models.Student) (int, error) {
	//TODO implement me
	panic("implement me")
}
