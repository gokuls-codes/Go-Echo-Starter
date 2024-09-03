package users

import (
	"github.com/gokuls-codes/go-echo-starter/templates/pages"
	"github.com/gokuls-codes/go-echo-starter/types"
	"github.com/gokuls-codes/go-echo-starter/utils"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{
		Store: store,
	}
}

func (h *Handler) RegisterRoutes(group *echo.Group) {
	
	group.GET("/login", func(c echo.Context) error {
		return utils.Render(c, pages.LoginPage(c.Get("theme") == "dark"))
	})

	group.GET("/register", func(c echo.Context) error {
		return utils.Render(c, pages.RegisterPage(c.Get("theme") == "dark"))
	})
}

