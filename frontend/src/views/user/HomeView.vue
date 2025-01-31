<script setup lang="ts">
import { Button } from "primevue";
import data from "@/assets/data.json";
import { useDistrictStore } from "@/stores/districts";
import { District } from "@/entities/interfaces/district";
import { ref } from "vue";

const map = data.map;
const districtsStore = useDistrictStore();
const districts = await districtsStore.getDistricts();

const activeDistrict = ref<District>(districts[0]);
const setActiveDistrict = (district: District) => {
  activeDistrict.value = district;
};
</script>

<template>
  <div class="home">
    <div class="w-[100%] h-[240px] overflow-hidden mc:h-[auto] md:w-[50%]">
      <svg
        class="w-[640px] h-[500px] scale-[0.46] basis-[500px] translate-x-[-25%] translate-y-[-125px] mc:basis-[auto] mc:scale-[0.6] mc:translate-x-[-18%] mc:translate-y-[0] md:w-[100%] md:scale-100 md:h-[500px] md:translate-x-[0%]"
      >
        <path
          v-for="district in districts"
          :d="map[district.name as keyof typeof map]"
          @click="setActiveDistrict(district)"
          class="stroke-white fill-[#79553D] hover:fill-[#3D2B1F] duration-300 cursor-pointer"
          :class="activeDistrict == district ? 'fill-[#3D2B1F]' : ''"
        />
      </svg>
    </div>

    <div
      class="flex flex-col gap-[24px] sm:w-2/5 text-left pt-[24px] pb-[36px] md:py-0"
    >
      <h1>{{ activeDistrict.name }}</h1>
      <p>{{ activeDistrict.description }}</p>
      <div class="flex gap-[24px]">
        <RouterLink :to="'/districts/' + activeDistrict.ID">
          <Button>Подробнее</Button>
        </RouterLink>
        <RouterLink to="/guidance/">
          <Button>Энциклопедия</Button>
        </RouterLink>
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss">
.home {
  width: 100%;
  display: flex;
  overflow: hidden;
  flex-direction: column;
  justify-content: center;
  min-height: calc(100vh - 80px);

  @media (min-width: 768px) {
    align-items: center;
    flex-direction: row;
    justify-content: space-between;
  }
}
</style>
