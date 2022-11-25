import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: HomeView,
    },
    {
      path: "/article",
      name: "article",
      component: () => import("@/views/guest/ArticleList.vue"),
    },
    {
      path: "/article/:id",
      name: "articleDetail",
      component: () => import("@/views/guest/ArticleDetail.vue"),
      props: () => {}
    },
    {
      path: "/tags",
      name: "tags",
      component: () => import("@/views/guest/Tags.vue"),
      props: () => {}
    },
    {
      path: "/about",
      name: "about",
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import("@/views/AboutView.vue"),
    },
  ],
});

export default router;
