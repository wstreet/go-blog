<script setup lang="ts">
import { defineProps } from "vue";

interface Article {
  id: string;
  title: string;
  content: string;
  // tags: string[];
  // categories: string[];
}

interface Props {
  article: Article;
  showContent: boolean;
}

const props = defineProps<Props>();
</script>

<template>
  <div class="mb-1.5">
    <div class="text-violet-500">
      <a
        :href="`/article/${props.article.id}`"
        :class="`${props.showContent ? 'text-2xl mb-6' : ''}`"
      >
        {{ props.article.title }}
      </a>
    </div>
    <div class="mt-6" v-if="props.showContent">{{ props.article.content }}</div>
  </div>
</template>
