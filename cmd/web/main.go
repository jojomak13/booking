package main

import (
	"github.com/alexedwards/scs/v2"
	"github.com/jojomak13/booking/pkg/config"
	"github.com/jojomak13/booking/pkg/handlers"
	"github.com/jojomak13/booking/pkg/render"
	"log"
	"net/http"
	"time"
)

var app config.AppConfig

func main() {
	app.IsProduction = false

	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.IsProduction
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	render.NewTemplates(&app)

	repo := handlers.New(&app)

	srv := &http.Server{
		Addr:    ":8000",
		Handler: routes(repo),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("cannot start the serve")
	}
}
