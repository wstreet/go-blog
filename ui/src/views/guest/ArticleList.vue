<script setup lang="ts">
import { onMounted, ref } from "vue";
import WArticle from "@/components/article/article.vue";
import articleService from "@/services/article";
import type { Article } from "@/domain/article";
import { getUrlParam } from "@/utils";

const tag = getUrlParam("tag");

const articles = ref<Article[]>([]);
onMounted(async () => {
  const res = await articleService.getList({ tag });
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
