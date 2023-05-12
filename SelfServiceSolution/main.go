package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type App struct {
	tpl *template.Template
}

func main() {
	fmt.Println("Start Ory SelfService Solution")

	app := &App{
		tpl: template.Must(template.ParseGlob("templates/*.html")),
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/", app.index)

	// Create a route along /public that will serve contents from the ./public/ folder.
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "public"))
	FileServer(r, "/public", filesDir)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	// Start the server
	// http.ListenAndServe(":"+port, mux)
	http.ListenAndServe(":"+port, r)
	fmt.Printf("Application launched and running on http://127.0.0.1:%s\n", port)
}

func (app *App) index(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte(fmt.Sprintf("title:%s", "Hello")))
	err := app.tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Fatalln("ERROR: ", nil)
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
