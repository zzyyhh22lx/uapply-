import { store } from '@/store/index';

const UAPPLY_URL = 'http://localhost:8888/uapply';

async function sendRequest(path: string,data: Record<string, any>,header: Record<string, any>= {}) {
  data = {
    account: store.state.account,
    ...data,
  }
  const options = {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      // ...header
    },
    body: data ? JSON.stringify(data) : undefined,
    credentials: "include" as RequestCredentials,
    mode: "cors" as RequestMode,
  };
  const response = await fetch(`${URL}${path}`, options);

  return response;
}

export async function Register(data: {
    account: string,
    password: string
}): Promise<string> {
  return fetch(`${UAPPLY_URL}/user/register`, {
    method: "post",
    body: JSON.stringify(data)
  }).then(async (res) => {
    const resultResponse = await res.json()
    console.log(resultResponse);
    return 'hylin';
  }).catch((err) => {
    return 'hylin';
  });
}

export async function Login(data: {
    account: string,
    password: string
}): Promise<string> {
  return fetch(`${UAPPLY_URL}/user/login`, {
    method: "post",
    body: JSON.stringify(data)
  }).then(async (res) => {
    const resultResponse = await res.json()
    console.log(resultResponse);
    return 'hylin';
  }).catch((err) => {
    return 'hylin';
  });
}
