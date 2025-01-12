package app

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/osamikoyo/pitition/internal/config"
	"github.com/osamikoyo/pitition/internal/handler"
	"github.com/osamikoyo/pitition/pkg/loger"
	"net/http"
	"time"
)

type App struct {
	handle handler.Handler
	logger loger.Logger
	server *http.Server
}

func Init() *App {
	cfg := config.New()
	logger := loger.New()

	var (
		err               error
		idle, read, write time.Duration
	)

	idle, err = time.ParseDuration(cfg.IdleTimeOut)
	read, err = time.ParseDuration(cfg.ReadTimeOut)
	write, err = time.ParseDuration(cfg.WriteTimeOut)
	if err != nil {
		logger.Error().Err(err)
	}

	return &App{
		handle: handler.New(),
		logger: logger,
		server: &http.Server{
			IdleTimeout:  idle,
			WriteTimeout: write,
			ReadTimeout:  read,
			Addr:         fmt.Sprintf("%s:%d", cfg.Hostname, int(cfg.Port)),
		},
	}
}

func (a *App) Run(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		a.logger.Info().Msg("Server shutdown! :3")
		a.server.Shutdown(ctx)
	}()

	mux := chi.NewRouter()
	a.server.Handler = mux

	a.logger.Info().Msg(fmt.Sprintf("Http Server started on %s ! :3", a.server.Addr))
	return a.server.ListenAndServe()
}
