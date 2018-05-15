package main

import (
	"log"

	"github.com/abenz1267/rufus"
)

func main() {
	app := &rufus.App{}

	if err := app.LoadConfigAndRouter(); err != nil {
		log.Fatal(err)
	}

	cache := &rufus.Cache{}
	app.Middleware.Cache = cache

	go registerRoutes(app)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
