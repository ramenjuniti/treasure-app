package model

type NoteTag struct {
	NoteID int64 `db:"note_id"`
	TagID  int64 `db:"Tag_id"`
}
