package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err.Error())
		os.Exit(1)
	}
}

func run() error {
	var handler http.Handler
	server := &http.Server{
		Addr:    os.Getenv("PORT"),
		Handler: handler,
	}

	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	serverErrors := make(chan error, 1)
	go func() {
		log.Println("startup portdomainapi", server.Addr)
		serverErrors <- server.ListenAndServe()
	}()

	select {
	case serverError := <-serverErrors:
		return errors.Unwrap(serverError)

	case sig := <-quit:
		log.Println("Server is shutting down", sig)

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if shutdownErr := server.Shutdown(ctx); shutdownErr != nil {
			defer func() {
				closeErr := server.Close()
				if closeErr != nil {
					log.Fatalln("Could not close server", closeErr)
				}
			}()
			log.Fatalln("Could not gracefully shutdown the server")
		}
		close(done)
	case <-done:
		return nil

	}
	return nil
}
