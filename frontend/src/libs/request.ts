import { store } from '@/store/index';
import { getLocalStorage, setLocalStorage } from './utils';
import { TOKEN_NAME, ID_NAME } from './constant';

export type LOGIN_RES = {
  code: number,
  data: {
    account: string,
    token?: string,
    id?: number
  }
};

export interface CV_TYPE {
  "user_id": number | undefined,
  "name": string | undefined,
  "age": number | undefined,
  "sex": number | undefined,
  "major": string | undefined,
  "interest": string | undefined,
  "phone": string | undefined,
  "email": string | undefined,
  "init": number | undefined,
  "qq": string | undefined,
  "wc": string | undefined,
}

const UAPPLY_URL = 'http://localhost:8888/uapply';

async function sendRequest(path: string, method: ('POST' | 'GET'), data?: Record<string, any>) {
  const bearerToken = getLocalStorage(TOKEN_NAME);
  data = {
    account: store.state.account,
    ...data,
  }
  const options = {
    method,
    headers: {
      "Content-Type": "application/json",
      "Authorization": `Bearer ${bearerToken}`,
    },
  };
  return fetch(`${path}`, method === 'GET' ? options : { ...options, body: JSON.stringify(data) });

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
      setLocalStorage(TOKEN_NAME, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo2LCJleHAiOjE2OTY3NTI3NTcsImlzcyI6InVhcHBseSJ9.Q9k5tMus-YjODbvq0ljlUZyUwMB37lEP_sJbspLKIB0")
      setLocalStorage(ID_NAME, '6');
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

/**
 * 获取用户简历
 * @returns 
 */
export async function getUserCV(): Promise<{
  code: number,
  data: CV_TYPE,
  msg: string
}> {
  return sendRequest(`${UAPPLY_URL}/user/get-cv/complete`, 'GET').then(async (res) => {
    const resultResponse = await res.json()
    return resultResponse;
  });
}

/**
 * 上传用户简历
 * @returns 
 */
export async function commitUserCV(cv: CV_TYPE) {
  return sendRequest(`${UAPPLY_URL}/user/save-cv/commit`, 'POST', cv).then(async (res) => {
    const resultResponse = await res.json()
    return resultResponse;
  });
}
