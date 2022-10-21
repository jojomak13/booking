package render

import (
	"fmt"
	"github.com/jojomak13/booking/pkg/config"
	"github.com/jojomak13/booking/pkg/models"
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

// RenderTemplate renders templates using html/template package
func RenderTemplate(w http.ResponseWriter, tmpl string, data models.Json) {
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
