package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/lantonster/liberate/cmd/wire"
)

func main() {
	srv := wire.InitializeServer()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	<-sigChan

	if err := srv.Stop(context.Background()); err != nil {
		log.Fatal(err)
	}
}
