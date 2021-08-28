package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Duch003/hello-world-go/pkg/config"
	"github.com/Duch003/hello-world-go/pkg/handlers"
	"github.com/Duch003/hello-world-go/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber string = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main() {
	//Change it to true while in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create temaplte cache:", err)
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf(fmt.Sprintf("Starting application on port %s", portNumber))

	// _ = http.ListenAndServe(portNumber, nil)
	srv := &http.Server {
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
