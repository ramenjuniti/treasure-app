package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/model"
)

func CreateNoteTag(db *sqlx.Tx, noteId int64, tagId int64) (sql.Result, error) {
	stmt, err := db.Prepare(`INSERT INTO note_tag (note_id, tag_id) VALUES (?, ?)`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(noteId, tagId)
}

func DeleteNoteTag(db *sqlx.Tx, noteId int64, tagId int64) (sql.Result, error) {
	stmt, err := db.Prepare(`DELETE FROM note_tag WHERE note_id = ? AND tag_id = ? `)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(noteId, tagId)
}

func FindNoteTagByNoteID(db *sqlx.DB, noteId int64) ([]model.Tag, error) {
	t := make([]model.Tag, 0)
	if err := db.Select(&t, `
SELECT tag.id as id, tag.name as name FROM note_tag 
INNER JOIN tag ON tag.id = note_tag.tag_id
WHERE note_tag.note_id = ?
`, noteId); err != nil {
		return nil, err
	}
	return t, nil
}
