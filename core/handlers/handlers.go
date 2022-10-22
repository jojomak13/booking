package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/jojomak13/booking/core/config"
	"github.com/jojomak13/booking/core/models"
	"github.com/jojomak13/booking/core/render"
	"log"
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

	render.RenderTemplate(w, r, "home.page.tmpl", models.Json{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	remoteIP := m.App.Session.Get(r.Context(), "remote_ip")

	render.RenderTemplate(w, r, "about.page.tmpl", models.Json{
		"name":      "Joseph",
		"remote_ip": remoteIP,
	})
}

// Major is the major page handler
func (m *Repository) Major(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors.page.tmpl", models.Json{})
}

// General is the general page handler
func (m *Repository) General(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.page.tmpl", models.Json{})
}

// Search is the search page handler
func (m *Repository) Search(w http.ResponseWriter, r *http.Request) {
	data := fmt.Sprintf("Start Date: %s and End Date: %s", r.Form.Get("start_date"), r.Form.Get("end_date"))
	w.Write([]byte(data))
}

type jsonResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// Availability is the search page handler
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	res := jsonResponse{
		Status:  true,
		Message: "Available",
	}
	out, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
