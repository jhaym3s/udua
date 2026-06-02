package cmd

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Application struct {
	Config Config
}

type Config struct {
	Addr     string
	DBConfig DBConfig
}

type DBConfig struct {
	Dsn string
}

func (app *Application) Mount() http.Handler {
	r := chi.NewRouter()

  // A good base middleware stack
  r.Use(middleware.RequestID) // this is important for rate-limiting and logging, 
  r.Use(middleware.ClientIPFromRemoteAddr) // pick one ClientIPFrom* based on your infra, see below
  r.Use(middleware.Logger)
  r.Use(middleware.Recoverer)

  // Set a timeout value on the request context (ctx), that will signal
  // through ctx.Done() that the request has timed out and further
  // processing should be stopped.
  r.Use(middleware.Timeout(60 * time.Second))

  r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("healthy"))
  })

	return r
}


func (app *Application) Run(h http.Handler) error {
	srv := &http.Server{
		Addr: app.Config.Addr,
		Handler: h,
		WriteTimeout: time.Second*30,
		ReadTimeout: time.Second*30,
		IdleTimeout: time.Second*60,
	}
	log.Println("Starting server on http://localhost:", app.Config.Addr)
	return srv.ListenAndServe()
}