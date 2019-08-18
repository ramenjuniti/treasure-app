import React from "react";
import { Card, Icon, Modal } from "antd";

import RefInput from "../RefInput";
import "./Ref.css";

const Ref = ({ id, title, description, link, onClickEdit, onClickDelete }) => (
  <Card
    className="Ref"
    type="inner"
    title={title}
    extra={
      <div className="Ref-icons">
        <Icon type="edit" onClick={onClickEdit} />
        <Icon type="delete" onClick={() => onClickDelete(id)} />
      </div>
    }
  >
    <a target="_blank" rel="noopener noreferrer" href={link}>
      {link}
    </a>
    <p>{description}</p>
  </Card>
);

export default Ref;
