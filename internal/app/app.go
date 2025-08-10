package app

import (
	"context"
	"fmt"
	"forum/internal/config"
	"forum/internal/repository"
	"forum/internal/repository/postgres"
	"forum/internal/service"
	"log/slog"
	"net/http"
)

type App struct {
	mux     *http.ServeMux
	service *service.Service
	lis     *listeners
}

func NewApp(cfg *config.HTTPServer) (*App, error) {
	db := postgres.MustNew(context.Background(), cfg.PGConfig)

	repo := repository.New(db)

	service := service.NewService(repo)

	lis, err := newListener(cfg)

	if err != nil {
		return nil, fmt.Errorf("can't start listener %w", err)
	}

	app := &App{
		service: service,
		lis:     lis,
	}

	return app, nil
}

func (a *App) Run() {
	a.initHTTP()
	a.runHTTP()
}

func (a *App) initHTTP() {
	handler := New(a.service)
	a.mux = handler.Router()
}

func (a *App) runHTTP() {
	server := &http.Server{
		Handler: a.mux,
	}

	slog.Info("Server is running", slog.String("addr", a.lis.http.Addr().String()))

	server.Serve(a.lis.http)
}
