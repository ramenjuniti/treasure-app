package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"github.com/voyagegroup/treasure-app/httputil"
	"github.com/voyagegroup/treasure-app/service"
)

type NoteTag struct {
	db *sqlx.DB
}

func NewNoteTag(db *sqlx.DB) *NoteTag {
	return &NoteTag{db: db}
}

func (nt *NoteTag) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	nidStr, ok := vars["note_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	nid, err := strconv.ParseInt(nidStr, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	tidStr, ok := vars["tab_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	tid, err := strconv.ParseInt(tidStr, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	noteTagService := service.NewNoteTag(nt.db)
	if err := noteTagService.Create(nid, tid); err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, nil, nil
}

func (nt *NoteTag) Destroy(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	nidStr, ok := vars["note_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	nid, err := strconv.ParseInt(nidStr, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	tidStr, ok := vars["tab_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	tid, err := strconv.ParseInt(tidStr, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	noteTagService := service.NewNoteTag(nt.db)
	if err := noteTagService.Destory(nid, tid); err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusNoContent, nil, nil
}
