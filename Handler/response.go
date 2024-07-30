package handler

import (
	"Backend/project/Models"
	"Backend/project/Middleware"
	
)
type userResponse struct {
	User struct {
		Email string `json:"email,omitempty"`
		Token       string `json:"token"`
	}
}
func newUserResponse(u *models.User) *userResponse {
	r := new(userResponse)
	r.User.Email=u.Email
	r.User.Token=middleware.GenerateJWT(u.Email)
	return r
}