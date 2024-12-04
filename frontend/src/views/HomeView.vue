<script setup lang="ts">
import data from "@/assets/data.json";

type District =
  | "bilibino"
  | "anadyr"
  | "levek"
  | "egvekinot"
  | "providence"
  | "chukotka";

const map = data.map;
const districts = Object.keys(map) as District[];
</script>

<template>
  <svg view-box="0 0 1000 1000" class="w-[100%] h-[100vh] scale-75">
    <path
      v-for="district in districts"
      :d="map[district]"
      class="stroke-white fill-[#79553D] hover:fill-[#3D2B1F] duration-300 cursor-pointer"
    />
  </svg>
</template>

<style scoped></style>
