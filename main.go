package main

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		&rest.Route{"GET", "/servers", HandlerAllServers},
		&rest.Route{"GET", "/servers/#addr", HandlerGetServer},
		&rest.Route{"DELETE", "/servers/#addr", HandlerDeleteServer},
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)

	log.Fatal(http.ListenAndServe(":3000", api.MakeHandler()))
}
