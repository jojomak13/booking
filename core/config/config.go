package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
)

// AppConfig holds the application configuration
type AppConfig struct {
	IsProduction  bool
	UseCache      bool
	TemplateCache map[string]*template.Template
	Session       *scs.SessionManager
}
