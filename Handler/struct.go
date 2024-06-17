package handler

import (
	"Backend/project/Store"
	"errors"
	"net/http"
	"github.com/labstack/echo/v4"
)
type Handler struct {
	userStore            store.Userstore
	
}
func NewHandler(userStore store.Userstore) (handler *Handler) {
return &Handler{
	userStore:            userStore,
	
}
}
func (h *Handler) BaseRouter(c echo.Context) error {
	return c.JSON(http.StatusCreated, errors.New("welcome to Parham food"))
}