package main

import (
	"github.com/abenz1267/rufus"
	"github.com/go-chi/chi"

	"github.com/abenz1267/rufus_example/handlers"
)

func registerRoutes(app *rufus.App) {
	for i := 0; i < app.Translation.Amount; i++ {
		language := <-app.RoutesSender

		if language == "default" {
			language = app.Language
		}

		r := chi.NewRouter()
		r.NotFound(handlers.NotFound{App: app, Language: language}.Get)

		handlers.Index{App: app, Language: language}.GetRoutes(r)

		app.RoutesReceiver <- r
	}
}
