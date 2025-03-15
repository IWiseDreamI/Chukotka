<script setup lang="ts">
import data from "@/assets/data.json";
import { District } from "@/entities/interfaces/district";
import { Village } from "@/entities/interfaces/village";
import { useDistrictStore } from "@/stores/districts";
import { useVillageStore } from "@/stores/villages";
import { ref } from "vue";
import { RouterLink } from "vue-router";

// Ваша карта
const map: {
  districts: Record<string, string>;
  villages: Record<string, { cx: string; cy: string; tx: string; ty: string }>;
} = data.map;

// Получаем районы и деревни из стора
const districtsStore = useDistrictStore();
const districts = await districtsStore.getDistricts();

const villagesStore = useVillageStore();
const villages = await villagesStore.getVillages();

// Активный элемент (район/деревня)
const activeEntity = ref<District | Village>(districts[0]);
const setActiveEntity = (item: District | Village) => {
  activeEntity.value = item;
};

// Функция поиска координат для деревни
const getVillageCords = (village: Village) => {
  const villagesNames = Object.keys(map.villages);
  const item = villagesNames.find((item) => village.name.includes(item));
  return item ? map.villages[item] : undefined;
};

// Tooltip — показывается при наведении на деревню
const containerRef = ref<HTMLElement | null>(null);

const tooltip = ref({
  visible: false,
  x: 0,
  y: 0,
  title: "",
  description: "",
});

const showTooltip = (event: MouseEvent, village: Village) => {
  tooltip.value.visible = true;
  if (containerRef.value) {
    const rect = containerRef.value.getBoundingClientRect();
    tooltip.value.x = event.clientX - rect.left + 10;
    tooltip.value.y = event.clientY - rect.top + 10;
  } else {
    tooltip.value.x = event.clientX + 10;
    tooltip.value.y = event.clientY + 10;
  }
  tooltip.value.title = village.name;
  tooltip.value.description = village.description || "Описание отсутствует";
};

const hideTooltip = () => {
  tooltip.value.visible = false;
};
</script>

<template>
  <div ref="containerRef" class="relative mx-auto max-w-[900px] w-full p-4">
    <div class="relative w-full aspect-[13/10]">
      <svg
        viewBox="0 0 650 500"
        preserveAspectRatio="xMidYMid meet"
        class="absolute top-0 left-0 w-full h-full"
      >
        <path
          v-for="district in districts"
          :key="district.ID"
          :d="map.districts[district.name]"
          class="stroke-white fill-[#79553D] hover:fill-[#3D2B1F] duration-300 cursor-pointer"
          @click="setActiveEntity(district)"
        />
        <g
          v-for="village in villages"
          :key="village.ID"
          @mouseover="showTooltip($event, village)"
          @mousemove="showTooltip($event, village)"
          @mouseleave="hideTooltip"
        >
          <RouterLink :to="'/villages/' + village.ID">
            <circle
              r="6.2"
              stroke="black"
              stroke-width="2"
              :cx="getVillageCords(village)?.cx"
              :cy="getVillageCords(village)?.cy"
              class="fill-white duration-300 cursor-pointer hover:fill-slate-400"
              @click="setActiveEntity(village)"
            />
            <g v-if="getVillageCords(village)">
              <text
                :x="getVillageCords(village)?.tx"
                :y="getVillageCords(village)?.ty"
                dx="5"
                dy="3"
                stroke="black"
                fill="white"
                stroke-width="0.3"
                font-size="14"
                font-family="sans-serif"
                letter-spacing="0.3"
              >
                {{ village.name }}
              </text>
            </g>
          </RouterLink>
        </g>
      </svg>
    </div>

    <!-- Tooltip -->
    <div
      v-if="tooltip.visible"
      class="absolute bg-white border border-gray-300 p-2 rounded shadow-lg text-sm"
      :style="{
        top: tooltip.y + 'px',
        left: tooltip.x + 'px',
        whiteSpace: 'pre-wrap',
      }"
    >
      <strong>{{ tooltip.title }}</strong>
    </div>
  </div>
</template>

<style scoped>
g {
  cursor: pointer;
}

text {
  background: #000;
}
</style>
