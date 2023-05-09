package main

import (
	"fmt"
	"net/http"
	"os"

	ory "github.com/ory/client-go"
)

type App struct {
	ory *ory.APIClient
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
		ory: ory.NewAPIClient(c),
	}

	mux := http.NewServeMux()

	// Dashboard
	mux.Handle("/", app.sessionMiddleware(app.dashboardHandler()))

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Printf("Application launched and running on http://127.0.0.1:%s\n", port)

	// Start the server
	http.ListenAndServe(":"+port, mux)
}
