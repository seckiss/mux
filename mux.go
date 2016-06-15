package mux

import (
	"net/http"

	"github.com/mholt/caddy"
	"github.com/mholt/caddy/caddyhttp/httpserver"
)

var DefaultCaddyServeMux = http.NewServeMux()

func init() {
	caddy.RegisterPlugin("mux", caddy.Plugin{
		ServerType: "http",
		Action:     setup,
	})
}

type MuxHandler struct {
	Next httpserver.Handler
}

// Get the default Caddy ServeMux
func ServeMux() *http.ServeMux {
	return DefaultCaddyServeMux
}

// Register the handler for the given pattern in the default Caddy ServeMux
func Handle(pattern string, handler http.Handler) {
	DefaultCaddyServeMux.Handle(pattern, handler)
}

// Registers the handler function for the given pattern in the default Caddy ServeMux
func HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	DefaultCaddyServeMux.HandleFunc(pattern, handler)
}

func (m MuxHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) (int, error) {
	_, pattern := DefaultCaddyServeMux.Handler(r)
	if len(pattern) > 0 {
		DefaultCaddyServeMux.ServeHTTP(w, r)
		return 0, nil
	} else {
		// no matching filter
		return m.Next.ServeHTTP(w, r)
	}
}
