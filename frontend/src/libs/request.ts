import { store } from '@/store/index';
import { getLocalStorage } from './utils';
import { TOKEN_NAME } from './constant';

export type LOGIN_RES = {
  code: number,
  data: {
    account: string,
    token?: string,
    id?: number
  }
};

const UAPPLY_URL = 'http://localhost:8888/uapply';

async function sendRequest(path: string,data: Record<string, any>,header: Record<string, any>= {}) {
  const bearerToken = getLocalStorage(TOKEN_NAME);
  data = {
    account: store.state.account,
    ...data,
  }
  const options = {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "Authorization": `Bearer ${bearerToken}`,
      // ...
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
/**
 * 获取用户名
 * @returns 
 */
export async function getAccount(): Promise<LOGIN_RES | Boolean> {
  const bearerToken = getLocalStorage(TOKEN_NAME);
  // 无缓存则返回false
  if(!bearerToken) {
    // 如果是测试环境
    if(window.location.hostname === 'localhost') {
      return Promise.resolve({
        "code": 1000,
        "data": {
          "account": "222",
          "id": 6,
        },
        "msg": "Success"
      });
    }
    return Promise.resolve(false)
  };
  return fetch(`${UAPPLY_URL}/user/get-account`, {
    headers: {
      "Authorization": `Bearer ${bearerToken}`,
    },
  }).then(async (res) => {
    const resultResponse = await res.json()
    return resultResponse;
  })
}
