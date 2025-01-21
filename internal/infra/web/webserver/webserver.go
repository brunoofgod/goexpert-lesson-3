package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        *chi.Mux
	Handlers      map[string]map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]map[string]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(method, path string, handler http.HandlerFunc) {
	if s.Handlers[method] == nil {
		s.Handlers[method] = make(map[string]http.HandlerFunc)
	}
	s.Handlers[method][path] = handler
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)

	for method, paths := range s.Handlers {
		for path, handler := range paths {
			s.Router.Method(method, path, handler)
		}
	}

	http.ListenAndServe(s.WebServerPort, s.Router)
}
