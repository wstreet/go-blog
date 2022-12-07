import { request } from "./request";


interface LoginParams {
  username: string,
  password: string
}
export async function adminLogin(params: LoginParams) {
  return request<LoginParams, any>({
    method: 'post',
    url: "/admin/login",
    data: params
  });
}