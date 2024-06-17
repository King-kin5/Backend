package handler

import (
	"Backend/project/Models"
	"net/http"
	"github.com/labstack/echo/v4"
)

func  (h *Handler)UserSignUp(c echo.Context)error{
	// Check database connection
	if err := h.userStore.CheckDBConnection(); err != nil {
		return c.JSON(http.StatusInternalServerError, models.NewResponse(nil, "database connection error: "+err.Error(), false))
	}
	user:=new(models.User)
	if err :=c.Bind(&user);err!= nil {
		return c.JSON(http.StatusBadRequest, models.NewResponse(nil, "bad request", false))
	}
	
	if err:=c.Validate(user);err != nil {
		return c.JSON(http.StatusBadRequest,models.NewResponse(nil,"invalid email",false))
	}
	if err:=models.CheckPasswordLevel(user.Password);err!=nil{
		return c.JSON(http.StatusBadRequest, models.NewResponse(nil, err.Error(), false))
	}
	// Check if the email already exists

	user.Credit = 1000000
	pass, err := models.PasswordHash(user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.NewResponse(nil, "internal server error", false))
	}
	user.Password = pass

	err = h.userStore.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.NewResponse(nil, "failed to create user", false))
	}
	user.Password = ""

	// Create the user response with JWT token
	//userResp := newUserResponse(user)

	// Create a response structure
	response := map[string]interface{}{
		"message": "User created successfully",
		
	}

	return c.JSON(http.StatusCreated, response)

}    


func UserLogin(h *Handler, c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, models.NewResponse(nil,"Bad request",false))
	}

	return c.JSON(http.StatusCreated,models.NewResponse(newUserResponse(user), "", true))
}

          