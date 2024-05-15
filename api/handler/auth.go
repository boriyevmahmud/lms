package handler

import (
	"backend_course/lms/api/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TeacherLogin godoc
// @Router       /teacher/login [POST]
// @Summary      Teacher login
// @Description  Teacher login
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        login body models.LoginRequest true "login"
// @Success      201  {object}  models.LoginResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *Handler) TeacherLogin(c *gin.Context) {
	loginReq := models.LoginRequest{}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		handleResponse(c, h.Log, "error while binding body", http.StatusBadRequest, err)
		return
	}
	fmt.Println("loginReq: ", loginReq)

	//TODO: need validate login & password

	loginResp, err := h.Service.Auth().TeacherLogin(c.Request.Context(), loginReq)
	if err != nil {
		handleResponse(c, h.Log, "unauthorized", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.Log, "Succes", http.StatusOK, loginResp)

}

// TeacherRegister godoc
// @Router       /teacher/register [POST]
// @Summary      Teacher register
// @Description  Teacher register
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        register body models.RegisterRequest true "register"
// @Success      201  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *Handler) TeacherRegister(c *gin.Context) {
	loginReq := models.RegisterRequest{}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		handleResponse(c, h.Log, "error while binding body", http.StatusBadRequest, err)
		return
	}

	//TODO: need validate login

	err := h.Service.Auth().TeacherRegister(c.Request.Context(), loginReq)
	if err != nil {
		handleResponse(c, h.Log, "Bad request", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "Succes", http.StatusOK, "Your request succeed")

}
