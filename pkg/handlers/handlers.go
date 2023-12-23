package handlers

import (
	"net/http"

	"github.com/abbassmortazavi/bookings/pkg/config"
	"github.com/abbassmortazavi/bookings/pkg/models"
	"github.com/abbassmortazavi/bookings/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository Type
type Repository struct {
	App *config.AppConfig
}

//NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, "home.html", &models.TempalateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	
	stringMap :=make(map[string]string)
	stringMap["test"] = "Hello Abbass."

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIp

	render.RenderTemplate(w, "about.html", &models.TempalateData{
		StringMap: stringMap,
	})
}
