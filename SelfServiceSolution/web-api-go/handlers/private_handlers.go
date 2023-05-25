package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/avn/ory-self-service/session"
)

type PrivateParams struct {
	session.SessionStore
	*template.Template
}

func (p PrivateParams) Private(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte(fmt.Sprintf("title:%s", "Hello")))
	err := p.ExecuteTemplate(w, "private.html", nil)
	if err != nil {
		log.Fatalln("ERROR: ", nil)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
