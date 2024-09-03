package server

import (
	"database/sql"

	"github.com/labstack/echo/v4"
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
	app.Static("/static", "assets")

	return app.Start(s.addr)
}