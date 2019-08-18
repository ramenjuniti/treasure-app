import React from "react";
import { Input } from "antd";

import "./NoteInput.css";

const NoteInput = ({
  title,
  description,
  onChangeTitle,
  onChangeDescription
}) => {
  return (
    <div className="NoteInput">
      <Input
        placeholder="Note Title"
        defaultValue={title && title}
        onChange={onChangeTitle}
      />
      <Input.TextArea
        placeholder="Note Description"
        defaultValue={description && description}
        onChange={onChangeDescription}
      />
    </div>
  );
};

export default NoteInput;
