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
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	request := fmt.Sprintf("select id from classes where school_id = $1 and number = $2 and letter = $3")
	row := l.db.QueryRowContext(ctx, request, student.SchoolId, student.ClassNum, student.ClassLet)
	if row.Err() != nil {
		return 0, fmt.Errorf("error to get classId: %v", row.Err())
	}
	var classId int
	if err := row.Scan(&classId); err != nil {
		return 0, fmt.Errorf("error to scan classId for student: %v", err)
	}

	request1 := fmt.Sprintf("insert into students (first_name, last_name, middle_name, class_id, school_id, email, phone_number) values ($1, $2, $3, $4, $5, $6, $7) returning id")
	row1 := l.db.QueryRowContext(ctx, request1, student.FirstName, student.LastName, student.MiddleName, classId, student.SchoolId, student.Email, student.PhoneNumber)
	if row.Err() != nil {
		return 0, fmt.Errorf("error to create student: %v", row.Err())
	}

	var studentId int
	if err := row1.Scan(&studentId); err != nil {
		return 0, fmt.Errorf("error to scan studetId: %v", err)
	}
	return studentId, nil
}

func (l *LibPostgres) GetAllClasses(schoolId string) ([]models.Class, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	request := fmt.Sprintf("select number, letter from classes where school_id = $1")
	rows, err := l.db.QueryContext(ctx, request, schoolId)
	if err != nil {
		return nil, fmt.Errorf("error to get all classes: %v", err)
	}

	var classes []models.Class
	for rows.Next() {
		var class models.Class
		if err = rows.Scan(&class.NumClass, &class.LetClass); err != nil {
			return nil, fmt.Errorf("error to scan classes: %v", err)
		}
		classes = append(classes, class)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("rows error to get classes")
	}
	return classes, nil
}

func (l *LibPostgres) GetAllStudents(schoolId string) ([]models.Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	request := fmt.Sprintf("select students.first_name, students.last_name, students.middle_name, classes.letter, classes.number, students.email, students.phone_number from students join classes on classes.id = students.class_id and students.school_id = $1")
	rows, err := l.db.QueryContext(ctx, request, schoolId)
	if err != nil {
		return nil, fmt.Errorf("error to get all students: %v", err)
	}

	var students []models.Student
	for rows.Next() {
		var student models.Student
		if err = rows.Scan(&student.FirstName, &student.LastName, &student.MiddleName, &student.ClassLet, &student.ClassNum, &student.Email, &student.PhoneNumber); err != nil {
			return nil, fmt.Errorf("error to scan students: %v", err)
		}
		students = append(students, student)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("rows error to get classes")
	}
	return students, nil
}
