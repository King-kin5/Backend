package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)// JWTConfig defines the config for JWT middleware.
type JWTConfig struct {
	Skipper    Skipper
	SigningKey interface{}
}

// Skipper defines a function to skip middleware.
type Skipper func(c echo.Context) bool

// jwtExtractor defines a function to extract JWT token from the request.
type jwtExtractor func(c echo.Context) (string, error)

var (
	// ErrJWTMissing is returned when the JWT is missing or malformed.
	ErrJWTMissing = echo.NewHTTPError(http.StatusUnauthorized, "Missing or malformed JWT")
	// ErrJWTInvalid is returned when the JWT is invalid or expired.
	ErrJWTInvalid = echo.NewHTTPError(http.StatusForbidden, "Invalid or expired JWT")
)

var JWTSecret = []byte("!!SECRET!!")

// GenerateJWT generates a JWT token.
func GenerateJWT(email string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, _ := token.SignedString(JWTSecret)
	return t
}

// JWT returns a JWT middleware.
func JWT(key interface{}) echo.MiddlewareFunc {
	c := JWTConfig{}
	c.SigningKey = key
	return JWTWithConfig(c)
}

// JWTWithConfig returns a JWT middleware with config.
func JWTWithConfig(config JWTConfig) echo.MiddlewareFunc {
	extractor := jwtFromHeader("Authorization", "Token")
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth, err := extractor(c)
			if err != nil {
				if config.Skipper != nil && config.Skipper(c) {
					return next(c)
				}
				return c.JSON(http.StatusUnauthorized, NewError(err))
			}
			token, err := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return config.SigningKey, nil
			})
			if err != nil {
				return c.JSON(http.StatusForbidden, NewError(ErrJWTInvalid))
			}
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				userID := uint(claims["id"].(float64))
				c.Set("user", userID)
				return next(c)
			}
			return c.JSON(http.StatusForbidden, NewError(ErrJWTInvalid))
		}
	}
}

// jwtFromHeader returns a jwtExtractor that extracts token from the request header.
func jwtFromHeader(header string, authScheme string) jwtExtractor {
	return func(c echo.Context) (string, error) {
		auth := c.Request().Header.Get(header)
		l := len(authScheme)
		if len(auth) > l+1 && auth[:l] == authScheme {
			return auth[l+1:], nil
		}
		return "", ErrJWTMissing
	}
}
