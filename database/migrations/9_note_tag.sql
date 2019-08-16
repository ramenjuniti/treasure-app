-- +goose Up
CREATE TABLE note_tag (
  note_id INT(10) UNSIGNED NOT NULL,
  tag_id INT(10) UNSIGNED NOT NULL,
  ctime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (note_id, tag_id),
  CONSTRAINT note_tag_fk_note FOREIGN KEY (note_id) REFERENCES note (id),
  CONSTRAINT note_tag_fk_tag FOREIGN KEY (tag_id) REFERENCES tag (id)
);

-- +goose Down
DROP TABLE note_tag;