import { request } from "./request";


interface TagsRes {
  items: string[];
  total: number;
}

export default {
  getTags: async () => {
    return request<void, TagsRes>({
      url: "/tags",
    });
  },
};
