package handler

import (
	"Backend/project/Models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) UserSignUp(c echo.Context) error {
	// Check database connection
	if err := h.userStore.CheckDBConnection(); err != nil {
		return c.JSON(http.StatusInternalServerError, models.NewResponse(nil, "database connection error: "+err.Error(), false))
	}

	user := new(models.User)
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, models.NewResponse(nil, "bad request", false))
	}

	if user.Email == "" {
		return c.JSON(http.StatusBadRequest, models.NewResponse(nil, "email must not be empty", false))
	}

	if err := c.Validate(user); err != nil {
		return c.JSON(http.StatusBadRequest, models.NewResponse(nil, "invalid email", false))
	}

	if err := models.CheckPasswordLevel(user.Password); err != nil {
		return c.JSON(http.StatusBadRequest, models.NewResponse(nil, err.Error(), false))
	}

	existingUserByName, _ := h.userStore.GetUserByName(user.Name)
	if existingUserByName != nil {
		return c.JSON(http.StatusBadRequest, models.NewResponse(nil, "this name is already registered", false))
	}

	// Check if the email already exists
	existingUser, _ := h.userStore.GetUserByEmail(user.Email)
	if existingUser != nil {
		return c.JSON(http.StatusBadRequest, models.NewResponse(nil, "this user already exists", false))
	}

	//existingUserByPhone, _ := h.userStore.GetUserByPhoneNumber(user.PhoneNumber)
	//if existingUserByPhone != nil {
	//	return c.JSON(http.StatusBadRequest, models.NewResponse(nil, "this phone number is already registered", false))
	//}

	user.Credit = 1000000
	pass, err := models.PasswordHash(user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.NewResponse(nil, "internal server error", false))
	}
	user.Password = pass

	err = h.userStore.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.NewResponse(nil, "failed to create user: "+err.Error(), false))
	}
	user.Password = ""

	// Generate OTP
	otp := generateOTP()
	log.Printf("Generated OTP for user %s:"  ,otp)

	// Send OTP via email


	// Create the user response with JWT token
	userResp := newUserResponse(user)

	// Create a response structure
	response := map[string]interface{}{
		"message": "user created successfully",
		"user":    userResp,
	}

	return c.JSON(http.StatusCreated, response)
}
func (h *Handler) UserLogin(c echo.Context) error {
    log.Println("UserLogin: Received login request")

    loginRequest := new(struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    })

    if err := c.Bind(loginRequest); err != nil {
        log.Printf("UserLogin: Failed to bind request: %v\n", err)
        return c.JSON(http.StatusBadRequest, models.NewResponse(nil, "bad request", false))
    }

    user, err := h.userStore.GetUserByEmail(loginRequest.Email)
    if err != nil {
        log.Printf("UserLogin: Failed to get user by email: %v\n", err)
        return c.JSON(http.StatusInternalServerError, models.NewResponse(nil, "internal server error", false))
    }
    if user == nil {
        log.Println("UserLogin: Invalid email or password")
        return c.JSON(http.StatusUnauthorized, models.NewResponse(nil, "email or password is incorrect", false))
    }

    if !models.CheckPasswordSame(user.Password, loginRequest.Password) {
        log.Println("UserLogin: Invalid  password")
        return c.JSON(http.StatusUnauthorized, models.NewResponse(nil, "email or password is incorrect", false))
    }

    log.Println("UserLogin: Login successful for user:", user.Email)
    user.Password = ""

    return c.JSON(http.StatusOK, models.NewResponse(newUserResponse(user), "login successful", true))
}
func (h *Handler) Getprofile (c echo.Context) error {
	name:=c.Param("username")
	user,err:=h.userStore.GetUserByName(name)
	 if err!=nil{
		log.Println("Failed to find by name")
		return c.JSON(http.StatusInternalServerError,models.NewResponse(nil, "internal server error", false))
	 }
	 if user==nil {
		log.Printf("Username not found:%v",user)
		return c.JSON(http.StatusNotFound,models.NewResponse(nil,"Username not found",false))
	 }
	 return c.JSON(http.StatusOK,models.NewResponse(newUserResponse(user),"User name found",true))
}