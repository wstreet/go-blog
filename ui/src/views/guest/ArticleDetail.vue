<script lang="ts">
import { defineComponent, computed } from "vue";
import WArticle from "@/components/article/article.vue";

export default defineComponent({
  components: {
    WArticle,
  },
  props: {
    id: { type: String, required: true },
  },
  setup(props) {
    return {
      article: {
        id: "article3",
        title: "title3",
        content: "content3",
      },
    };
  },
  fetchData() {
    console.log("fetchData");
  },
});
</script>

<template>
   <WArticle :article="article" :show-content="true"></WArticle>
</template>
