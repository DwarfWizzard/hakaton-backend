package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type MobileLoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// MobileLogin - POST /auth/login
func (s *Service) MobileLogin(c echo.Context) error {
	rq := &MobileLoginRequest{}

	if err := c.Bind(&rq); err != nil {
		s.logger.Error("Binding error", err)
		return errInvalidInput
	}

	if len(rq.Login) == 0 || len(rq.Password) == 0 {
		return errInvalidInput
	}

	user, err := s.repo.GetUserByLogin(rq.Login)
	if err != nil {
		s.logger.Error("Get user error", err)
		if err == gorm.ErrRecordNotFound {
			return errNotAuthorized
		}
		return errInternalServerError
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(rq.Password)); err != nil {
		s.logger.Error(err)
		return errNotAuthorized
	}

	tokenPair, err := generateTokenPair(user.Id)
	if err != nil {
		s.logger.Error("Generate token pair error", err)
		return errInternalServerError
	}

	return c.JSON(http.StatusOK, &Response{
		Result: tokenPair,
	})
}
