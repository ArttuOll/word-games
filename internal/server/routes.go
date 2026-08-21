package server

import (
	"net/http"

	"word-games/cmd/web"
	"word-games/cmd/web/wordmaster"

	"github.com/a-h/templ"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/", s.redirectToWordMaster)

	// Word Master
	mux.Handle("/word-master", templ.Handler(wordmaster.WordMaster()))
	mux.HandleFunc("POST /guess", wordmaster.GuessHandler)

	// Serve static files
	fileServer := http.FileServer(http.FS(web.Files))
	mux.Handle("/assets/", fileServer)

	// Wrap the mux with CORS middleware
	return s.corsMiddleware(mux)
}

func (s *Server) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Replace "*" with specific origins if needed
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "false") // Set to "true" if credentials are required

		// Handle preflight OPTIONS requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Proceed with the next handler
		next.ServeHTTP(w, r)
	})
}

func (s *Server) redirectToWordMaster(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/word-master", http.StatusMovedPermanently)
		return
	}
}
