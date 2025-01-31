<script setup lang="ts">
import { computed } from "vue";
import { useRoute } from "vue-router";
import { useVillageStore } from "@/stores/villages";

import Markdown from "@/components/Markdown.vue";

const route = useRoute();

const villages = useVillageStore();
const villageId = computed(() => Number(route.params.id));
const item = await villages.getVillageById(villageId.value);
</script>

<template>
  <section
    class="flex flex-col items-start gap-[24px] w-full pt-[40px] xl:pt-[120px]"
  >
    <h2>{{ item.name }}</h2>

    <div class="flex flex-col gap-[16px]">
      <Markdown :content="String(item.description)" />
      <Markdown :content="String(item.contacts)" />
      <Markdown :content="String(item.nationality)" />
      <Markdown :content="String(item.population)" />
      <Markdown :content="String(item.organizations)" />
      <Markdown :content="String(item.transport)" />
      <Markdown :content="String(item.connection)" />
      <Markdown :content="String(item.activity)" />
      <Markdown :content="String(item.information)" />
      <Markdown :content="String(item.people)" />
      <Markdown :content="String(item.culture)" />
      <Markdown :content="String(item.socials)" />
    </div>
  </section>
</template>

<style scoped></style>
