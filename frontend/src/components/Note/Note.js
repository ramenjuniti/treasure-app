import React from "react";
import { Card } from "antd";

import "./Note.css";

const Note = ({ title, description, onClickNote }) => (
  <Card className="Note" title={title} onClick={onClickNote}>
    <p>{description}</p>
  </Card>
);

export default Note;
