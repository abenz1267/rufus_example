package handlers

import (
	"net/http"
	"strings"

	"github.com/abenz1267/rufus"
	"github.com/go-chi/chi"
)

// About handler struct
type About struct {
	App         *rufus.App `json:"-"`
	Language    string     `json:"language,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
}

// GetRoutes registers routes for About handler
func (h About) GetRoutes(r *chi.Mux) {
	var url strings.Builder
	url.WriteString("/")

	if h.App.Translation.Amount > 1 {
		url.WriteString(h.App.TranslateURL("about", h.Language))
	} else {
		url.WriteString("about")
	}

	r.Get(url.String(), h.get)
}

func (h About) get(w http.ResponseWriter, r *http.Request) {
	h.Title = "About Page"
	h.Description = "This is a description"

	resp := rufus.Response{Status: http.StatusOK, TemplateFile: "about", Data: h}

	h.App.Response = resp

	h.App.Render(w, r)
}
