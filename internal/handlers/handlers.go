package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/abbassmortazavi/bookings/internal/config"
	"github.com/abbassmortazavi/bookings/internal/forms"
	"github.com/abbassmortazavi/bookings/internal/models"
	"github.com/abbassmortazavi/bookings/internal/render"
	"log"
	"net/http"
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
	render.RenderTemplate(w, r, "home.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Abbass."

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIp

	render.RenderTemplate(w, r, "about.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) General(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.html", &models.TemplateData{})
}

func (m *Repository) Major(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors.html", &models.TemplateData{})
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.html", &models.TemplateData{})
}

func (m *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.html", &models.TemplateData{})
}
func (m *Repository) PostSearchAvailability(w http.ResponseWriter, r *http.Request) {

	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", start, end)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) PostSearchAvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Hi there",
	}
	out, err := json.MarshalIndent(resp, "", "      ")
	if err != nil {
		log.Println(err)
	}
	log.Println(string(out))
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.html", &models.TemplateData{})
}

func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.html", &models.TemplateData{
		Form: forms.New(nil),
	})
}

func (m *Repository) PostMakeReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}
	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}
	form := forms.New(r.PostForm)
	form.Has("first_name", r)

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.RenderTemplate(w, r, "make-reservation.html", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

}
