package main

import (
	"log"

	"github.com/Daci1/url-shortener-atad/internal/db"
	"github.com/Daci1/url-shortener-atad/internal/server"
)

func main() {
	db.Init()
	defer db.Close()

	e := server.NewServer()

	log.Fatal(e.Start(":8080"))
}
