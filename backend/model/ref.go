package model

type Ref struct {
	ID          int64  `db:"id" json:"id"`
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	Link        string `db:"link" json:"link"`
	UserID      int64  `db:"user_id" json:"user_id"`
	NoteID      int64  `db:"note_id" json:"note_id"`
}
