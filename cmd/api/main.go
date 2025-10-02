package main

import (
	"fmt"
	"io"
	"log"
	"myapp/config"
	"net/http"
)

func main() {

	c := config.New()

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	srv := http.Server{
		Addr:         fmt.Sprintf(":%v", c.Server.Port),
		Handler:      mux,
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutIdle,
	}

	log.Println("Starting server ", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world")
}
