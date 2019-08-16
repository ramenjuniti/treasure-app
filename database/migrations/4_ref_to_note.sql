-- +goose Up
CREATE TABLE ref_to_note (
  ref_id INT(10) UNSIGNED NOT NULL,
  note_id INT(10) UNSIGNED NOT NULL,
  CONSTRAINT ref_to_note_fk_ref FOREIGN KEY (ref_id) REFERENCES ref (id),
  CONSTRAINT ref_to_note_fk_note FOREIGN KEY (note_id) REFERENCES note (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DROP TABLE ref_to_note;