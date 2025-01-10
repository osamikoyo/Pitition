package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/osamikoyo/pitition/internal/service"
	"github.com/osamikoyo/pitition/pkg/loger"
	"net/http"
)

type Handler struct {
	Pit service.Pititioner
}

func New() Handler {
	return Handler{
		Pit: service.New(),
	}
}

type handlerFunc func(w http.ResponseWriter, r *http.Request) error

func handler(hf handlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if err := hf(writer, request); err != nil {
			loger.New().Error().Err(err)
		}
	}
}

func (h Handler) ping(w http.ResponseWriter, r *http.Request) error {
	_, err := w.Write([]byte("PONG"))
	return err
}

func (h Handler) RegisterRoutes(mux *chi.Mux) {
	mux.Use(middleware.Logger)

	mux.Get("/ping", handler(h.ping))
}
