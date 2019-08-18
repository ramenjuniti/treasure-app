import React from "react";
import { Route, Switch } from "react-router";
import { Layout } from "antd";
import "antd/dist/antd.css";

import "./App.css";
import Header from "../../components/Header";
import Notes from "../../pages/Notes";
import NoteDetail from "../../pages/NoteDetail";

class App extends React.Component {
  componentDidMount = () => {
    this.props.auth();
  };

  render = () => {
    return (
      <Layout>
        <Header />
        <Layout.Content>
          <div className="App">
            <Switch>
              <Route exact path="/" component={Notes} />
              <Route exact path="/notes/:id" component={NoteDetail} />
            </Switch>
          </div>
        </Layout.Content>
      </Layout>
    );
  };
}

export default App;
