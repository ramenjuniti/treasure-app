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

type Tag struct {
	db *sqlx.DB
}

func NewTag(db *sqlx.DB) *Tag {
	return &Tag{db: db}
}

func (t *Tag) Index(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	tags, err := repository.AllNote(t.db)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, tags, nil
}

func (t *Tag) Show(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	idTag, err := repository.FindTag(t.db, tid)
	if err != nil && err == sql.ErrNoRows {
		return http.StatusNotFound, nil, err
	} else if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, idTag, nil
}

func (t *Tag) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	newTag := &model.Tag{}
	if err := json.NewDecoder(r.Body).Decode(&newTag); err != nil {
		return http.StatusBadRequest, nil, err
	}

	tagService := service.NewTag(t.db)
	id, err := tagService.Create(newTag)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	newTag.ID = id

	return http.StatusCreated, newTag, nil
}

func (t *Tag) Update(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	reqTag := &model.Tag{}
	if err := json.NewDecoder(r.Body).Decode(&reqTag); err != nil {
		return http.StatusBadRequest, nil, err
	}

	tagService := service.NewTag(t.db)
	err = tagService.Update(tid, reqTag)
	if err != nil && errors.Cause(err) == sql.ErrNoRows {
		return http.StatusNotFound, nil, err
	} else if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusNoContent, nil, nil
}

func (t *Tag) Destroy(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	tagService := service.NewNote(t.db)
	err = tagService.Destroy(tid)
	if err != nil && errors.Cause(err) == sql.ErrNoRows {
		return http.StatusNotFound, nil, err
	} else if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusNoContent, nil, nil
}
