package storage

import "backend_course/lms/api/models"

type IStorage interface {
	CloseDB()
	StudentStorage() StudentStorage
}

type StudentStorage interface {
	Create(student models.Student) (string, error)
}
