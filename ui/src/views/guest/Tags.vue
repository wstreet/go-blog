<script setup lang="ts">
import { ref, onMounted } from 'vue'
import tagsService from '@/services/tags'

const tags = ref<string[]>([]);
onMounted(async () => {
  const res = await tagsService.getTags();
  tags.value = res.items || [];
});
</script>
<template>
  <div>
    <ul>
      <li v-for="tag in tags"> 
        <a :href="`/article?tag=${tag}`">{{tag}}</a>
      </li>
    </ul>
  </div>
</template>