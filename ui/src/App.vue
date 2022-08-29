<script setup lang="ts">
import { RouterLink, RouterView } from "vue-router";
import HelloWorld from "./components/HelloWorld.vue";
import WHeader from "@/components/header/header.vue";
</script>

<template>
  <WHeader />
  <RouterView />
</template>
