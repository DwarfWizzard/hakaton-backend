package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	errNotAuthorized       = echo.NewHTTPError(http.StatusUnauthorized, "not authorized")
	errNotAuthenticated    = echo.NewHTTPError(http.StatusUnauthorized, "not authenticated")
	errUserNotFound        = echo.NewHTTPError(http.StatusBadRequest, "user not found")
	errInternalServerError = echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	errExpiredJWTToken     = echo.NewHTTPError(http.StatusUnauthorized, "jwt token expired")
	errIvalidJWTToken      = echo.NewHTTPError(http.StatusBadRequest, "jwt invalid token")
	errIvalidRefreshToken  = echo.NewHTTPError(http.StatusBadRequest, "refresh invalid token")
	errInvalidRequest      = echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	errInvalidInput        = echo.NewHTTPError(http.StatusBadRequest, "invalid input")
	errInvalidParamType    = echo.NewHTTPError(http.StatusBadRequest, "invalid type of query param")
	errEmptyRequiredParam  = echo.NewHTTPError(http.StatusBadRequest, "requiered param can`t be empty")
)

type Error struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

func (e *Error) SendResponse(c echo.Context) error {
	if e == nil || c.Response().Committed {
		return nil
	}

	return c.JSON(e.Code, &Response{Error: e})
}

func (s *Service) HttpErrorHandler(err error, c echo.Context) {
	if err == nil {
		return
	}

	var eErr *Error
	switch e := err.(type) {
	case *echo.HTTPError:
		eErr = &Error{
			Code:        e.Code,
			Description: e.Message.(string),
		}
	default:
		eErr = &Error{
			Code:        -1,
			Description: e.Error(),
		}
	}

	if c.Response().Committed {
		s.logger.Error("Response commited. Cant send error response", err)
		return
	}

	respErr := eErr.SendResponse(c)

	if respErr != nil {
		s.logger.Error("Cant send error response", err)
	}
}
