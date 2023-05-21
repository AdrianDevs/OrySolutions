package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	ory "github.com/ory/client-go"
)

type App struct {
	tpl *template.Template
	ory *ory.APIClient
}

var functionMap = template.FuncMap{
	"IncludeHTML": IncludeHTML,
}

func main() {
	fmt.Println("Start Go Kratos")

	proxyPort := os.Getenv("PROXY_PORT")
	if proxyPort == "" {
		proxyPort = "4000"
	}

	fmt.Printf("Proxy port: %v\n", proxyPort)

	// register a new Ory client with the URL set to the Ory CLI Proxy
	// we can also read the URL from the env or a config file
	c := ory.NewConfiguration()
	c.Servers = ory.ServerConfigurations{{URL: fmt.Sprintf("http://localhost:%s/.ory", proxyPort)}}

	app := &App{
		tpl: template.Must(template.New("").Funcs(functionMap).ParseGlob("templates/*.html")),
		ory: ory.NewAPIClient(c),
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", app.landingPage)
	r.Route("/user", func(r chi.Router) {
		r.Use(app.sessionMiddleware)
		r.Get("/", app.userDashboardHandler)
	})

	// Create a route along /public that will serve contents from the ./public/ folder.
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "public"))
	FileServer(r, "/public", filesDir)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Printf("Application launched and running on http://127.0.0.1:%s\n", port)

	// Start the server
	// http.ListenAndServe(":"+port, mux)
	http.ListenAndServe(":"+port, r)
}
