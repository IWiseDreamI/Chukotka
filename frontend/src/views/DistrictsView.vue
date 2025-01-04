<script setup lang="ts">
import data from "@/assets/data.json";
import { District } from "@/types";

const districtsData = data.data;
const districts = Object.keys(districtsData) as District[];
</script>

<template>
  <section class="flex flex-wrap gap-[24px] pt-[120px] w-full">
    <RouterLink
      :to="'/districts/' + district"
      v-for="district in districts"
      class="lg:w-[calc(33%-16px)] text-left border px-[24px] py-[12px] rounded-md"
    >
      <article class="flex flex-col gap-[24px]">
        <h3>{{ districtsData[district].name }}</h3>
      </article>
    </RouterLink>
  </section>
</template>

<style scoped></style>
