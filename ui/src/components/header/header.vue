<template>
  <header>
    <nav>
      <ul class="flex flex-row">
        <li
          v-for="menu in menus"
          :key="menu.path"
          :class="`basis-1/${menus.length} pr-6 pl-6`"
        >
         <RouterLink :to="menu.path"> {{ menu.title }}</RouterLink>
        </li>
      </ul>
    </nav>
  </header>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import siteInfo from "./config.json";

export default defineComponent({
  name: "WHeader",
  setup(props, { emit }) {
    return { menus: siteInfo.menus };
  },
});
</script>
