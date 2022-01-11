package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)


func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	router.HandlerFunc(http.MethodGet, "/v1/instruction/:id", app.getOneInstruction)
	router.HandlerFunc(http.MethodGet, "/v1/instructions", app.getAllInstructions)

	return router
}