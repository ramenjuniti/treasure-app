package model

type Note struct {
	ID          int64  `db:"id" json:"id"`
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	UserID      int64  `db:"user_id" json:"user_id"`
}

type NoteDetail struct {
	Note
	Refs []Ref `json:"refs"`
	Tags []Tag `json:"tags"`
}
