package handlers

import (
	"net/http"
	"strings"

	"github.com/avn/ory-self-service/session"
	"github.com/go-chi/chi"
)

// type PublicParams struct {
// 	session.SessionStore
// 	*template.Template
// }

// workDir, _ := os.Getwd()
// filesDir := http.Dir(filepath.Join(workDir, "public"))

type FileParams struct {
	session.SessionStore
	Root http.Dir
}

// FileServer sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func (p FileParams) FileServer(path string, r chi.Router) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(p.Root))
		fs.ServeHTTP(w, r)
	})
}

// FileServer sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
// func FileServer(path string, r chi.Router, root http.FileSystem) {
// 	if strings.ContainsAny(path, "{}*") {
// 		panic("FileServer does not permit any URL parameters.")
// 	}

// 	if path != "/" && path[len(path)-1] != '/' {
// 		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
// 		path += "/"
// 	}
// 	path += "*"

// 	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
// 		rctx := chi.RouteContext(r.Context())
// 		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
// 		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
// 		fs.ServeHTTP(w, r)
// 	})
// }
