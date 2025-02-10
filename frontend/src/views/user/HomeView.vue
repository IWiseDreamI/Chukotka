<script setup lang="ts">
import { Button } from "primevue";
import data from "@/assets/data.json";
import Markdown from "@/components/Markdown.vue";
import { useDistrictStore } from "@/stores/districts";
import { useVillageStore } from "@/stores/villages";
import { District } from "@/entities/interfaces/district";
import { Village } from "@/entities/interfaces/village";
import { ref, computed } from "vue";

const map: {
  districts: Record<string, string>;
  villages: Record<string, { cx: string; cy: string }>;
} = data.map;
const districtsStore = useDistrictStore();
const districts = await districtsStore.getDistricts();

const villagesStore = useVillageStore();
const villages = await villagesStore.getVillages();

const activeEntity = ref<District | Village>(districts[0]);
const setActiveEntity = (item: District | Village) => {
  activeEntity.value = item;
};

const isVilllage = computed<boolean>(() =>
  activeEntity.value.name.toLowerCase().includes("район") ? false : true
);

const getVillageCords = (village: Village) => {
  const villagesNames = Object.keys(map.villages);
  const item = villagesNames.find((item) => village.name.includes(item));

  return item ? map.villages[item] : undefined;
};
</script>

<template>
  <div class="home">
    <div class="w-[100%] h-[240px] overflow-hidden mc:h-[auto] md:w-[50%]">
      <svg
        class="w-[640px] h-[500px] basis-[500px] translate-x-[-25%] translate-y-[-125px] mc:basis-[auto] mc:translate-x-[-18%] mc:translate-y-[0] md:w-[100%] md:h-[500px] md:translate-x-[0%]"
      >
        <path
          v-for="district in districts"
          :d="map.districts[district.name]"
          @click="setActiveEntity(district)"
          class="stroke-white fill-[#79553D] hover:fill-[#3D2B1F] duration-300 cursor-pointer"
          :class="activeEntity == district ? 'fill-[#3D2B1F]' : ''"
        />
        <circle
          v-for="village in villages"
          r="6.2"
          stroke="black"
          stroke-width="2"
          :cx="getVillageCords(village)?.cx"
          :cy="getVillageCords(village)?.cy"
          @click="setActiveEntity(village)"
          class="fill-white duration-300 cursor-pointer hover:fill-slate-400"
        />
      </svg>
    </div>

    <div
      class="flex flex-col gap-[24px] sm:w-2/5 text-left pt-[24px] pb-[36px] md:py-0"
    >
      <h1>{{ activeEntity.name }}</h1>
      <Markdown :content="String(activeEntity.description)" />
      <div class="flex gap-[24px]">
        <RouterLink
          :to="(isVilllage ? '/villages/' : '/districts/') + activeEntity.ID"
        >
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
