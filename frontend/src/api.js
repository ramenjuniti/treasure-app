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

export const getNotes = () => {
  return fetch(`${API_ENDPOINT}/notes`, {
    method: "get"
  }).then(res => {
    if (res.ok) {
      return res.json();
    } else {
      throw Error(`Request rejected with status ${res.status}`);
    }
  });
};

export const getNoteDetail = noteId => {
  return fetch(`${API_ENDPOINT}/notes/${noteId}`, {
    method: "get"
  }).then(res => {
    if (res.ok) {
      return res.json();
    } else {
      throw Error(`Request rejected with status ${res.status}`);
    }
  });
};

export const postNote = (idToken, body) => {
  return fetch(`${API_ENDPOINT}/notes`, {
    method: "post",
    headers: new Headers({
      Authorization: `Bearer ${idToken}`
    }),
    credentials: "same-origin",
    body
  }).then(res => {
    if (res.ok) {
      return res.json();
    } else {
      throw Error(`Request rejected with status ${res.status}`);
    }
  });
};

export const putNote = (idToken, id, body) => {
  return fetch(`${API_ENDPOINT}/notes/${id}`, {
    method: "put",
    headers: new Headers({
      Authorization: `Bearer ${idToken}`
    }),
    credentials: "same-origin",
    body
  }).then(res => {
    if (res.ok) {
      return;
    } else {
      throw Error(`Request rejected with status ${res.status}`);
    }
  });
};

export const deleteNote = (idToken, id) => {
  return fetch(`${API_ENDPOINT}/notes/${id}`, {
    method: "delete",
    headers: new Headers({
      Authorization: `Bearer ${idToken}`
    }),
    credentials: "same-origin"
  }).then(res => {
    if (res.ok) {
      return;
    } else {
      throw Error(`Request rejected with status ${res.status}`);
    }
  });
};

export const getRefs = () => {};

export const postRef = (idToken, note_id, body) => {
  return fetch(`${API_ENDPOINT}/notes/${note_id}/refs`, {
    method: "post",
    headers: new Headers({
      Authorization: `Bearer ${idToken}`
    }),
    credentials: "same-origin",
    body
  }).then(res => {
    if (res.ok) {
      return;
    } else {
      throw Error(`Request rejected with status ${res.status}`);
    }
  });
};

export const putRef = (idToken, id, body) => {
  return fetch(`${API_ENDPOINT}/refs/${id}`, {
    method: "put",
    headers: new Headers({
      Authorization: `Bearer ${idToken}`
    }),
    credentials: "same-origin",
    body
  }).then(res => {
    if (res.ok) {
      return;
    } else {
      throw Error(`Request rejected with status ${res.status}`);
    }
  });
};

export const deleteRef = (idToken, id) => {
  return fetch(`${API_ENDPOINT}/refs/${id}`, {
    method: "delete",
    headers: new Headers({
      Authorization: `Bearer ${idToken}`
    }),
    credentials: "same-origin"
  }).then(res => {
    if (res.ok) {
      return;
    } else {
      throw Error(`Request rejected with status ${res.status}`);
    }
  });
};

export const getTags = () => {};
export const postTag = idToken => {};
export const updateTag = idToken => {};
export const deleteTag = idToken => {};
