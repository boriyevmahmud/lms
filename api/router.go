package api

import (
	"backend_course/lms/api/handler"
	"backend_course/lms/service"
	"backend_course/lms/storage"
	"net/http"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// New ...
// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(store storage.IStorage, service service.IServiceManager) *gin.Engine {
	h := handler.NewStrg(store, service)

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Use(authMiddleware)

	r.POST("/student", h.CreateStudent)
	r.PUT("/student/:id", h.UpdateStudent)
	r.PATCH("/student/:id", h.UpdateStudentStatus)
	r.GET("/students", h.GetAllStudents)
	r.DELETE("/student/:id", h.DeleteStudent)
	r.GET("/student/:id", h.GetStudent)

	r.POST("/teacher", h.CreateTeacher)
	r.PUT("/teacher/:id", h.UpdateTeacher)
	r.GET("/teachers", h.GetAllTeachers)
	r.DELETE("/teacher/:id", h.DeleteTeacher)
	r.GET("/teacher/:id", h.GetTeacher)

	return r
}

func authMiddleware(c *gin.Context) {
	if c.GetHeader("Authorization") != "secret" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	c.Next()
}
