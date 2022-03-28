package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	rpc "github.com/marioidival/brincation-go/rpc"
	"github.com/mitchellh/mapstructure"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err.Error())
	}
}

func loadPorts(client rpc.PortService, path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	for dec.More() {
		m := make(map[string]interface{}, 0)
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Println(err.Error())
			continue
		}

		for k, v := range m {
			var o rpc.Port
			mapstructure.Decode(v, &o)
			o.Id = k
			_, err := client.CreatePort(context.Background(), &o)
			if err != nil {
				log.Println(err.Error())
				continue
			}
		}
	}
}

func run() error {
	client := rpc.NewPortServiceProtobufClient(os.Getenv("PORTS_SERVICE_URL"), &http.Client{})

	go loadPorts(client, os.Getenv("PORTS_FILE"))

	router := mux.NewRouter()
	router.HandleFunc("/get-port/{id}", func(w http.ResponseWriter, r *http.Request) {
	}).Methods("GET")

	server := &http.Server{
		Addr:    os.Getenv("PORT"),
		Handler: router,
	}

	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	serverErrors := make(chan error, 1)
	go func() {
		log.Println("startup client api", server.Addr)
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
