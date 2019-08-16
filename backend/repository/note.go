package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/model"
)

func AllNote(db *sqlx.DB) ([]model.Note, error) {
	note := make([]model.Note, 0)
	if err := db.Select(&note, `SELECT id, title, description, user_id FROM note`); err != nil {
		return nil, err
	}
	return note, nil
}

func FindNote(db *sqlx.DB, id int64) (*model.Note, error) {
	note := model.Note{}
	if err := db.Get(&note, `SELECT id, title, description, user_id FROM note WHERE id = ?`, id); err != nil {
		return nil, err
	}
	return &note, nil
}

func CreateNote(db *sqlx.Tx, note *model.Note) (sql.Result, error) {
	stmt, err := db.Prepare(`INSERT INTO note (title, description, user_id) VALUES (?, ?, ?)`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(note.Title, note.Description, note.UserID)
}

func UpdateNote(db *sqlx.Tx, id int64, note *model.Note) (sql.Result, error) {
	stmt, err := db.Prepare(`UPDATE note SET title = ?, description = ? WHERE id = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(note.Title, note.Description, id)
}

func DestroyNote(db *sqlx.Tx, id int64) (sql.Result, error) {
	stmt, err := db.Prepare(`DELETE FROM note WHERE id = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(id)
}
