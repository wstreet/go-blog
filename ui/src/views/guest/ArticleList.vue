<script setup lang="ts">
import { onMounted, ref } from "vue";
import WArticle from "@/components/article/article.vue";
import guestService from "@/services/guest";

const articles = ref([
  {
    id: "article1",
    title: "title1",
    content: "content1",
  },
  {
    id: "article2",
    title: "title2",
    content: "content2",
  },
  {
    id: "article3",
    title: "title3",
    content: "content3",
  },
]);
onMounted(async () => {
  const res = await guestService.article.getList();
  articles.value = res.items || [];
});
</script>

<template>
  <div>
    <div class="text-2xl mb-6">Article</div>
    <WArticle
      v-for="(article, index) in articles"
      :key="index"
      :article="article"
      :show-content="false"
    ></WArticle>
  </div>
</template>
