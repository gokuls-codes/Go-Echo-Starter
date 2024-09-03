package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("session")
		if err != nil {
			return c.String(401, "Unauthorized")
		}
		fmt.Println(cookie.Value)
		return next(c)
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
			c.SetCookie(cookie)
			return next(c)

		}
		c.Set("theme", cookie.Value)
		return next(c)
	}
}