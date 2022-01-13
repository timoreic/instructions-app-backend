package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	router.HandlerFunc(http.MethodGet, "/v1/instructions", app.getAllInstructions)
	router.HandlerFunc(http.MethodGet, "/v1/instruction/:id", app.getOneInstruction)
	router.HandlerFunc(http.MethodGet, "/v1/instructions/:category_id", app.getAllInstructionsByCategory)

	router.HandlerFunc(http.MethodGet, "/v1/categories", app.getAllCategories)

	return app.enableCORS(router)
}
