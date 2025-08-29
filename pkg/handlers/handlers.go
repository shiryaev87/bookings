package handlers

import (
	"net/http"

	"github.com/shiryaev87/bookings/pkg/config"
	"github.com/shiryaev87/bookings/pkg/models"
	"github.com/shiryaev87/bookings/pkg/render"
)

// Templatedata holds data sent from handlers to template

// Repo the repository used by the handlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// cgerate new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home  домашняя страница
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "This is the home page")
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About   страница о
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
