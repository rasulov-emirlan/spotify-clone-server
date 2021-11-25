package server

import (
	"context"
	"log"
	"spotify-clone/server/config"
	"spotify-clone/server/internal/store"
	"spotify-clone/server/internal/store/sqlstore"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	router *echo.Echo
	store  store.Store
	jwtkey string
}

// New() is our constructor for server
// here we specify everything that needs to be done before staring the server
func New() (*Server, error) {
	dbconfig, err := config.NewSQLDBconfig()
	if err != nil {
		log.Fatal(err)
	}
	jwtkey, err := config.NewJWTToken()
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
		jwtkey: jwtkey,
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
	s.router.Use(middleware.CORS())

	s.router.GET("/ping", func(c echo.Context) error {
		return c.JSON(200, "pong")
	})
	songs := s.router.Group("/songs")
	{
		songs.Use(middleware.JWT([]byte(s.jwtkey)))
		songs.POST("/", s.handleSongsCreate())
	}

	playlists := s.router.Group("/playlists")
	{
		playlists.Use(middleware.JWT([]byte(s.jwtkey)))
		playlists.POST("/", s.handlePlaylistsCreate())
		playlists.GET("/:playlist_id", s.handlePlaylistsGetSongsFromPlaylist())
		playlists.PATCH("", s.handlePlaylistsAddSong())
	}

	auth := s.router.Group("/auth")
	{
		auth.POST("/register", s.handleUserRegistration())
		auth.POST("/login", s.handleUserLogin())
	}

	s.router.Static("/database/", "../database/")
	return nil
}

type response map[string]interface {
}

func (s *Server) Error(code int, message string, err error, c echo.Context) {
	c.JSON(code, response{"message": message})
	log.Println(message, err)
}
