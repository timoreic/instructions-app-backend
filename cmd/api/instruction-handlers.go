package main

import (
	"backend/models"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneInstruction(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Println(errors.New("invalid id parameter"))
	}

	app.logger.Println("ID is", id)

	instruction := models.Instruction {
		ID: id,
		Title: "Some Instruction",
		Description: "Some description",
		Steps: []string{"Go", "Slices", "Are", "Powerful"},
		Rating: 5,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = app.writeJSON(w, http.StatusOK, instruction, "instruction")
}

func (app *application) getAllInstructions(w http.ResponseWriter, r *http.Request) {
	
}