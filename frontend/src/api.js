const API_ENDPOINT = process.env.REACT_APP_BACKEND_API_BASE;

export const getPrivateMessage = idToken => {
  return fetch(`${API_ENDPOINT}/private`, {
    method: "get",
    headers: new Headers({
      Authorization: `Bearer ${idToken}`
    }),
    credentials: "same-origin"
  }).then(res => {
    if (res.ok) {
      return res.json();
    } else {
      throw Error(`Request rejected with status ${res.status}`);
    }
  });
};

export const getPublicMessage = () => {
  return fetch(`${API_ENDPOINT}/public`);
};

export const getNote = () => {};
export const postNote = idToken => {};
export const updateNote = idToken => {};
export const deleteNote = idToken => {};

export const getRef = () => {};
export const postRef = idToken => {};
export const updateRef = idToken => {};
export const deleteRef = idToken => {};

export const getTag = () => {};
export const postTag = idToken => {};
export const updateTag = idToken => {};
export const deleteTag = idToken => {};
