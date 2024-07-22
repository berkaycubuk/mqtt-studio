package user

import (
	"fmt"
	"net/http"

	"github.com/berkaycubuk/mqtt-studio/service/auth"
	"github.com/berkaycubuk/mqtt-studio/types"
	"github.com/berkaycubuk/mqtt-studio/utils"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *echo.Group) {
	router.POST("/login", h.handleLogin)
	router.POST("/register", h.handleRegister)
}

func (h *Handler) handleLogin(c echo.Context) error {
	return c.String(http.StatusOK, "Hi!")
}

func (h *Handler) handleRegister(c echo.Context) error {
	// get payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(c, &payload); err != nil {
		utils.WriteError(c, http.StatusBadRequest, err)
		return err
	}

	// check is user exists
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(c, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return fmt.Errorf("user with email %s already exists", payload.Email)
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, err)
		return err
	}

	// create new user
	err = h.store.CreateUser(types.User{
		FullName: payload.FullName,
		Email: payload.Email,
		Password: hashedPassword,
	})
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, err)
		return err
	}

	return c.JSON(http.StatusCreated, nil)
}
