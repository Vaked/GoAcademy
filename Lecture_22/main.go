package main

import (
	"log"
	"net/http"
	"os"
	api "scraper/handlers"
	"time"
	"os/signal"
	"context"
)

func main() {

	l := log.New(os.Stdout, "scraper-api ", log.LstdFlags)
	sh := api.NewScraperAPI(l)
	sm := http.NewServeMux()
	sm.Handle("/top", sh)

	s := http.Server{
		Addr: "localhost:9000",
		Handler: sm,
		ErrorLog: l,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 120 * time.Second,
	}

	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}