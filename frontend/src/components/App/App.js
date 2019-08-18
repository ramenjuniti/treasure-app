import React from "react";
import { Route, Switch } from "react-router";
import "./App.css";
import { getPrivateMessage, getPublicMessage } from "../../api";

class App extends React.Component {
  state = {
    message: "",
    errorMessage: ""
  };

  componentDidMount = () => {
    this.props.auth();
  };

  getPrivateMessage = () => {
    this.props.user
      .getIdToken()
      .then(token => {
        return getPrivateMessage(token);
      })
      .then(resp => {
        this.setState({
          message: resp.message
        });
      })
      .catch(error => {
        this.setState({
          errorMessage: error.toString()
        });
      });
  };

  render = () => {
    const { login, logout, user } = this.props;
    const { message, errorMessage } = this.state;

    const app = () => (
      <div>
        <div>{message}</div>
        <p>{errorMessage}</p>
        <button onClick={this.getPrivateMessage.bind(this)}>
          Get Private Message
        </button>
        <button onClick={logout}>Logout</button>
      </div>
    );

    if (user === null) {
      return <button onClick={login}>Please login</button>;
    }
    return (
      <div>
        <Switch>
          <Route exact path="/" render={app} />
        </Switch>
      </div>
    );
  };
}

export default App;
