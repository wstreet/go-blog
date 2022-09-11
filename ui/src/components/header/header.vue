<script setup lang="ts">
import { reactive } from "vue";
import siteInfo from "./config.json";
const menus = reactive(siteInfo.menus);
</script>

<template>
  <header class="mb-12">
    <a href="">
      <div class="text-2xl font-bold">{{siteInfo.site_name}}</div>
    </a>
    <nav>
      <ul class="flex justify-start">
        <li
          v-for="menu in menus"
          :key="menu.path"
          :class="`basis-1/${menus.length} pr-4 text-violet-500 hover:text-purple-700`"
        >
          <RouterLink :to="menu.path"> {{ menu.title }}</RouterLink>
        </li>
      </ul>
    </nav>
  </header>
</template>
