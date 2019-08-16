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

type Ref struct {
	db *sqlx.DB
}

func NewRef(db *sqlx.DB) *Ref {
	return &Ref{db: db}
}

func (ref *Ref) Index(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	refs, err := repository.AllRef(ref.db)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, refs, nil
}

func (ref *Ref) Show(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	rid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	idRef, err := repository.FindRef(ref.db, rid)
	if err != nil && err == sql.ErrNoRows {
		return http.StatusNotFound, nil, err
	} else if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, idRef, nil
}

func (ref *Ref) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	newRef := &model.Ref{}
	if err := json.NewDecoder(r.Body).Decode(&newRef); err != nil {
		return http.StatusBadRequest, nil, err
	}

	vars := mux.Vars(r)
	nidStr, ok := vars["note_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	nid, err := strconv.ParseInt(nidStr, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	newRef.NoteID = nid

	user, err := httputil.GetUserFromContext(r.Context())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	newRef.UserID = user.ID

	refService := service.NewRef(ref.db)
	id, err := refService.Create(newRef)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	newRef.ID = id

	return http.StatusCreated, newRef, nil
}

func (ref *Ref) Update(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	rid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	refRef := &model.Ref{}
	if err := json.NewDecoder(r.Body).Decode(&refRef); err != nil {
		return http.StatusBadRequest, nil, err
	}

	refService := service.NewRef(ref.db)
	err = refService.Update(rid, refRef)
	if err != nil && errors.Cause(err) == sql.ErrNoRows {
		return http.StatusNotFound, nil, err
	} else if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusNoContent, nil, nil
}

func (ref *Ref) Destroy(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	rid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	refService := service.NewRef(ref.db)
	err = refService.Destroy(rid)
	if err != nil && errors.Cause(err) == sql.ErrNoRows {
		return http.StatusNotFound, nil, err
	} else if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusNoContent, nil, nil
}
