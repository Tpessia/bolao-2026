//go:generate echo "Generating..."

package main

import (
	server "bolao-2026/internal/server"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// use go viper for env files

func main() {
	srv, errChan := server.Listen()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-errChan:
		log.Fatalf("Server error: %v", err)
	case sig := <-sigChan:
		log.Printf("Signal received: %v. Shutting down gracefully...", sig)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("Server forced to shutdown: %v", err)
		}
		log.Println("Server exited cleanly")
	}
}
