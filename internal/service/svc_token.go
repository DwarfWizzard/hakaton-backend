package service

import (
	"net/http"
	"time"

	"github.com/DwarfWizzard/hakaton-backend/internal/domain"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

const (
	ACCESS_TOKEN_SECRET  = "Fj&C/,-QD8i%f:9mXva56Z4dTmYVDJ!LCMfuAw.HJ@/Ca7%vux&+]lKyXy6S[#}R"
	REFRESH_TOKEN_SECRET = "3gf2aN):WtQ)/ZepVk@ud&-8?YcxiZ(cu!YDh-E]yf=jcvBnnae#pKh59N@S?3U6"
	ACCESS_TOKEN_TTL     = 15 * time.Minute
	REFRESH_TOKEN_TTL    = 30 * 24 * time.Hour
)

type TokenPair struct {
	Access  string `json:"token"`
	Refresh string `json:"refresh"`
}

type TokenClaims struct {
	jwt.StandardClaims
	UserId uint64 `json:"userId"`
}

type ApiKey struct {
	User *domain.User
}

// MobileRefreshToken - GET /auth/refresh?token=
func (s *Service) MobileRefreshToken(c echo.Context) error {
	tokenString := c.QueryParam("token")

	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("invalid signing method", jwt.ValidationErrorMalformed)
		}

		return []byte(REFRESH_TOKEN_SECRET), nil
	})

	if err != nil {
		s.logger.Error("Parse token error", err)
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return errExpiredJWTToken
			}
		}
		return errIvalidJWTToken
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		s.logger.Error("token claims are not of type *tokenClaims")
		return errIvalidJWTToken
	}

	tokenPair, err := generateTokenPair(claims.UserId)
	if err != nil {
		s.logger.Error("generateTokenPair error", err)
		return errInternalServerError
	}

	return c.JSON(http.StatusOK, &Response{Result: tokenPair})
}

func generateTokenPair(userId uint64) (*TokenPair, error) {
	now := time.Now().UTC()
	access, err := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: now.Add(ACCESS_TOKEN_TTL).Unix(),
			IssuedAt:  now.Unix(),
		},
		userId,
	}).SignedString([]byte(ACCESS_TOKEN_SECRET))

	refresh, err := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: now.Add(REFRESH_TOKEN_TTL).Unix(),
			IssuedAt:  now.Unix(),
		},
		userId,
	}).SignedString([]byte(REFRESH_TOKEN_SECRET))

	if err != nil {
		return nil, err
	}

	return &TokenPair{
		Access:  access,
		Refresh: refresh,
	}, nil
}
