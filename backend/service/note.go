package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/voyagegroup/treasure-app/dbutil"

	"github.com/voyagegroup/treasure-app/model"
	"github.com/voyagegroup/treasure-app/repository"
)

type Note struct {
	db *sqlx.DB
}

func NewNote(db *sqlx.DB) *Note {
	return &Note{db}
}

func (note *Note) FindNoteDetail(id int64) (*model.NoteDetail, error) {
	n, err := repository.FindNote(note.db, id)
	if err != nil {
		return nil, err
	}

	tags, err := repository.FindNoteTagByNoteID(note.db, id)
	if err != nil {
		return nil, err
	}

	refs, err := repository.FindRefByNoteID(note.db, id)
	if err != nil {
		return nil, err
	}

	noteDetail := &model.NoteDetail{
		Note: *n,
		Tags: tags,
		Refs: refs,
	}

	return noteDetail, nil
}

func (note *Note) Update(id int64, newNote *model.Note) error {
	_, err := repository.FindNote(note.db, id)
	if err != nil {
		return errors.Wrap(err, "failed find note")
	}

	if err := dbutil.TXHandler(note.db, func(tx *sqlx.Tx) error {
		_, err := repository.UpdateNote(tx, id, newNote)
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		return err
	}); err != nil {
		return errors.Wrap(err, "failed note update transaction")
	}
	return nil
}

func (note *Note) Destroy(id int64) error {
	_, err := repository.FindNote(note.db, id)
	if err != nil {
		return errors.Wrap(err, "failed find note")
	}

	if err := dbutil.TXHandler(note.db, func(tx *sqlx.Tx) error {
		_, err := repository.DestroyNote(tx, id)
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		return err
	}); err != nil {
		return errors.Wrap(err, "failed note delete transaction")
	}
	return nil
}

func (note *Note) Create(createNote *model.Note) (int64, error) {
	var createdId int64
	if err := dbutil.TXHandler(note.db, func(tx *sqlx.Tx) error {
		result, err := repository.CreateNote(tx, createNote)
		if err != nil {
			return err
		}
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		createdId = id
		return err
	}); err != nil {
		return 0, err
	}
	return createdId, nil
}
