package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Mario-Benedict/note-api/helper"
	"github.com/Mario-Benedict/note-api/model"
	"github.com/go-chi/chi/v5"
)

type NoteController struct{}

var noteModel = new(model.NoteModel)

func (NoteController) GetAll(res http.ResponseWriter, _ *http.Request) {
	notes, err := noteModel.All()

	if err != nil {
		helper.WriteError(res, err, http.StatusInternalServerError)
		return
	}

	helper.WriteJSON(res, notes, http.StatusOK)
}

func (NoteController) GetByID(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(req, "id"))

	if err != nil {
		helper.WriteError(res, err, http.StatusBadRequest)
		return
	}

	note, err := noteModel.GetByID(id)

	if err != nil {
		helper.WriteError(res, err, http.StatusInternalServerError)
		return
	}

	helper.WriteJSON(res, note, http.StatusOK)
}

func (NoteController) Create(res http.ResponseWriter, req *http.Request) {
	var note model.Note

	err := helper.ReadJSON(req, &note)

	if err != nil {
		helper.WriteError(res, err, http.StatusBadRequest)
		return
	}

	if (note.Title == "") || (note.Content == "") || (note.ID != 0) {
		err := errors.New("invalid body request")

		helper.WriteError(res, err, http.StatusBadRequest)
		return
	}

	err = noteModel.Create(note.Title, note.Content)

	if err != nil {
		helper.WriteError(res, err, http.StatusInternalServerError)
		return
	}

	helper.WriteJSON(res, map[string]string{
		"message": "Note created successfully",
	}, http.StatusCreated)
}

func (NoteController) Update(res http.ResponseWriter, req *http.Request) {
	var note model.Note

	err := helper.ReadJSON(req, &note)

	if err != nil {
		helper.WriteError(res, err, http.StatusBadRequest)
		return
	}

	if (note.ID == 0) || (note.Title == "") || (note.Content == "") {
		err := errors.New("invalid body request")

		helper.WriteError(res, err, http.StatusBadRequest)
		return
	}

	err = noteModel.Update(note.ID, note.Title, note.Content)

	if err != nil {
		helper.WriteError(res, err, http.StatusInternalServerError)
		return
	}

	helper.WriteJSON(res, map[string]string{
		"message": "Note updated successfully",
	}, http.StatusOK)
}

func (NoteController) Delete(res http.ResponseWriter, req *http.Request) {
	var note model.Note

	err := helper.ReadJSON(req, &note)

	if err != nil {
		helper.WriteError(res, err, http.StatusBadRequest)
		return
	}

	if note.ID == 0 {
		err := errors.New("invalid body request")

		helper.WriteError(res, err, http.StatusBadRequest)
		return
	}

	err = noteModel.Delete(note.ID)

	if err != nil {
		helper.WriteError(res, err, http.StatusInternalServerError)
		return
	}

	helper.WriteJSON(res, map[string]string{
		"message": "Note deleted successfully",
	}, http.StatusOK)
}
