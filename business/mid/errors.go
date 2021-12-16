package mid

import (
	"context"
	"github.com/vince15dk/sms-webhook-api/foundation/api"
	"log"
	"net/http"
)

// Errors handles wrap up function for child handler
func Errors(log *log.Logger) api.Middleware {
	// This is the actual middleware function to be executed.
	m := func(handler api.Handler) api.Handler {
		// This is the custom handler
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			// net/http package handler
			if err := handler(ctx, w, r); err != nil {
				log.Printf("ERROR	: %v", err)
				if err := api.RespondError(ctx, w, err); err != nil {
					return err
				}
				if ok := api.IsShutdown(err); ok {
					return err
				}
			}
			return nil
		}
		return h
	}
	return m
}
