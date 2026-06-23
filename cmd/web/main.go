package main

import (
	"fmt"
	"log"
	"net/http"
	"weapp/pkg/config"
	"weapp/pkg/handlers"
	"weapp/pkg/render"
)

const portNumber = ":8080"

// main is the main function
func main() {

	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot crreate template cache")
	}

	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
