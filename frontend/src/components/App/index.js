import { connect } from "react-redux";
import firebase from "../../firebase";

import App from "./App";
import { loginAction, logoutAction } from "../../actions/Auth";

const mapStateToProps = state => {
  console.log(state);
  return {
    user: state.auth.user,
    message: state.get
  };
};

const mapDispatchToProps = dispatch => {
  return {
    login() {
      firebase.login();
    },
    auth() {
      firebase.auth().onAuthStateChanged(user => {
        if (!user) {
          return;
        }
        dispatch(loginAction(user));
      });
    },
    logout() {
      firebase.logout();
      dispatch(logoutAction());
    }
  };
};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(App);
