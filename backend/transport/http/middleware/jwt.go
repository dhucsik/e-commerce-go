package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dhucsik/e-commerce-go/config"
	"github.com/dhucsik/e-commerce-go/models"
	"github.com/labstack/echo/v4"
)

type JWTAuth struct {
	jwtKey []byte
	jwtTTL int
}

func NewJWTAuth(cfg *config.Config) *JWTAuth {
	return &JWTAuth{
		jwtKey: []byte(cfg.JWTKey),
		jwtTTL: cfg.JWTTTL,
	}
}

func (m *JWTAuth) generateJWT(userID, userRole, userEmail string) (string, error) {
	expirationTime := time.Now().Add(time.Duration(m.jwtTTL) * time.Second)
	claims := models.JWTClaim{
		UserID:    userID,
		UserRole:  userRole,
		UserEmail: userEmail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(m.jwtKey)
}

func (m *JWTAuth) SignInMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			return err
		}

		userData, ok := c.Get(string(models.ContextKey)).(*models.ContextUserData)
		if !ok {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid context user data")
		}

		token, err := m.generateJWT(userData.UserID, userData.UserRole, userData.UserEmail)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, token)
	}
}

func (m *JWTAuth) validateToken(signedToken string) (*models.JWTClaim, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&models.JWTClaim{},
		func(t *jwt.Token) (interface{}, error) {
			return m.jwtKey, nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*models.JWTClaim)
	if !ok {
		return nil, errors.New("couldn't parse claims")
	}
	if claims.ExpiresAt < time.Now().Unix() {
		return nil, errors.New("token expired")
	}

	return claims, nil
}

func (m *JWTAuth) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := extractToken(c.Request())

		claims, err := m.validateToken(token)
		if err != nil {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}

		ctx := context.WithValue(c.Request().Context(), models.ContextKey, &models.ContextUserData{
			UserID:    claims.UserID,
			UserRole:  claims.UserRole,
			UserEmail: claims.UserEmail,
		})

		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}

func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")

	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}

	return ""
}
