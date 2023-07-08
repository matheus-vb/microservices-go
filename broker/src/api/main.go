package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = "80"

type Config struct{}

func main() {
	app := Config{}

	log.Println("Starting broker service on port ", port)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: app.setRoutes(),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
