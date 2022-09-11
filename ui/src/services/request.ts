import axios from "axios";
import type { AxiosRequestConfig, AxiosInstance } from "axios";
import { getLocal } from "../utils";

type RequestParams<P> = {
  method?: "post" | "get";
  url: string;
  data?: P;
  config?: AxiosRequestConfig;
};

type Response<T> = {
  code: number;
  data: T;
  msg: string;
  error: string;
};

let instance: AxiosInstance;

function init() {
  instance = axios.create({
    baseURL: "/api/v1/",
  });
  instance.defaults.headers.post["Content-Type"] = "application/json";

  // 添加请求拦截器
  instance.interceptors.request.use(
    function (config: AxiosRequestConfig) {
      // 读取Authorization
      const authorization = getLocal("Authorization");
      if (!authorization) {
        // TODO: go to login page
      }
      if (!config.headers) {
        config.headers = {};
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

export const request = async <P, T>(
  params: RequestParams<P>
): Promise<T> => {
  const {
    method = "get",
    url,
    data: httpParams,
    config = {} as AxiosRequestConfig,
  } = params;

  const options: AxiosRequestConfig = {
    method,
    url,
    ...config,
  };

  if (method === "get") {
    options.params = httpParams;
  } else {
    options.data = httpParams;
  }

  const { data } = await instance.request<Response<T>>(options);
  // TODO: handle error
  return data.data;
};
