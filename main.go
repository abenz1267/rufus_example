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

	go registerRoutes(app)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
