package mid

import (
	"context"
	"github.com/vince15dk/sms-webhook-api/foundation/api"
	"log"
	"net/http"
)

func Errors(log *log.Logger) api.Middleware {

	m := func(handler api.Handler) api.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

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
