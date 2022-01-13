package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneInstruction(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Println(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	instruction, err := app.models.DB.Get(id)
	if err != nil {
		app.logger.Println(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, instruction, "instruction")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getAllInstructions(w http.ResponseWriter, r *http.Request) {
	instructions, err := app.models.DB.All()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, instructions, "instructions")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getAllCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := app.models.DB.CategoriesAll()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, categories, "categories")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) deleteInstruction(w http.ResponseWriter, r *http.Request) {

}

func (app *application) insertInstruction(w http.ResponseWriter, r *http.Request) {

}

func (app *application) updateInstruction(w http.ResponseWriter, r *http.Request) {

}

func (app *application) searchInstructions(w http.ResponseWriter, r *http.Request) {

}
