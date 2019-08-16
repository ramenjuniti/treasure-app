package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/voyagegroup/treasure-app/dbutil"

	"github.com/voyagegroup/treasure-app/model"
	"github.com/voyagegroup/treasure-app/repository"
)

type Tag struct {
	db *sqlx.DB
}

func NewTag(db *sqlx.DB) *Tag {
	return &Tag{db}
}

func (t *Tag) Update(id int64, newTag *model.Tag) error {
	_, err := repository.FindRef(t.db, id)
	if err != nil {
		return errors.Wrap(err, "failed find ref")
	}

	if err := dbutil.TXHandler(t.db, func(tx *sqlx.Tx) error {
		_, err := repository.UpdateTag(tx, id, newTag)
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		return err
	}); err != nil {
		return errors.Wrap(err, "failed ref update transaction")
	}
	return nil
}

func (t *Tag) Destroy(id int64) error {
	_, err := repository.FindTag(t.db, id)
	if err != nil {
		return errors.Wrap(err, "failed find tag")
	}

	if err := dbutil.TXHandler(t.db, func(tx *sqlx.Tx) error {
		_, err := repository.DestroyRef(tx, id)
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		return err
	}); err != nil {
		return errors.Wrap(err, "failed tag delete transaction")
	}
	return nil
}

func (t *Tag) Create(createTag *model.Tag) (int64, error) {
	var createdId int64
	if err := dbutil.TXHandler(t.db, func(tx *sqlx.Tx) error {
		result, err := repository.CreateTag(tx, createTag)
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
