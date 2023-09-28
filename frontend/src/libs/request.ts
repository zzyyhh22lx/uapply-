import { store } from '@/store/index';

type LOGIN_RES = {
  code: number,
  data: {
    account: string,
    token: string
  }
};

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
/**
 * 注册
 * @param data 
 * @returns 
 */
export async function Register(data: {
    account: string,
    password: string
}): Promise<LOGIN_RES> {
  return fetch(`${UAPPLY_URL}/user/register`, {
    method: "post",
    body: JSON.stringify(data)
  }).then(async (res) => {
    const resultResponse = await res.json()
    return resultResponse;
  }).catch((err) => {
    return {};
  });
}
/**
 * 登录
 * @param data 
 * @returns 
 */
export async function Login(data: {
    account: string,
    password: string
}): Promise<LOGIN_RES> {
  return fetch(`${UAPPLY_URL}/user/login`, {
    method: "post",
    body: JSON.stringify(data)
  }).then(async (res) => {
    const resultResponse = await res.json()
    return resultResponse;
  }).catch((err) => {
    return {};
  });
}
