package mid

import (
	"context"
	"github.com/vince15dk/sms-webhook-api/foundation/api"
	"log"
	"net/http"
)

// Logger writes some information about the request to the logs in the
func Logger(log *log.Logger) api.Middleware {
	// This is the actual middleware function to be executed.
	m := func(handler api.Handler) api.Handler {
		// Create the handler that will be attached in the middleware chain.
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			// This is a handler which is h in errors.go
			err := handler(ctx, w, r)

			log.Printf("%s : completed : %s -> %s", r.Method, r.URL.Path,r.RemoteAddr)
			return err
		}
		return h
	}
	return m
}
