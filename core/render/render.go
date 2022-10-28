package render

import (
	"fmt"
	"github.com/jojomak13/booking/core/config"
	"github.com/jojomak13/booking/core/models"
	"github.com/justinas/nosurf"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}
var app *config.AppConfig

// NewTemplates sets the config for the render package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// AddDefaultData bind some data in each render
func AddDefaultData(data models.Json, r *http.Request) models.Json {
	data["csrf_token"] = nosurf.Token(r)

	// SessionFlash
	data["success"] = app.Session.PopString(r.Context(), "success")
	data["warning"] = app.Session.PopString(r.Context(), "warning")
	data["error"] = app.Session.PopString(r.Context(), "error")

	return data
}

// RenderTemplate renders templates using html/template package
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, data models.Json) {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("There is no view called " + tmpl)
	}

	data = AddDefaultData(data, r)

	err := t.Execute(w, data)
	if err != nil {
		fmt.Println("error parsing template ", err)
		return
	}
}

// CreateTemplateCache create a cached map for templates
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		fmt.Println("Caching page " + page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
