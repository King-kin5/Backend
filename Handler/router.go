package handler

import (
	 "Backend/project/Middleware"

	"github.com/labstack/echo/v4"
)
const (
	URLSignUp = "/signup"
	URLUser   = "/user"
	URLLogin="/login"
	URLProfile="/profile"

)

func (h *Handler) RegisterRoutes(g *echo.Group) {
	g.GET("/", h.BaseRouter)

	//  routes
	g.POST(URLUser+URLSignUp, h.UserSignUp)  // /user/signup
	g.POST(URLUser+URLLogin,h.UserLogin) // /user/login

user := g.Group(URLUser, middleware.USER(middleware.JWTSecret))
user.GET(URLUser+URLProfile, h.Getprofile)
	
}
