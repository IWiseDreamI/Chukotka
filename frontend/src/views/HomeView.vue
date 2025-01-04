<script setup lang="ts">
import { Button } from "primevue";
import data from "@/assets/data.json";
import { ref } from "vue";
import { District } from "@/types";

const map = data.map;
const districtsData = data.data;
const districts = Object.keys(map) as District[];

const activeDistrict = ref<District>("bilibino");
const setActiveDistrict = (district: District) => {
  activeDistrict.value = district;
};
</script>

<template>
  <div
    class="w-[100%] h-[calc(100vh-80px)] flex justify-end items-center overflow-hidden"
  >
    <svg
      view-box="0 0 1000 1000"
      class="min-w-[1024px] h-[800px] scale-50 absolute left-0 translate-x-[-16%]"
    >
      <path
        v-for="district in districts"
        :d="map[district]"
        @click="setActiveDistrict(district)"
        class="stroke-white fill-[#79553D] hover:fill-[#3D2B1F] duration-300 cursor-pointer"
        :class="activeDistrict == district ? 'fill-[#3D2B1F]' : ''"
      />
    </svg>
    <div class="flex flex-col gap-[24px] w-[50%] text-left">
      <h1>{{ districtsData[activeDistrict].name }}</h1>
      <p>{{ districtsData[activeDistrict].description }}</p>
      <div class="flex gap-[24px]">
        <RouterLink :to="'/districts/' + activeDistrict">
          <Button>Подробнее</Button>
        </RouterLink>
        <RouterLink :to="'/guidance/' + activeDistrict">
          <Button>Энциклопедия</Button>
        </RouterLink>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
