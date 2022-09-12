import { request } from "./request";
import type { Article } from "@domains/article";

interface ArticleListRes {
  items: Article[];
  total: number;
}

export default {
  getList: async () => {
    return request<void, ArticleListRes>({
      url: "/articles",
    });
  },
  getDetail: async (id: string) => {
    return request<void, Article>({
      url: `/articles/${id}`,
    });
  },
};
