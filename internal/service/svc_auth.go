package service

import (
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (s *Service) Auth(c echo.Context) (*ApiKey, error) {
	auth := c.Request().Header.Get("Authorization")
	if !strings.HasPrefix(auth, "Bearer ") {
		return nil, errNotAuthorized
	}

	tokenString := auth[7:]
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("invalid signing method", jwt.ValidationErrorMalformed)
		}

		return []byte(ACCESS_TOKEN_SECRET), nil
	})

	if err != nil {
		s.logger.Error("Invalid token", err)
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors & jwt.ValidationErrorExpired != 0 {
				return nil, errExpiredJWTToken
			}
		}
		return nil, errIvalidJWTToken
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		s.logger.Error("token claims are not of type *tokenClaims")
		return nil, errIvalidJWTToken
	}

	user, err := s.repo.User.GetUserById(claims.UserId)
	if err != nil {
		s.logger.Error("Get user error", err)
		if err == gorm.ErrRecordNotFound {
			return nil, errUserNotFound
		}
		return nil, errInternalServerError
	}

	return &ApiKey{
		User: user,
	}, nil
}