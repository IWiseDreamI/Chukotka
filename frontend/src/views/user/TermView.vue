<script setup lang="ts">
import { computed } from "vue";
import { useRoute } from "vue-router";
import { useTermStore } from "@/stores/terms";

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

    <div class="w-full min-h-[400px] text-left" v-html="item.content" />
  </section>
</template>

<style scoped></style>
