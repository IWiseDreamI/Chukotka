<script setup lang="ts">
import { computed } from "vue";
import { useRoute } from "vue-router";
import { useTermStore } from "@/stores/terms";

import Markdown from "@/components/Markdown.vue";

const route = useRoute();

const terms = useTermStore();
const termId = computed(() => Number(route.params.id));
const item = await terms.getTermById(termId.value);
</script>

<template>
  <section
    class="flex flex-col items-start gap-[24px] w-full pt-[40px] xl:pt-[120px]"
  >
    <h2>{{ item.title }}</h2>

    <Markdown class="min-w-[400px]" :content="item.content" />
  </section>
</template>

<style scoped></style>
