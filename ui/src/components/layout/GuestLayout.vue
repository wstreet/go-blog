<script setup lang="ts">
import WHeader from "@/components/header/header.vue";
import Layout from "@/components/layout/layout.vue";
</script>

<template>
  <Layout>
    <WHeader />
    <RouterView />
  </Layout>
</template>
