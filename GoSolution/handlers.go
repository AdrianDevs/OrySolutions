package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

func (app *App) landingPage(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte(fmt.Sprintf("title:%s", "Hello")))
	err := app.tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Fatalln("ERROR: ", nil)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *App) userDashboardHandler(writer http.ResponseWriter, request *http.Request) {
	// tmpl, err := template.New("index.html").ParseFiles("index.html")
	// if err != nil {
	// 	http.Error(writer, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	session, err := json.Marshal(getSession(request.Context()))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	// err = tmpl.ExecuteTemplate(writer, "index.html", string(session))
	err = app.tpl.ExecuteTemplate(writer, "userDashboard.html", string(session))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

// FileServer sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}

func IncludeHTML(path string) template.HTML {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("includeHTML - error reading file: %v\n", err)
		return ""
	}

	return template.HTML(string(b))
}
