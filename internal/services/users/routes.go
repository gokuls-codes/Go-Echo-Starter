package users

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gokuls-codes/go-echo-starter/internal/services/auth"
	"github.com/gokuls-codes/go-echo-starter/templates/components"
	"github.com/gokuls-codes/go-echo-starter/templates/pages"
	"github.com/gokuls-codes/go-echo-starter/types"
	"github.com/gokuls-codes/go-echo-starter/utils"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(group *echo.Group) {
	
	group.GET("/login", func(c echo.Context) error {
		return utils.Render(c, pages.LoginPage(c.Get("theme") == "dark"))
	})

	group.GET("/register", func(c echo.Context) error {
		return utils.Render(c, pages.RegisterPage(c.Get("theme") == "dark"))
	})

	group.POST("/register", h.HandleRegister)
	group.POST("/login", h.HandleLogin)
}

func (h *Handler) HandleRegister(c echo.Context) error {
	p := new(types.RegisterPayload)
	err := c.Bind(p)

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	
	err = utils.Validate.Struct(p)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		return c.String(http.StatusBadRequest, fmt.Sprintf("Invalid payload %v", errors))
	}

	_, err = h.store.GetUserByEmail(p.Email) 

	if err == nil {
		return c.String(http.StatusConflict, fmt.Sprintf("User with email %s already exists", p.Email))
	}

	hashedPassword, err := auth.HashPassword(p.Password)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	err = h.store.CreateUser(&types.User{
		Name: p.Name,
		Email: p.Email,
		Password: hashedPassword,
		CreatedAt: time.Now(),
	})

	if err != nil {
		log.Println(err.Error())
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	return utils.Render(c, components.SuccessMessage())
}

func (h *Handler) HandleLogin(c echo.Context) error {
	return c.String(200, "Login")
}