package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/avn/ory-self-service/api_client"
	"github.com/avn/ory-self-service/handlers"
	middlewares "github.com/avn/ory-self-service/middleware"
	"github.com/avn/ory-self-service/options"
	"github.com/avn/ory-self-service/session"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Start Ory SelfService Solution")

	// Load options from command line and/or environment variables
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Failed to load .env file: %s", err)
	}
	opt := options.NewOptions().SetFronEnvVariables()
	if err := opt.Validate(); err != nil {
		log.Fatalf("Error parsing command line: %v", err)
	} else {
		opt.Display()
	}

	// Setup Kratos API client
	api_client.InitPublicClient(*opt.KratosPublicURL)
	api_client.InitAdminClient(*opt.KratosAdminURL)

	// Setup sesssion store in cookies
	var store = sessions.NewCookieStore(opt.CookieStoreKeyPairs...)

	// Load HTML templates
	tpl := template.Must(template.ParseGlob("templates/*.html"))

	// Set public route parameters
	public := handlers.PublicParams{
		SessionStore: session.SessionStore{Store: store},
		Template:     tpl,
	}

	// Set auth route parameters
	auth := handlers.AuthParams{
		RegistrationFlowRedirectURL: opt.RegistrationFlowURL(),
		LoginFlowRedirectURL:        opt.LoginFlowURL(),
		SettingsFlowRedirectURL:     opt.SettingsFlowURL(),
		RecoveryFlowRedirectURL:     opt.RecoveryFlowURL(),
		LogoutFlowRedirectURL:       opt.LogoutFlowURL(),
	}

	// Set auth middleware parameters
	authM := middlewares.KratosAuthParams{
		SessionStore:      session.SessionStore{Store: store},
		WhoAmIURL:         opt.WhoAmIFlowURL(),
		RedirectUnauthURL: fmt.Sprintf("%s\n", opt.LoginPageURL()),
	}

	// Set private route parameters
	private := handlers.PrivateParams{
		SessionStore: session.SessionStore{Store: store},
		Template:     tpl,
	}

	// Set file server parameters that will serve contents from the ./public/ folder.
	workDir, _ := os.Getwd()
	server := handlers.FileParams{
		SessionStore: session.SessionStore{Store: store},
		Root:         http.Dir(filepath.Join(workDir, "public")),
	}

	// Static assets are wrapped in a hash fs that allows for aggesive http caching
	// var fsys = hashfs.NewFS(staticFS)

	fmt.Printf("Login page URL:  %s\n", opt.LoginPageURL())
	fmt.Printf("Desired login flow URL: %s\n", "http://127.0.0.1:4433/self-service/login/browser")
	fmt.Printf("Actual login flow URL:  %s\n", opt.LoginFlowURL())

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.StripSlashes)
	// Public routes
	r.Get("/", public.Home)
	r.Get("/health", public.Health)
	r.Get("/landing", public.Landing)
	// Authentication routes
	r.Get("/auth/registration", auth.Registration)
	r.Get("/auth/login", auth.Login)
	r.Get("/auth/settings", auth.Settings)
	r.Get("/auth/recovery", auth.Recovery)
	r.Get("/auth/logout", auth.Logout)
	// Private routes
	// r.Get("/private", private.Private)
	r.Route("/private", func(r chi.Router) {
		r.Use(authM.KratosAuthMiddleware)
		r.Get("/", private.Private)
	})
	// File routes along /public
	server.FileServer("/public", r)

	fmt.Printf("Listen and serve on %s\n", opt.Address())
	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", opt.Port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // pass instance of chi router
	}
	if err = srv.ListenAndServe(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
