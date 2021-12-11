package handlers

import (
	"github.com/vince15dk/sms-webhook-api/business/mid"
	"github.com/vince15dk/sms-webhook-api/foundation/api"
	"log"
	"net/http"
	"os"
)

// API constructs an http.Handler with all application routes defined.
func API(build string, shutdown chan os.Signal, log *log.Logger) *api.App {
	app := api.NewApp(shutdown, mid.Logger(log), mid.Errors(log))

	v1(app, build)

	return app
}

func v1(app *api.App, build string) {
	const version = "v1"
	ug := WebHook{
		build: build,
	}
	app.Handle(http.MethodPost, version, "/:dep/:groups/sms", ug.sendAPItoSMS)
}
