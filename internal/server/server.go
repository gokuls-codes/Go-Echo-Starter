package server

import (
	"database/sql"
	"net/http"

	"github.com/gokuls-codes/go-echo-starter/internal/services/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	customMiddleware "github.com/gokuls-codes/go-echo-starter/internal/middleware"
)

type Server struct {
	addr string
	db *sql.DB
}

func NewServer(addr string, db *sql.DB) *Server {
	return &Server{
		addr: addr,
		db: db,
	}
}

func (s *Server) Start() error {
	app := echo.New()

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.Static("/static", "assets")

	app.Use(customMiddleware.Theme)

	userGroup := app.Group("/auth")
	userStore := users.NewStore(s.db)
	userHandler := users.NewHandler(userStore)
	userHandler.RegisterRoutes(userGroup)

	homeGroup := app.Group("/")
	homeGroup.Use(customMiddleware.Auth(userStore))

	homeGroup.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})

	return app.Start(s.addr)
}