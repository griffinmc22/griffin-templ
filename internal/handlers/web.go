package handlers

import (
	"net/http"

	"github.com/griffinmc22/griffin-templ/internal/templates"
)

type WebHandler struct{}

func NewWebHandler() *WebHandler {
	return &WebHandler{}
}

func (h *WebHandler) ServeHTTPAbout(w http.ResponseWriter, r *http.Request) {
	c := templates.About()
	err := templates.Layout(c, "My website").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
