package main

import (
	"context"
	"github.com/ardanlabs/conf"
	"github.com/pkg/errors"
	"github.com/vince15dk/sms-webhook-api/app/sms-api/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// build is the git version of this program. IT is set using build flags in the makefile
var build = "develop"

func main(){
	log := log.New(os.Stdout, "SMS : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	if err := run(log); err != nil{
		log.Println("main: error", err)
		os.Exit(1)
	}
}

func run(log *log.Logger) error{
	// =========================================================================
	// Configuration

	var cfg struct{
		conf.Version
		Web struct{
			APIHost string `conf:"default:0.0.0.0:8080"`
			ReadTimeout time.Duration `conf:"default:5s"`
			WriteTimeout time.Duration `conf:"default:5s"`
			ShutdownTimeout time.Duration `conf:"default:5s"`
		}
	}

	cfg.Version.SVN = build
	cfg.Version.Desc = "copyright v1.0.0"

	if err := conf.Parse(os.Args[1:], "SMS", &cfg); err != nil{
		switch err{
		case conf.ErrHelpWanted:
			usage, err := conf.Usage("SMS", &cfg)
			if err != nil{
				return errors.Wrap(err, "generating config usage")
			}
			log.Println(usage)
			return nil
		case conf.ErrVersionWanted:
			version, err := conf.VersionString("SMS", &cfg)
			if err != nil{
				return errors.Wrap(err, "generating config version")
			}
			log.Println(version)
			return nil
		}
		return errors.Wrap(err, "parsing config")
	}

	defer log.Println("main: Completed")

	out, err := conf.String(&cfg)
	if err != nil{
		return errors.Wrap(err, "generating config for output")
	}
	log.Printf("main: Config :\n%v\n", out)

	// =========================================================================
	// Start API Service

	log.Println("main: Initializing API support")

	// Make a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	api := http.Server{
		Addr: cfg.Web.APIHost,
		Handler: handlers.API(build, shutdown, log),
		ReadTimeout: cfg.Web.ReadTimeout,
		WriteTimeout: cfg.Web.WriteTimeout,
	}

	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this error.
	serverErrors := make(chan error, 1)

	// Start the service listening for requests.
	go func() {
		log.Printf("main: API listening on %s", api.Addr)
		serverErrors <- api.ListenAndServe()
	}()
	// =========================================================================
	// Shutdown

	// Blocking main and waiting for shutdown.
	select{
	case err := <-serverErrors:
		return errors.Wrap(err, "server error")

	case sig := <-shutdown:
		log.Printf("main: %v : Start shutdown", sig)

		// Give outstanding requests a deadline for completion
		ctx, cancel := context.WithTimeout(context.Background(), cfg.Web.ShutdownTimeout)
		defer cancel()

		// Asking listener to shutdown and shed load.
		// hits block until one of theses happen when either all the requests that are in flight at the time we ake for the shutdown completely which is what we want or timeout occurs
		if err := api.Shutdown(ctx); err != nil{
			// when timeout occurs we call this method, hopefully we do not have to do this since timeout is long enough
			api.Close()
			return errors.Wrap(err, "could not stop server gracefully")
		}
	}
	return nil
}