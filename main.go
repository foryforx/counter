package main

import (
	"context"
	"counter/internal/api"
	"counter/internal/model"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Println("Starting the service...")

	// counter start
	seqGen := model.Initialize()
	defer seqGen.Stop()

	// handler
	http.HandleFunc("/getCounter", api.GetCounterHandler(seqGen))

	// server
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("PORT environment variable was not set")
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Println("Recovered from panic for the counter service:", r)
				// Store the last counter value in database here. seqGen.CurrentCounter
			}
		}()
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down the server...")
	seqGen.Stop()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

}
