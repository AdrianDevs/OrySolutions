package main

import (
	"context"
	"log"
	"net/http"

	ory "github.com/ory/client-go"
)

// Save the cookies for any upstream calls to the Ory APIs
func withCookies(ctx context.Context, v string) context.Context {
	return context.WithValue(ctx, "req.cookies", v)
}

func getCookies(ctx context.Context) string {
	return ctx.Value("req.cookies").(string)
}

// Save the session to display it on the dashboard
func withSession(ctx context.Context, v *ory.Session) context.Context {
	return context.WithValue(ctx, "req.session", v)
}

func getSession(ctx context.Context) *ory.Session {
	return ctx.Value("req.session").(*ory.Session)
}

func (app *App) sessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("Handling middleware request\n")

		// Set the cookies on the ory client
		var cookies string

		// This example passes all request.Cookies to 'ToSession' function
		// However, you can pass only the value of ory_session_projectid
		// cookie to the endpoint
		cookies = request.Header.Get("Cookie")

		// Check if we have a session
		session, _, err := app.ory.FrontendApi.ToSession(request.Context()).Cookie(cookies).Execute()
		if (err != nil && session == nil) || (err == nil && !*session.Active) {
			if err != nil {
				log.Printf("Middleware - Error - FrontendAPI: %v\n", err.Error())
			}
			if session != nil {
				log.Printf("Middleware - Session is nil\n")
			}
			if session != nil && !*session.Active {
				log.Printf("Middleware - Session is not active\n")
			}

			// This will rediret the user to the managed Ory Login UI
			log.Println("User does not have a sesion - redirect to login")
			http.Redirect(writer, request, "/.ory/self-service/login/browser", http.StatusSeeOther)
			return
		}

		ctx := withCookies(request.Context(), cookies)
		ctx = withSession(ctx, session)

		// Continue to the requested page (in our case the Dashboard)
		log.Println("User has a valid sesion - continue to dashboard")
		next.ServeHTTP(writer, request.WithContext(ctx))
		return
	})
}
