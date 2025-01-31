<script setup lang="ts">
import { useDistrictStore } from "@/stores/districts";

const districtsStore = useDistrictStore();
districtsStore.fetchDistricts();

const districts = await districtsStore.getDistricts();
</script>

<template>
  <section class="flex flex-wrap gap-[24px] w-full pt-[40px] xl:pt-[120px]">
    <RouterLink
      :to="'/districts/' + district.ID"
      v-for="district in districts"
      class="lg:w-[calc(33%-16px)] text-left border px-[24px] py-[12px] rounded-md"
    >
      <article class="flex gap-[12px] items-center">
        <img
          class="w-[96px] h-[96px]"
          :src="'/assets/images/emblems/' + district.name + '.png'"
          :alt="'Герб ' + district.name"
        />
        <h3>{{ district.name }}</h3>
      </article>
    </RouterLink>
  </section>
</template>

<style scoped></style>
