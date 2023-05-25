package middlewares

import (
	"fmt"
	"log"
	"net/http"

	"github.com/avn/ory-self-service/session"
)

// KratosAuthParams configure the KratosAuth http handler
type KratosAuthParams struct {
	session.SessionStore

	// WhoAmIURL is the API endpoint fo the Kratis 'whoami' call that returns the
	// details of an authenticated session
	WhoAmIURL string

	// RedirectUnauthURL is where we will rerirect to if the session is
	// not associated with a valid user
	RedirectUnauthURL string
}

// KratoAuthMiddleware retrieves the user from the session via Kratos WhoAmIURL,
// and if the user is authenticated the request will proceed through the middleware chain.
// If the session is not authenticated, redirects to the RedirectUnauthURL
func (p KratosAuthParams) KratosAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("KratosAuthMiddleware")
		_, csrfErr := r.Cookie("csrf_token")
		_, sessionErr := r.Cookie("ory_kratos_session")

		if !(csrfErr == nil && sessionErr == nil) {
			log.Printf("- Error retrieving cookies:")
			log.Printf("-- csrf_token: %v\n", csrfErr)
			log.Printf("-- ory_kratos_session: %v\n", sessionErr)
			log.Printf("-- redirect to %s\n", p.RedirectUnauthURL)
			http.Redirect(w, r, p.RedirectUnauthURL, http.StatusPermanentRedirect)
			return
		}

		log.Println("HOW DID I GET HERE")
		next.ServeHTTP(w, r)
	})
}
