package main

import (
	"context"
	"github.com/osamikoyo/pitition/internal/app"
	"github.com/osamikoyo/pitition/pkg/loger"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	if err := app.Init().Run(ctx); err != nil {
		loger.New().Error().Err(err)
	}
}
