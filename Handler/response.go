package handler

import (
	"Backend/project/Models"
)
type userResponse struct {
	User struct {
		Email string `json:"email,omitempty"`
	}
}
func newUserResponse(u *models.User) *userResponse {
	r := new(userResponse)
	
	return r
}