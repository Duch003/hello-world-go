package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Duch003/hello-world-go/pkg/config"
	"github.com/Duch003/hello-world-go/pkg/handlers"
	"github.com/Duch003/hello-world-go/pkg/render"
)

const portNumber string = ":8080"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create temaplte cache:", err)
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
