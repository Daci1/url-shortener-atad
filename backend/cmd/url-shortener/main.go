package main

import (
	"github.com/Daci1/url-shortener-atad/internal/server"
	"log"
)

func main() {
	e := server.NewServer()

	log.Fatal(e.Start(":8080"))
}
