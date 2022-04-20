package main

import (
	"fmt"
	"github.com/jjang65/go-hello-word/pkg/config"
	"github.com/jjang65/go-hello-word/pkg/handlers"
	"github.com/jjang65/go-hello-word/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8081"

// main is the main application function
func main() {
	var app config.AppConfig

	// Create templateCache initially to cache templates
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	// Assign templateCache to app.TemplateCache in app config
	app.TemplateCache = tc

	// Set app.UseCache to be false, meaning no templateCache will be used
	//If set to ture, templateCache will be created, newly added template won't be rendered
	// unless app server is compiled again
	app.UseCache = false

	// Passing app reference to use app config in the render package
	render.NewTemplates(&app)

	// create a new repo passing app config to be used in the handlers package
	repo := handlers.NewRepo(&app)
	// Pass pointer to repository to use in the handlers package
	handlers.NewHandlers(repo)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
