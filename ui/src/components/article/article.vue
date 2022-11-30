<script setup lang="ts">
import { defineProps } from "vue";
import type { Article } from "@/domain/article";
import { formatDate, formatTime } from "@/utils";

interface Props {
  article: Article;
  showContent: boolean;
}

const props = defineProps<Props>();
</script>

<template>
  <div class="mb-1.5">
    <div class="text-slate-500">
      <span class="pr-4" v-if="!props.showContent">{{
        formatDate(props.article.updated_at)
      }}</span>
      <a
        :href="`/article/${props.article.id}`"
        :class="`${props.showContent ? 'text-2xl mb-6' : ''} underline`"
      >
        {{ props.article.title }}
      </a>
    </div>
    <div v-if="props.showContent">
      <span class="pr-2 text-violet-400">{{ props.article.user_name }}</span>
      <span class="pr-1">发表于</span
      ><span class="text-violet-400 pr-2">{{
        formatTime(props.article.created_at)
      }}</span>
      <span class="pr-1">更新于</span
      ><span class="text-violet-400 pr-2">{{
        formatTime(props.article.updated_at)
      }}</span>
    </div>
    <div class="mt-6" v-if="props.showContent">{{ props.article.content }}</div>
  </div>
</template>
