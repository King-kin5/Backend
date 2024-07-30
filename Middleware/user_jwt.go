package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)
type(
	//USERJWTConfig defines the config for UserJWT middleware.
	USERJWTConfig struct {
		Skipper    Skipper
		SigningKey interface{}
	}
	// Skipper defines a function to skip middleware.
 UserSkipper func(c echo.Context) bool
// jwtExtractor defines a function to extract JWT token from the request.
  UserjwtExtractor func(c echo.Context) (string, error)

)
var(
	USERErrJWTInvalid = echo.NewHTTPError(http.StatusForbidden, "invalid or expired jwt")

)
func USER (key interface{})echo.MiddlewareFunc{
	c:=JWTConfig{}
	c.SigningKey=key
	return USERJWTFROMHEADER(c)
}
//  returns a JWT middleware with config.
func USERJWTFROMHEADER (config JWTConfig)echo.MiddlewareFunc{
	 extractor:=jwtFromHeader("Authorzation","Token")
	 return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth,err:=extractor(c)
			if auth ==""&& config.Skipper != nil &&config.Skipper(c) {
				return next(c)
			}
			if err!=nil {
				return c.JSON(http.StatusForbidden,USERErrJWTInvalid)
			}
			if auth ==""{
				 if config.Skipper!=nil{
					return next(c)
				 }
				 return c.JSON(http.StatusUnauthorized,NewError(errors.New("Missing or malformed jwt")))
			}
			token, err := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return config.SigningKey, nil
			})
			if err != nil {
				return c.JSON(http.StatusForbidden, NewError(USERErrJWTInvalid))
			}
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				email := claims["email"]
				c.Set("email", email)
				return next(c)
			}
			return c.JSON(http.StatusForbidden,NewError(USERErrJWTInvalid))
		}
	 }
}