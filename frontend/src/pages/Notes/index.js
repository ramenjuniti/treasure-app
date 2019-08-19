import { connect } from "react-redux";
import { push } from "connected-react-router";

import Notes from "./Notes";

const mapStateToProps = state => {
  return {
    user: state.auth.user
  };
};

const mapDispatchToProps = dispatch => {
  return {
    push(path) {
      dispatch(push(path));
    }
  };
};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Notes);