package handlers

import (
	"net/http"

	"github.com/abenz1267/rufus"
	"github.com/go-chi/chi"
)

// NotFound handler for 404 requests
type NotFound struct {
	App          *rufus.App `json:"-"`
	Language     string     `json:"language,omitempty"`
	Title        string     `json:"title,omitempty"`
	Description  string     `json:"description,omitempty"`
	ErrorMessage string     `json:"error_message,omitempty"`
}

// GetRoutes registers routes for Index handler
func (h NotFound) GetRoutes(r *chi.Mux) {
	r.Get("/404", h.Get)
}

// Get is the handler for the custom 404 page
func (h NotFound) Get(w http.ResponseWriter, r *http.Request) {
	h.Title = "404"
	h.Description = "Page not found"
	h.ErrorMessage = "Page not found"

	resp := rufus.Response{Status: http.StatusNotFound, TemplateFile: "error", Data: h}

	h.App.Response = resp
	h.App.Render(w, r)
}
