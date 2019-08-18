import React from "react";
import { message, Empty, Card, Tag, Button, Icon, Modal } from "antd";
import parser from "ogp-parser";

import {
  getNoteDetail,
  putNote,
  deleteNote,
  postRef,
  putRef,
  deleteRef
} from "../../api";
import "./NoteDetail.css";
import Ref from "../../components/Ref";
import NoteInput from "../../components/NoteInput";
import RefInput from "../../components/RefInput";

class NoteDetail extends React.Component {
  state = {
    id: 0,
    title: "",
    description: "",
    refs: [],
    tags: [],
    refId: 0,
    refTitle: "",
    refDescription: "",
    refLink: "",
    showUpdateNoteModal: false,
    showNewRefModal: false,
    showUpdateRefModal: false
  };

  componentDidMount = () => {
    this.getIdNote();
  };

  getIdNote = () => {
    const noteId = this.props.location.pathname.split("/").slice(-1)[0];
    getNoteDetail(noteId)
      .then(json => {
        this.setState({
          id: json.id,
          title: json.title,
          description: json.description,
          refs: json.refs.reverse(),
          tags: json.tags
        });
      })
      .catch(err => {
        message.error(err.message);
      });
  };

  postUpdateNote = () => {
    const { user } = this.props;
    const { id, title, description } = this.state;

    if (!user) {
      message.error("please login");
      return;
    }

    user
      .getIdToken()
      .then(token => {
        return putNote(
          token,
          id,
          JSON.stringify({
            title: title,
            description: description
          })
        );
      })
      .then(() => {
        this.getIdNote();
        message.success("Success!!");
      })
      .catch(err => {
        message.error(err.message);
      });
  };

  postNewRef = () => {
    const { user } = this.props;
    const { id, refTitle, refDescription, refLink } = this.state;

    if (!user) {
      message.error("please login");
      return;
    }

    user
      .getIdToken()
      .then(token => {
        return postRef(
          token,
          id,
          JSON.stringify({
            title: refTitle,
            description: refDescription,
            link: refLink
          })
        );
      })
      .then(() => {
        this.getIdNote();
        message.success("Success!!");
      })
      .catch(err => {
        message.error(err.message);
      });
  };

  postUpdateRef = () => {
    const { user } = this.props;
    const { refId, refTitle, refDescription, refLink } = this.state;

    if (!user) {
      message.error("please login");
      return;
    }

    user
      .getIdToken()
      .then(token => {
        return putRef(
          token,
          refId,
          JSON.stringify({
            title: refTitle,
            description: refDescription,
            link: refLink
          })
        );
      })
      .then(() => {
        this.getIdNote();
        message.success("Success!!");
      })
      .catch(err => {
        message.error(err.message);
      });
  };

  deleteIdNote = () => {
    const { user } = this.props;
    const { id } = this.state;

    user
      .getIdToken()
      .then(token => {
        return deleteNote(token, id);
      })
      .then(() => {
        this.getIdNote();
        message.success("Success!!");
      })
      .catch(err => {
        message.error(err.message);
      });
  };

  deleteIdRef = id => {
    const { user } = this.props;

    if (!user) {
      message.error("please login");
      return;
    }

    user
      .getIdToken()
      .then(token => {
        return deleteRef(token, id);
      })
      .then(() => {
        this.getIdNote();
        message.success("Success!!");
      })
      .catch(err => {
        message.error(err.message);
      });
  };

  handleUpdateNoteModalOk = () => {
    this.setState({ showUpdateNoteModal: false });
    this.postUpdateNote();
  };

  handleNewRefModalOk = () => {
    this.setState({ showNewRefModal: false });
    this.postNewRef();
  };

  handleUpdateRefModalOk = () => {
    this.setState({ showUpdateRefModal: false });
    this.postUpdateRef();
  };

  onClickEdit = (id, title, description, link) => {
    this.setState({
      refId: id,
      refTitle: title,
      refDescription: description,
      refLink: link
    });
    console.log(id, title, description, link);
    this.setState({
      showUpdateRefModal: true
    });
  };

  render = () => {
    const {
      id,
      title,
      description,
      refs,
      tags,
      refTitle,
      refDescription,
      refLink,
      showUpdateNoteModal,
      showNewRefModal,
      showUpdateRefModal
    } = this.state;

    const { user } = this.props;

    return (
      <Card
        className="NoteDetail"
        title={title}
        extra={
          <div className="NoteDetail-icons">
            <Icon
              type="edit"
              onClick={() => {
                user
                  ? this.setState({ showUpdateNoteModal: true })
                  : message.error("please login");
              }}
            />
            <Icon type="delete" />
          </div>
        }
      >
        <p>{description}</p>
        {refs.length !== 0 ? (
          refs.map((ref, key) => (
            <Ref
              key={key}
              id={ref.id}
              title={ref.title}
              description={ref.description}
              link={ref.link}
              onClickEdit={() => {
                user
                  ? this.onClickEdit(
                      ref.id,
                      ref.title,
                      ref.description,
                      ref.link
                    )
                  : message.error("plase login");
              }}
              onClickDelete={this.deleteIdRef}
            />
          ))
        ) : (
          <Empty />
        )}
        <Button
          className="NoteDetail-Add-Ref"
          type="primary"
          icon="plus"
          onClick={() => {
            user
              ? this.setState({ showNewRefModal: true })
              : message.error("please login");
          }}
        >
          Add Ref
        </Button>
        <div className="NoteDetail-tags">
          {tags.length !== 0 &&
            tags.map((tag, key) => <Tag key={key}>{tag.name}</Tag>)}
        </div>
        <Modal
          title="Update Note"
          visible={showUpdateNoteModal}
          onOk={this.handleUpdateNoteModalOk}
          onCancel={() => {
            this.setState({ showUpdateNoteModal: false });
          }}
        >
          <NoteInput
            title={title}
            description={description}
            onChangeTitle={e => this.setState({ title: e.target.value })}
            onChangeDescription={e =>
              this.setState({ description: e.target.value })
            }
          />
        </Modal>
        <Modal
          title="new Ref"
          visible={showNewRefModal}
          onOk={this.handleNewRefModalOk}
          onCancel={() => this.setState({ showNewRefModal: false })}
        >
          <RefInput
            onChangeTitle={e => this.setState({ refTitle: e.target.value })}
            onChangeDescription={e =>
              this.setState({ refDescription: e.target.value })
            }
            onChangeLink={e => this.setState({ refLink: e.target.value })}
          />
        </Modal>
        <Modal
          title="Update Ref"
          visible={showUpdateRefModal}
          onOk={() => this.handleUpdateRefModalOk()}
          onCancel={() =>
            this.setState(() => this.setState({ showUpdateRefModal: false }))
          }
        >
          <RefInput
            title={refTitle}
            description={refDescription}
            link={refLink}
            onChangeTitle={e => {
              console.log("kita");
              this.setState({ refTitle: e.target.value });
            }}
            onChangeDescription={e =>
              this.setState({ refDescription: e.target.value })
            }
            onChangeLink={e => {
              this.setState({ refLink: e.target.value });
            }}
          />
        </Modal>
      </Card>
    );
  };
}

export default NoteDetail;
