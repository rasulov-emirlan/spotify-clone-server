package server

import (
	"context"
	"log"
	"spotify-clone/server/config"
	_ "spotify-clone/server/docs"
	"spotify-clone/server/internal/fs"
	"spotify-clone/server/internal/fs/googlefs"
	"spotify-clone/server/internal/models"
	"spotify-clone/server/internal/server/handlers"
	"spotify-clone/server/internal/store"
	"spotify-clone/server/internal/store/sqlstore"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Spotify Clone Server
// @version 1.1
// @description This is a backend server for spotify clone.

// @contact.name Rasulov Emirlan
// @contact.email rasulov-emirlan@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @Accept json
type Server struct {
	Router *echo.Echo
	Store  store.Store
	FS     fs.FileSystem
	JWTkey string
}

// New() is our constructor for server
// here we specify everything that needs to be done before staring the server
func New() (*Server, error) {
	dbconfig, err := config.NewSQLDBlink()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(dbconfig)
	jwtkey, err := config.NewJWTToken()
	if err != nil {
		log.Fatal(err)
	}
	s, err := sqlstore.New(dbconfig)
	if err != nil {
		log.Fatal(err)
	}
	fs, err := googlefs.NewFileSystem()
	if err != nil {
		log.Fatal(err)
	}
	e := echo.New()
	return &Server{
		Router: e,
		Store:  s,
		FS:     fs,
		JWTkey: jwtkey,
	}, nil
}

func (s *Server) Start(port string) error {
	if err := s.plugRoutes(); err != nil {
		return err
	}
	if err := s.Router.Start(":" + port); err != nil {
		return err
	}

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.Router.Shutdown(ctx)
}

func (s *Server) plugRoutes() error {
	s.Router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderAccessControlAllowOrigin},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	s.Router.Use(middleware.RemoveTrailingSlash())
	s.Router.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
	}))
	s.Router.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "swagger")
		},
	}))

	jwtmiddleware := middleware.JWT([]byte(s.JWTkey))

	s.Router.GET("/ping", func(c echo.Context) error {
		return c.JSON(200, "pong")
	})

	songs := s.Router.Group("/songs")
	{
		songs.POST("", handlers.SongsCreate(s.Store, s.FS), jwtmiddleware, handlers.MiddlewareCheckRole("singer"))
		songs.GET("", handlers.SongsFromAtoB(s.Store))
	}

	playlists := s.Router.Group("/playlists")
	{
		playlists.POST("", handlers.PlaylistsCreate(s.Store, s.FS), jwtmiddleware)
		playlists.GET("", handlers.ListAllPlaylists(s.Store))
		playlists.GET("/:id", handlers.GetSongsFromPlaylist(s.Store))
		playlists.PUT("", handlers.PlaylistsAddSong(s.Store), jwtmiddleware)
	}

	genres := s.Router.Group("/genres")
	{
		genres.POST("", handlers.GenresCreate(s.Store, s.FS), jwtmiddleware, handlers.MiddlewareCheckRole(models.Admin))
		genres.PATCH("", handlers.GenresAddLocalization(s.Store), jwtmiddleware, handlers.MiddlewareCheckRole(models.Admin))
		genres.PUT("", handlers.GenresAddSong(s.Store), jwtmiddleware, handlers.MiddlewareCheckRole(models.Admin, models.Singer))
		genres.GET("", handlers.ListAllGenres(s.Store))
		genres.GET("/:genre", handlers.GenresSongs(s.Store))
	}

	languages := s.Router.Group("/languages")
	{
		languages.POST("", handlers.LanguagesCreate(s.Store), jwtmiddleware, handlers.MiddlewareCheckRole(models.Admin))
		languages.GET("", handlers.LanguagesListAll(s.Store))
	}

	countries := s.Router.Group("/countries")
	{
		countries.POST("", handlers.CountriesCreate(s.Store), jwtmiddleware, handlers.MiddlewareCheckRole(models.Admin))
		countries.GET("", handlers.CountriesListAll(s.Store))
		countries.DELETE("/:id", handlers.CountriesDelete(s.Store), jwtmiddleware, handlers.MiddlewareCheckRole(models.Admin))
	}

	users := s.Router.Group("/users")
	{
		users.POST("/favorite/songs/:id", handlers.UsersAddFavoriteSong(s.Store), jwtmiddleware)
		users.GET("/favorite/songs/", handlers.UsersListFavoriteSongs(s.Store))
		users.DELETE("/favorite/songs/:id", handlers.UsersRemoveFavoriteSong(s.Store), jwtmiddleware)
	}

	auth := s.Router.Group("/auth")
	{
		auth.POST("/register", handlers.UserRegistration(s.Store, s.JWTkey))
		auth.POST("/login", handlers.UserLogin(s.Store, s.JWTkey))
	}

	s.Router.GET("/swagger/*", echoSwagger.WrapHandler)

	return nil
}
