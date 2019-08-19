import { connect } from "react-redux";
import firebase from "../../firebase";
import { push } from "connected-react-router";

import Header from "./Header";
import { loginAction, logoutAction } from "../../actions/Auth";

const mapStateToProps = state => {
  return {
    user: state.auth.user
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
    },
    push(path) {
      dispatch(push(path));
    }
  };
};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Header);
