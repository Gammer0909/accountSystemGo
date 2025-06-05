package main

import (
	"context"
	"errors"
	"log"
	"muxHello/server"
	"os"
	"os/signal"
	"syscall"
	"time"

	"net/http"
)

func main() {

	db, err := os.OpenFile("./data.csv", os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("Failed opening file ./data.csv")
	}
	s := server.NewServer(db)
	defer s.CloseServer()

	http.HandleFunc("/api/signup", s.SignUp)
	http.HandleFunc("/api/signin", s.SignIn)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Println("Serving http://localhost:8080")

	// https://dev.to/mokiat/proper-http-shutdown-in-go-3fji thnx
	server := &http.Server{
		Addr: ":8080",
	}

	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
		defer shutdownRelease()

		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Fatalf("HTTP shutdown error: %v", err)
		}
		log.Println("Graceful shutdown complete.")
	}()

	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("HTTP server error: %v", err)
	}
	log.Println("Stopped serving new connections.")

	s.CloseServer()

}
