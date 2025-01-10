package app

import (
	"github.com/osamikoyo/pitition/internal/config"
	"github.com/osamikoyo/pitition/internal/handler"
)

type App struct {
	cfg    config.Config
	handle handler.Handler
}

func Init() *App {
	return &App{}
}
