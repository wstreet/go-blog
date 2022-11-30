<script setup lang="ts">
import { ref, onMounted } from "vue";
import WArticle from "@/components/article/article.vue";
import articleService from "@/services/article";
import type { Article } from "@/domain/article";
import { useRoute } from "vue-router";

const article = ref<Article>({} as Article);
const loading = ref<boolean>();
const route = useRoute();
const id = route.params.id as string;

onMounted(() => {
  queryArticle();
});

const queryArticle = async () => {
  loading.value = true;
  const res = await articleService.getDetail(id);
  article.value = res;
  loading.value = false;
};
</script>

<template>
  <div v-if="loading">loading</div>
  <WArticle v-else :article="article" :show-content="true"></WArticle>
</template>
