package server

import (
	"log"
	"spotify-clone/server/config"
	"spotify-clone/server/internal/store"
	"spotify-clone/server/internal/store/sqlstore"

	"github.com/labstack/echo/v4"
)

type Server struct {
	router *echo.Echo
	store  store.Store
}

func New() (*Server, error) {
	dbconfig, err := config.NewDBConfig()
	if err != nil {
		log.Fatal(err)
	}
	s, err := sqlstore.New(dbconfig)
	if err != nil {
		log.Fatal(err)
	}
	e := echo.New()
	return &Server{
		router: e,
		store:  s,
	}, nil
}

func (s *Server) Start(port string) error {
	if err := s.plugRoutes(); err != nil {
		return err
	}
	if err := s.router.Start(":" + port); err != nil {
		return err
	}

	return nil
}

func (s *Server) plugRoutes() error {
	s.router.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})

	s.router.POST("/songs", s.handleSongsCreate())

	s.router.Static("/static/", "../database/")
	return nil
}

func (s *Server) Error(code int, message string, err error, c echo.Context) {
	c.JSON(code, message)
	log.Println(message, err)
}
