package main

import (
	"context"
	"github.com/alitvinenko/ecareer_bot/internal/app"
	"log"
)

func main() {
	ctx := context.Background()
	application := app.NewDaemon(ctx)

	err := application.Run(ctx)
	if err != nil {
		log.Fatalf("error on run application: %v", err)
	}
}
