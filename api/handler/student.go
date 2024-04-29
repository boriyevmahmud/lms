package handler

import (
	"backend_course/lms/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) CreateStudent(c *gin.Context) {
	student := models.Student{}

	if err := c.ShouldBindJSON(&student); err != nil {
		handleResponse(c, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}
	// if err := check.ValidateCarYear(car.Year); err != nil {
	// 	handleResponse(c, "error while validating car year, year: "+strconv.Itoa(car.Year), http.StatusBadRequest, err.Error())

	// 	return
	// }

	id, err := h.Store.StudentStorage().Create(student)
	if err != nil {
		handleResponse(c, "error while creating car", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, "Created successfully", http.StatusOK, id)
}
