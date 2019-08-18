export const LOGIN = "LOGIN";
export const loginAction = user => {
  return {
    type: LOGIN,
    payload: { user }
  };
};

export const LOGOUT = "LOGOUT";
export const logoutAction = () => {
  return {
    type: LOGOUT
  };
};
