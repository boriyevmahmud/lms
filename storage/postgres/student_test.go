package postgres

import (
	"backend_course/lms/api/models"
	"fmt"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateStudent(t *testing.T) {
	studentRepo := NewStudent(db)

	reqStudent := models.Student{
		FirstName: faker.Name(),
		Age:       10,
		LastName:  faker.Word(),
	}

	id, err := studentRepo.Create(reqStudent)
	if assert.NoError(t, err) {
		createdStudent, err := studentRepo.GetStudent(id)
		if assert.NoError(t, err) {
			assert.Equal(t, reqStudent.FirstName, createdStudent.FirstName)
			assert.Equal(t, reqStudent.Age, createdStudent.Age)
			assert.Equal(t, reqStudent.LastName, createdStudent.LastName)
		} else {
			return
		}
		fmt.Println("Created student", createdStudent)
	}
}
