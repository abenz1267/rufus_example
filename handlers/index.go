package handlers

import (
	"fmt"
	"net/http"

	"github.com/abenz1267/rufus"
	"github.com/go-chi/chi"
)

// Index handler struct
type Index struct {
	App         *rufus.App `json:"-"`
	Language    string     `json:"language,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
}

// GetRoutes registers routes for Index handler
func (h Index) GetRoutes(r *chi.Mux) {
	r.Get("/", h.get)
	r.Post("/", h.post)
}

func (h Index) get(w http.ResponseWriter, r *http.Request) {
	h.Title = "Index Page"
	h.Description = "This is a description"

	resp := rufus.Response{Status: http.StatusOK, TemplateFile: "index", Data: h}

	h.App.Response = resp

	h.App.Render(w, r)
}

func (h Index) post(w http.ResponseWriter, r *http.Request) {
	// do something
	deleted := h.App.Cache.Invalidate(r.RequestURI)
	fmt.Println("Cache deleted:", deleted)
}
