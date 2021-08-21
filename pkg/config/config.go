package config

import (
	"log"
	"text/template"
)

// AppConfig holds tha application configuration
type AppConfig struct {
	UseCache bool
	TemplateCache map[string]*template.Template
	InfoLog *log.Logger
}