package api

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	db "github.com/flukis/simplebank/db/sqlc"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Server struct {
	store    db.Store
	router   *echo.Echo
	validate *validator.Validate
}

func NewServer(store db.Store, validate *validator.Validate) *Server {
	server := &Server{
		store:    store,
		validate: validate,
	}
	router := echo.New()

	router.POST("/account", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.FetchAccount)

	server.router = router
	return server
}

func (s *Server) Start(addr string) {
	go func() {
		if err := s.router.Start(addr); err != nil && err != http.ErrServerClosed {
			s.router.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.router.Shutdown(ctx); err != nil {
		s.router.Logger.Fatal(err)
	}
}

type Meta struct {
	Limit int32 `json:"limit"`
	Page  int32 `json:"page"`
}
