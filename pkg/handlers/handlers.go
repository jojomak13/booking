package handlers

import (
	"github.com/jojomak13/booking/pkg/config"
	"github.com/jojomak13/booking/pkg/models"
	"github.com/jojomak13/booking/pkg/render"
	"net/http"
)

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// New creates a new repository
func New(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr

	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", models.Json{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	remoteIP := m.App.Session.Get(r.Context(), "remote_ip")

	render.RenderTemplate(w, "about.page.tmpl", models.Json{
		"name":      "Joseph",
		"remote_ip": remoteIP,
	})
}
