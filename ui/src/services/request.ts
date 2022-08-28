import axios, { AxiosRequestConfig } from "axios";
import { getLocal } from "../utils";

let instance;

function init() {
  const instance = axios.create({
    baseURL: "/api/v1/",
  });
  instance.defaults.headers.post["Content-Type"] = "application/json";

  // 添加请求拦截器
  instance.interceptors.request.use(
    function (config) {
      // 读取Authorization
      const authorization = getLocal("Authorization");
      if (!authorization) {
        // TODO: go to login page
      }
      config.headers["Authorization"] = authorization;
      return config;
    },
    function (error) {
      // 对请求错误做些什么
      return Promise.reject(error);
    }
  );

  // 添加响应拦截器
  instance.interceptors.response.use(
    function (response) {
      // 2xx 范围内的状态码都会触发该函数。
      // 对响应数据做点什么
      return response;
    },
    function (error) {
      // 超出 2xx 范围的状态码都会触发该函数。
      // 对响应错误做点什么
      return Promise.reject(error);
    }
  );
}
init();

export default async function request<T>(
  config: AxiosRequestConfig
): Promise<T> {
  const res = await instance.request();

  return Promise.resolve("" as T);
}
