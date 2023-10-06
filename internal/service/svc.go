package service

import (
	"github.com/DwarfWizzard/hakaton-backend/internal/repository"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Service struct {
	logger *logrus.Logger
	repo   *repository.Repository
}

func New(logger *logrus.Logger, repo *repository.Repository) *Service {
	return &Service{
		logger: logger,
		repo:   repo,
	}
}

func (s *Service) BindApi(endpoint func(*Service, echo.Context) error) echo.HandlerFunc {
	return func(c echo.Context) error {
		return endpoint(s, c)
	}

}

func (s *Service) BindApiWithAuth(apifunc func(*Service, echo.Context, *ApiKey) error) echo.HandlerFunc {
	return func(c echo.Context) error {
		apikey, err := s.Auth(c)
		if err != nil {
			s.logger.Error("Auth error:", err)
			if err == errInternalServerError {
				return errInternalServerError
			}
			return errNotAuthenticated
		}
		return apifunc(s, c, apikey)
	}
}
