const tokenKey = "authToken";

const get = () => localStorage.getItem(tokenKey);

const set = (token: string) => localStorage.setItem(tokenKey, token);

const remove = () => localStorage.removeItem(tokenKey);

export const authToken = Object.assign(
  {},
  {
    get,
    set,
    remove
  }
);
