package postgres

import (
	"backend_course/lms/api/models"
	"database/sql"

	"github.com/google/uuid"
)

type studentRepo struct {
	db *sql.DB
}

func NewStudent(db *sql.DB) studentRepo {
	return studentRepo{
		db: db,
	}
}

func (s *studentRepo) Create(student models.Student) (string, error) {

	id := uuid.New()

	query := ` INSERT INTO students (id, first_name) VALUES ($1, $2) `

	_, err := s.db.Exec(query, id, student.FirstName)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

//exec
//query
//queryrow
