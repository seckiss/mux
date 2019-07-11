package mux

import (
	"github.com/caddyserver/caddy"
	"github.com/caddyserver/caddy/caddyhttp/httpserver"
)

func setup(c *caddy.Controller) error {
	httpserver.GetConfig(c.Key).AddMiddleware(func(next httpserver.Handler) httpserver.Handler {
		return MuxHandler{Next: next}
	})
	return nil
}
