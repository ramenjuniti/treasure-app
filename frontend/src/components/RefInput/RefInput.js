import React from "react";
import { Input } from "antd";

import "./RefInput.css";

const RefInput = ({
  title,
  description,
  link,
  onChangeTitle,
  onChangeDescription,
  onChangeLink
}) => {
  return (
    <div className="RefInput">
      <Input
        placeholder="Ref Title"
        defaultValue={title && title}
        onChange={onChangeTitle}
      />
      <Input.TextArea
        placeholder="Ref Description"
        defaultValue={description && description}
        onChange={onChangeDescription}
      />
      <Input
        placeholder="Ref Link"
        defaultValue={link && link}
        onChange={onChangeLink}
      />
    </div>
  );
};

export default RefInput;
