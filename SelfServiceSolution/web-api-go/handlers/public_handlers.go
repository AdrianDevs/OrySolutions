package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/avn/ory-self-service/session"
)

type PublicParams struct {
	session.SessionStore
	*template.Template
}

func (p PublicParams) Home(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte(fmt.Sprintf("title:%s", "Hello")))
	dataMap := map[string]interface{}{
		"kratosSession": p.GetKratosSession(r),
		"headers":       []string{},
		"pageHeading":   "Homepage",
	}

	err := p.ExecuteTemplate(w, "index.html", dataMap)
	if err != nil {
		log.Fatalln("ERROR: ", nil)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (p PublicParams) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	s := struct {
		OathKeeper string
		Traefik    string
		Prometheus string
		Grafana    string
		Hydra      string
		Kratos     string
		Keto       string
		WebAPI     string
		WebApp     string
		MailHog    string
	}{
		OathKeeper: "In Progress",
		Traefik:    "To do",
		Prometheus: "In Progress",
		Grafana:    "In Progress",
		Hydra:      "To do",
		Kratos:     "To do",
		Keto:       "To do",
		WebAPI:     "In Progress",
		WebApp:     "To do",
		MailHog:    "To do",
	}

	err := json.NewEncoder(w).Encode(s)
	if err != nil {
		fmt.Println("Error encoding json in encode")
		fmt.Println(err.Error())
		log.Println(err.Error())
	}
}

func (p PublicParams) Landing(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte(fmt.Sprintf("title:%s", "Hello")))
	dataMap := map[string]interface{}{
		"kratosSession": p.GetKratosSession(r),
		"headers":       []string{},
		"pageHeading":   "Landing Page",
	}

	err := p.ExecuteTemplate(w, "landing.html", dataMap)
	if err != nil {
		log.Fatalln("ERROR: ", nil)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
