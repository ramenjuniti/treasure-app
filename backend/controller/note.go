package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/voyagegroup/treasure-app/httputil"
	"github.com/voyagegroup/treasure-app/model"
	"github.com/voyagegroup/treasure-app/repository"
	"github.com/voyagegroup/treasure-app/service"
)

type Note struct {
	db *sqlx.DB
}

func NewNote(db *sqlx.DB) *Note {
	return &Note{db: db}
}

func (note *Note) Index(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	notes, err := repository.AllNote(note.db)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, notes, nil
}

func (note *Note) Show(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	nid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	noteService := service.NewNote(note.db)
	noteDetail, err := noteService.FindNoteDetail(nid)
	if err != nil && err == sql.ErrNoRows {
		return http.StatusNotFound, nil, err
	} else if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, noteDetail, nil
}

func (note *Note) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	newNote := &model.Note{}
	if err := json.NewDecoder(r.Body).Decode(&newNote); err != nil {
		return http.StatusBadRequest, nil, err
	}

	user, err := httputil.GetUserFromContext(r.Context())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	newNote.UserID = user.ID

	noteService := service.NewNote(note.db)
	id, err := noteService.Create(newNote)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	newNote.ID = id

	return http.StatusCreated, newNote, nil
}

func (note *Note) Update(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	nid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	reqNote := &model.Note{}
	if err := json.NewDecoder(r.Body).Decode(&reqNote); err != nil {
		return http.StatusBadRequest, nil, err
	}

	noteService := service.NewNote(note.db)
	err = noteService.Update(nid, reqNote)
	if err != nil && errors.Cause(err) == sql.ErrNoRows {
		return http.StatusNotFound, nil, err
	} else if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusNoContent, nil, nil
}

func (note *Note) Destroy(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	nid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	noteService := service.NewNote(note.db)
	err = noteService.Destroy(nid)
	if err != nil && errors.Cause(err) == sql.ErrNoRows {
		return http.StatusNotFound, nil, err
	} else if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusNoContent, nil, nil
}
