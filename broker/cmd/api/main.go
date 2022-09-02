package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = "3000"

type Config struct{}

func main() {
	app := Config{}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", PORT),
		Handler: app.routes(),
	}

	fmt.Println("Starting on port:", PORT)

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
