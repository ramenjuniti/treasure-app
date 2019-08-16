package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/voyagegroup/treasure-app/dbutil"

	"github.com/voyagegroup/treasure-app/model"
	"github.com/voyagegroup/treasure-app/repository"
)

type Ref struct {
	db *sqlx.DB
}

func NewRef(db *sqlx.DB) *Ref {
	return &Ref{db}
}

func (ref *Ref) Update(id int64, newRef *model.Ref) error {
	_, err := repository.FindRef(ref.db, id)
	if err != nil {
		return errors.Wrap(err, "failed find ref")
	}

	if err := dbutil.TXHandler(ref.db, func(tx *sqlx.Tx) error {
		_, err := repository.UpdateRef(tx, id, newRef)
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

func (ref *Ref) Destroy(id int64) error {
	_, err := repository.FindRef(ref.db, id)
	if err != nil {
		return errors.Wrap(err, "failed find ref")
	}

	if err := dbutil.TXHandler(ref.db, func(tx *sqlx.Tx) error {
		_, err := repository.DestroyRef(tx, id)
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		return err
	}); err != nil {
		return errors.Wrap(err, "failed ref delete transaction")
	}
	return nil
}

func (ref *Ref) Create(createRef *model.Ref) (int64, error) {
	var createdId int64
	if err := dbutil.TXHandler(ref.db, func(tx *sqlx.Tx) error {
		result, err := repository.CreateRef(tx, createRef)
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
