import { request } from "./request";
import type { Article } from "@domains/article";

export default {
  article: {
    getList: async () => {
      return request<
        void,
        {
          items: Article[];
          total: number;
        }
      >({
        url: "/articles",
      });
    },
  },
};
