-- +goose Up
ALTER TABLE ref ADD note_id int(10) UNSIGNED NOT NULL;
ALTER TABLE ref ADD CONSTRAINT ref_fk_note FOREIGN KEY (note_id) REFERENCES note(id);