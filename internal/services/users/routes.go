package users

import (
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
		cookie, err := c.Cookie(("session"))
		if err == nil {
			log.Println(cookie.Value)
			if auth.CheckIfLoggedIn(cookie.Value, h.store) {
				return c.Redirect(http.StatusSeeOther, "/")
			}
		}
		return utils.Render(c, pages.LoginPage(c.Get("theme") == "dark"))
	})

	group.GET("/register", func(c echo.Context) error {
		cookie, err := c.Cookie(("session"))
		if err == nil {
			log.Println(cookie.Value)
			if auth.CheckIfLoggedIn(cookie.Value, h.store) {
				return c.Redirect(http.StatusSeeOther, "/")
			}
		}
		return utils.Render(c, pages.RegisterPage(c.Get("theme") == "dark"))
	})

	group.POST("/register", h.HandleRegister)
	group.POST("/login", h.HandleLogin)
}

func (h *Handler) HandleRegister(c echo.Context) error {
	p := new(types.RegisterPayload)
	err := c.Bind(p)

	if err != nil {
		c.Response().WriteHeader(http.StatusUnprocessableEntity)
		log.Println(err.Error())
		return utils.Render(c, components.ErrorMessage("Invalid data"))
	}
	
	err = utils.Validate.Struct(p)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		log.Printf("Invalid payload %v", errors)

		c.Response().WriteHeader(http.StatusUnprocessableEntity)
		return utils.Render(c, components.ErrorMessage("Invalid data"))
	}

	_, err = h.store.GetUserByEmail(p.Email) 

	if err == nil {
		log.Printf("Email %v already registered", p.Email)
		c.Response().WriteHeader(http.StatusConflict)
		return utils.Render(c, components.ErrorMessage("Email already exists"))
	}

	hashedPassword, err := auth.HashPassword(p.Password)

	if err != nil {
		log.Println(err.Error())
		c.Response().WriteHeader(http.StatusInternalServerError)
		return utils.Render(c, components.ErrorMessage("Something went wrong"))
	}

	err = h.store.CreateUser(&types.User{
		Name: p.Name,
		Email: p.Email,
		Password: hashedPassword,
		CreatedAt: time.Now(),
	})

	if err != nil {
		log.Println(err.Error())
		c.Response().WriteHeader(http.StatusUnprocessableEntity)
		return utils.Render(c, components.ErrorMessage("Something went wrong"))
	}

	return utils.Render(c, components.SuccessMessage())
}

func (h *Handler) HandleLogin(c echo.Context) error {

	payload:= new(types.LoginPayload)
	err := c.Bind(payload)

	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	err = utils.Validate.Struct(payload)

	if err != nil {
		errors := err.(validator.ValidationErrors)
		log.Printf("Invalid payload %v", errors)

		c.Response().WriteHeader(http.StatusUnprocessableEntity)
		return utils.Render(c, components.ErrorMessage("Invalid data"))
	}

	u, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		log.Println(err.Error())
		c.Response().WriteHeader(http.StatusUnprocessableEntity)
		return utils.Render(c, components.ErrorMessage("Invalid email or password"))
	}

	if !auth.ComparePasswords(u.Password, []byte(payload.Password)) {
		log.Println("Invalid email or password")
		c.Response().WriteHeader(http.StatusUnprocessableEntity)
		return utils.Render(c, components.ErrorMessage("Invalid email or password"))
	}

	sess, err := auth.GenerateSessionCookie(u, h.store)
	if err != nil {
		log.Println(err.Error())
		c.Response().WriteHeader(http.StatusInternalServerError)
		return utils.Render(c, components.ErrorMessage("Unable to login"))
	}

	log.Println(sess)
	cookie := new(http.Cookie)
	cookie.Name = "session"
	cookie.Value = sess.SessionToken
	cookie.Expires = sess.ExpiresAt
	cookie.Path = "/"
	cookie.HttpOnly = true
	c.SetCookie(cookie)

	c.Set("user", u)
	return c.String(200, "Logged in successfully")
}