package handlers

import (
	"net/http"

	"github.com/abenz1267/rufus"
	"github.com/go-chi/chi"
)

// Index handler struct
type Index struct {
	App      *rufus.App
	Language string
}

// GetRoutes registers routes for Index handler
func (h Index) GetRoutes(r *chi.Mux) {
	r.Get("/", h.get)
}

func (h Index) get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
