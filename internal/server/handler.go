package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	_ "bolao-2026/docs"
	"bolao-2026/internal/config"
	"bolao-2026/internal/utils"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title			Bolao 2026
// @version		1.0
func Listen() (*http.Server, <-chan error) {
	mux := http.NewServeMux()

	registerHandlers(mux)

	port := utils.GetPort()
	address := fmt.Sprintf(":%d", port)

	server := &http.Server{
		Addr:         address,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	errChan := make(chan error, 1)

	go func() {
		log.Printf("Starting server at http://localhost:%d/swagger\n", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errChan <- err
		}
	}()

	return server, errChan
}

func registerHandlers(mux *http.ServeMux) {
	// Swagger UI
	mux.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.UIConfig(map[string]string{
			"displayRequestDuration": "true",
		}),
	))

	// Config
	configService := config.ConfigService{}
	configHandler := config.ConfigHandler{ConfigService: configService}
	configHandler.RegisterRoutes(mux)
}
