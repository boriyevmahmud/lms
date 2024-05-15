package service

import (
	"backend_course/lms/api/models"
	"backend_course/lms/config"
	"backend_course/lms/pkg"
	"backend_course/lms/pkg/jwt"
	"backend_course/lms/pkg/smtp"
	"backend_course/lms/storage"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

type authService struct {
	storage storage.IStorage
}

func NewAuthService(storage storage.IStorage) authService {
	return authService{storage: storage}
}

func (s authService) TeacherLogin(ctx context.Context, req models.LoginRequest) (models.LoginResponse, error) {
	resp := models.LoginResponse{}

	teacher, err := s.storage.TeacherStorage().GetTeacherByLogin(ctx, req.Login)
	if err != nil {
		// log
		return resp, err
	}

	if err = pkg.CompareHashAndPassword(teacher.Password, req.Password); err != nil {
		return resp, errors.New("password doesn't match")
	}

	m := make(map[interface{}]interface{})
	m["user_id"] = teacher.Id
	m["user_role"] = config.TEACHER_TYPE
	accessToken, refreshToken, err := jwt.GenJWT(m)
	if err != nil {
		return resp, err
	}
	resp.AccessToken = accessToken
	resp.RefreshToken = refreshToken

	return resp, nil
}

func (s authService) TeacherRegister(ctx context.Context, req models.RegisterRequest) error {

	_, err := s.storage.TeacherStorage().GetTeacherByLogin(ctx, req.Mail)
	if err == pgx.ErrNoRows {
		otp := pkg.GenerateOTP()
		msg := fmt.Sprintf("Your OTP: %v. DON'T give anyone", otp)
		err := s.storage.Redis().SetX(ctx, req.Mail, otp, time.Minute*2)
		if err != nil {
			return err
		}

		err = smtp.SendMail(req.Mail, msg)
		if err != nil {
			return err
		}

	} else if err != nil {
		return err
	} else {
		return errors.New("email already exists")
	}

	return nil
}
