package main

import (
	"log"
	"net/http"

	"github.com/eslopezm/Agenda_Personal/commons"
	"github.com/eslopezm/Agenda_Personal/routes"
	"github.com/gorilla/mux"
)

func main() {
	commons.Migrate()

	router := mux.NewRouter()

	routes.SetPersonaRoutes(router)

	server := http.Server{
		Addr:    ":9000",
		Handler: router,
	}

	log.Println("--- Servidor ejecutandose en el puerto 127.0.0.1:6")

	log.Println(server.ListenAndServe())
}
