package server

import (
	"context"
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

// New() is our constructor for server
// here we specify everything that needs to be done before staring the server
func New() (*Server, error) {
	dbconfig, err := config.NewSQLDBconfig()
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

func (s *Server) Shutdown(ctx context.Context) error {
	return s.router.Shutdown(ctx)
}

func (s *Server) plugRoutes() error {
	// just for testing purposes
	s.router.GET("/ping", func(c echo.Context) error {
		return c.JSON(200, "pong")
	})

	// the main REST API
	s.router.POST("/songs", s.handleSongsCreate())
	s.router.POST("/users", s.handleUserRegistration())
	// the main REST API

	// this part serves our files withing ../database/ folder
	s.router.Static("/static/", "../database/")
	return nil
}

type Response map[string]interface {
}

func (s *Server) Error(code int, message string, err error, c echo.Context) {
	c.JSON(code, Response{"message": message})
	log.Println(message, err)
}
