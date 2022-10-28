package main

import (
	"encoding/gob"
	"github.com/alexedwards/scs/v2"
	"github.com/jojomak13/booking/core/config"
	"github.com/jojomak13/booking/core/handlers"
	"github.com/jojomak13/booking/core/models"
	"github.com/jojomak13/booking/core/render"
	"log"
	"net/http"
	"time"
)

var app config.AppConfig

func main() {
	gob.Register(models.Reservation{})

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
