import { createRouter, createWebHistory } from "vue-router";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "root",
      component: () => import("@/components/layout/GuestLayout.vue"),
      // redirect: "/home",
      children: [
        {
          path: "/home",
          name: "home",
          component: () => import("@/views/HomeView.vue"),
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
          component: () => import("@/views/AboutView.vue"),
        },
      ]
    },
    {
      path: "/admin/login",
      name: "login",
      component: () => import("@/views/admin/Login.vue"),
    },
  ],
});

export default router;
