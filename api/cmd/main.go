package main

import (
	"context"

	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/EgorTarasov/lct-2024/api/internal"
	"github.com/rs/zerolog/log"
)

// @title			lct api
// @version		1.0
// @description	This is a sample swagger for Fiber
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.email	fiber@swagger.io
// @license.name	BSD 3-Clause License
// @license.url	https://raw.githubusercontent.com/EgorTarasov/true-tech/main/LICENSE
// @host			api.lct.larek.tech
// @BasePath		/ //no-lint.
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	go func() {
		if err := internal.Run(ctx, &wg); err != nil {
			panic(err)
		}
		os.Exit(0)
	}()

	waitForExitSignal(cancel, &wg)
}

func waitForExitSignal(cancel context.CancelFunc, wg *sync.WaitGroup) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	log.Info().Msg("SIGINT signal")
	<-sigCh
	log.Info().Msg("\nWaiting tasks for shutdown")
	wg.Wait()
	log.Info().Msg("\nAll tasks finished")
	cancel()
}
