package main

import (
	"fmt"
	"log"
	"myapp/api/router"
	"myapp/config"
	"net/http"
)

func main() {

	c := config.New()

	r := router.New()

	srv := http.Server{
		Addr:         fmt.Sprintf(":%v", c.Server.Port),
		Handler:      r,
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutIdle,
	}

	log.Println("Starting server ", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}
