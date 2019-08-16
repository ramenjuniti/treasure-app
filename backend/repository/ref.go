package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/model"
)

func FindRefByNoteID(db *sqlx.DB, noteId int64) ([]model.Ref, error) {
	ref := make([]model.Ref, 0)
	if err := db.Select(&ref, `
SELECT id, title, description, link, user_id, note_id FROM ref WHERE note_id = ?
`, noteId); err != nil {
		return nil, err
	}
	return ref, nil
}

func AllRef(db *sqlx.DB) ([]model.Ref, error) {
	ref := make([]model.Ref, 0)
	if err := db.Select(&ref, `SELECT id, title, description, link, user_id, note_id FROM ref`); err != nil {
		return nil, err
	}
	return ref, nil
}

func FindRef(db *sqlx.DB, id int64) (*model.Ref, error) {
	ref := model.Ref{}
	if err := db.Get(&ref, `SELECT id, title, description, link, user_id, note_id FROM ref WHERE id = ?`, id); err != nil {
		return nil, err
	}
	return &ref, nil
}

func CreateRef(db *sqlx.Tx, ref *model.Ref) (sql.Result, error) {
	stmt, err := db.Prepare(`INSERT INTO ref (title, description, link, user_id, note_id) VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(ref.Title, ref.Description, ref.Link, ref.UserID, ref.NoteID)
}

func UpdateRef(db *sqlx.Tx, id int64, ref *model.Ref) (sql.Result, error) {
	stmt, err := db.Prepare(`UPDATE ref SET title = ?, description = ?, link = ? WHERE id = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(ref.Title, ref.Description, ref.Link, id)
}

func DestroyRef(db *sqlx.Tx, id int64) (sql.Result, error) {
	stmt, err := db.Prepare(`DELETE FROM ref WHERE id = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(id)
}
