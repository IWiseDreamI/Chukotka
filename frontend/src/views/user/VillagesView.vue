<script setup lang="ts">
import { useVillageStore } from "@/stores/villages";

const villagesStore = useVillageStore();

const villages = await villagesStore.getVillages();
</script>

<template>
  <section class="flex flex-wrap gap-[24px] w-full py-[40px] xl:pt-[120px]">
    <RouterLink
      v-for="village in villages"
      :to="'/villages/' + village.ID"
      class="lg:w-[calc(33%-16px)] text-left border px-[24px] py-[12px] rounded-md"
    >
      <article class="flex gap-[12px] items-center">
        <h3>{{ village.name }}</h3>
      </article>
    </RouterLink>
  </section>
</template>

<style scoped></style>
