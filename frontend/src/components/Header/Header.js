import React from "react";
import { Layout, Button, Avatar } from "antd";

import "./Header.css";

const Header = props => {
  const { login, logout, user } = props;

  return (
    <Layout.Header className="Header">
      <h1 className="Header-left">reftumu</h1>
      <div className="Header-right">
        {user ? (
          <div className="Header-button-area">
            <div className="Header-user-area">
              <Avatar src={user.photoURL} />
              <span>{user.displayName}</span>
            </div>
            <Button onClick={logout}>Logout</Button>
          </div>
        ) : (
          <Button onClick={login}>Login</Button>
        )}
      </div>
    </Layout.Header>
  );
};

export default Header;
