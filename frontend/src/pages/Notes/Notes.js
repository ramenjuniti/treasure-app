import React from "react";
import { message, Empty, Button, Modal } from "antd";

import { getNotes, postNote } from "../../api";
import "./Notes.css";
import Note from "../../components/Note";
import NoteInput from "../../components/NoteInput";

class Notes extends React.Component {
  state = {
    notes: [],
    showModal: false,
    newTitle: "",
    newDescription: ""
  };

  componentDidMount = () => {
    this.getAllNotes();
  };

  getAllNotes = () => {
    getNotes()
      .then(json => {
        this.setState({ notes: json.reverse() });
      })
      .catch(err => {
        message.error(err.message);
      });
  };

  onClickNote = id => {
    this.props.push(`notes/${id}`);
  };

  handleModalOpen = () => {
    this.props.user
      ? this.setState({ showModal: true })
      : message.error("please login");
  };

  handleModalClose = () => {
    this.setState({ showModal: false });
  };

  postNewNote = () => {
    const { user } = this.props;
    const { newTitle, newDescription } = this.state;

    if (!user) {
      message.error("please login");
      return;
    }

    user
      .getIdToken()
      .then(token => {
        return postNote(
          token,
          JSON.stringify({
            title: newTitle,
            description: newDescription
          })
        );
      })
      .then(() => {
        this.getAllNotes();
        message.success("Success!!");
      })
      .catch(err => {
        message.error(err.message);
      });
  };

  handleModalOk = () => {
    this.handleModalClose();
    this.postNewNote();
  };

  render = () => {
    const { notes, showModal, newTitle, newDescription } = this.state;
    console.log(this.state);
    return (
      <div className="Notes">
        <Button
          className="Notes-Add-Button"
          type="primary"
          icon="plus"
          onClick={this.handleModalOpen}
        >
          Add Note
        </Button>
        {notes.length !== 0 ? (
          notes.map((note, key) => (
            <Note
              key={key}
              title={note.title}
              description={note.description}
              onClickNote={() => this.onClickNote(note.id)}
            />
          ))
        ) : (
          <Empty />
        )}
        <Modal
          title="new Note"
          visible={showModal}
          onOk={this.handleModalOk}
          onCancel={this.handleModalClose}
        >
          <NoteInput
            newTitle={newTitle}
            newDescription={newDescription}
            onChangeTitle={e => this.setState({ newTitle: e.target.value })}
            onChangeDescription={e =>
              this.setState({ newDescription: e.target.value })
            }
          />
        </Modal>
      </div>
    );
  };
}

export default Notes;
