import { request } from "./request";
import type { Article } from "@/domain/article";

interface ArticleListRes {
  items: Article[];
  total: number;
}

interface ArticleListParams {
  tag?: string
}

export default {
  getList: async (params: ArticleListParams) => {
    return request<ArticleListParams, ArticleListRes>({
      url: "/articles",
      data: params
    });
  },
  getDetail: async (id: string) => {
    return request<void, Article>({
      url: `/articles/${id}`,
    });
  },
};
