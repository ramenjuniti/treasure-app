package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/dbutil"

	"github.com/voyagegroup/treasure-app/repository"
)

type NoteTag struct {
	db *sqlx.DB
}

func NewNoteTag(db *sqlx.DB) *NoteTag {
	return &NoteTag{db}
}

func (nt *NoteTag) Create(nid, tid int64) error {
	if err := dbutil.TXHandler(nt.db, func(tx *sqlx.Tx) error {
		_, err := repository.CreateNoteTag(tx, nid, tid)
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		return err
	}); err != nil {
		return err
	}
	return nil
}

func (nt *NoteTag) Destory(nid, tid int64) error {
	if err := dbutil.TXHandler(nt.db, func(tx *sqlx.Tx) error {
		_, err := repository.DeleteNoteTag(tx, nid, tid)
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		return err
	}); err != nil {
		return err
	}
	return nil
}
