import React from "react";
import { render } from "react-dom";
import thunk from "redux-thunk";
import logger from "redux-logger";
import { createStore, applyMiddleware, compose, combineReducers } from "redux";
import { Provider } from "react-redux";
import { createBrowserHistory } from "history";
import {
  connectRouter,
  routerMiddleware,
  ConnectedRouter
} from "connected-react-router";

import auth from "./reducers/auth";
import App from "./components/App";

export const history = createBrowserHistory();

const rootReducer = combineReducers({
  auth,
  router: connectRouter(history)
});

const store = createStore(
  connectRouter(history)(rootReducer),
  compose(applyMiddleware(thunk, logger, routerMiddleware(history)))
);

render(
  <Provider store={store}>
    <ConnectedRouter history={history}>
      <App />
    </ConnectedRouter>
  </Provider>,
  document.getElementById("root")
);
