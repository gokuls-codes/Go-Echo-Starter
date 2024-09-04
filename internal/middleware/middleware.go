package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/gokuls-codes/go-echo-starter/internal/services/auth"
	"github.com/gokuls-codes/go-echo-starter/types"
	"github.com/labstack/echo/v4"
)

func Auth(store types.UserStore) echo.MiddlewareFunc {
	return func (next echo.HandlerFunc) echo.HandlerFunc  {
		return func (c echo.Context) error {
			cookie, err := c.Cookie("session")
			if err != nil {
				log.Println(err.Error())
				return c.Redirect(http.StatusFound, "/auth/login")
			}
			user, loggedIn := auth.CheckIfLoggedIn(cookie.Value, store)
			if !loggedIn {
				return c.Redirect(http.StatusFound, "/auth/login")
			}
			c.Set("user", user)
			return next(c)
		}
	}
}

func Theme(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("theme")
		if err != nil {
			c.Set("theme", "dark")
			cookie := new(http.Cookie)
			cookie.Name = "theme"
			cookie.Value = "dark"
			cookie.Expires = time.Now().Add(24 * time.Hour)
			cookie.Path = "/"
			c.SetCookie(cookie)
			return next(c)
		}
		c.Set("theme", cookie.Value)
		return next(c)
	}
}