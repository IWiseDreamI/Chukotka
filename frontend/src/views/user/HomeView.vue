<script setup lang="ts">
import { Button } from "primevue";
import data from "@/assets/data.json";
import Markdown from "@/components/Markdown.vue";
import { useDistrictStore } from "@/stores/districts";
import { useVillageStore } from "@/stores/villages";
import { District } from "@/entities/interfaces/district";
import { Village } from "@/entities/interfaces/village";
import { ref, computed } from "vue";
import { RouterLink } from "vue-router";

const map: {
  districts: Record<string, string>;
  villages: Record<string, { cx: string; cy: string }>;
} = data.map;

const districtsStore = useDistrictStore();
const districts = await districtsStore.getDistricts();

const villagesStore = useVillageStore();
const villages = await villagesStore.getVillages();

const defaultContent = {
  name: "Сёла Чукотки",
  description: `<b>Добро пожаловать в проект, посвящённый уникальным и самобытным сёлам Чукотского автономного округа!</b></br>
                <br />Чукотка — это суровый, но невероятно красивый край, где раскинулись поселения, в которых живут сильные духом люди. Здесь сохранились традиции коренных народов, переплетаются культура и современность, а природа завораживает своей первозданной мощью.<br />
                <br />На этой карте вы найдёте все сёла Чукотки — от побережья Ледовитого океана до отдалённых районов тундры. Мы собираем истории, фотографии и сведения о каждом населённом пункте, чтобы сохранить и передать знания о жизни в этом удивительном регионе.<br />
                <br />Исследуйте карту, узнавайте больше о сёлах и делитесь своими историями!
`,
};

const activeEntity = ref<District | Village | typeof defaultContent>(
  defaultContent
);
const setActiveEntity = (item: District | Village) => {
  activeEntity.value = item;
};

const isVillage = computed<boolean>(() =>
  activeEntity.value?.name.toLowerCase().includes("район") ? false : true
);

const getVillageCords = (village: Village) => {
  const villagesNames = Object.keys(map.villages);
  const item = villagesNames.find((item) => village.name.includes(item));
  return item ? map.villages[item] : undefined;
};

// Контейнер для позиционирования tooltip (относительно него)
const containerRef = ref<HTMLElement | null>(null);

// Состояние tooltip
const tooltip = ref({
  visible: false,
  x: 0,
  y: 0,
  title: "",
  description: "",
});

// Функция показа tooltip: вычисляет координаты относительно containerRef
const showTooltip = (event: MouseEvent, item: District | Village) => {
  tooltip.value.visible = true;
  if (containerRef.value) {
    const rect = containerRef.value.getBoundingClientRect();
    tooltip.value.x = event.clientX - rect.left + 10;
    tooltip.value.y = event.clientY - rect.top + 10;
  } else {
    tooltip.value.x = event.clientX + 10;
    tooltip.value.y = event.clientY + 10;
  }
  tooltip.value.title = item.name;
  tooltip.value.description = item.description || "Описание отсутствует";
};

const hideTooltip = () => {
  tooltip.value.visible = false;
};
</script>

<template>
  <div ref="containerRef" class="home relative">
    <div
      class="flex items-center justify-center w-[1000px] h-[360px] mc:h-full md:w-full lg:h-[500px] xl:w-[60%]"
    >
      <svg
        class="w-full h-[500px] scale-[0.6] translate-x-[-220px] sm:scale-75 sm:translate-x-[-150px] md:translate-x-0 md:w-[640px] md:scale-100 df:scale-[1.2]"
      >
        <path
          v-for="district in districts"
          :key="district.ID"
          :d="map.districts[district.name]"
          @click="setActiveEntity(district)"
          @mouseover="showTooltip($event, district)"
          @mousemove="showTooltip($event, district)"
          @mouseleave="hideTooltip"
          class="stroke-white fill-[#79553D] hover:fill-[#3D2B1F] duration-300 cursor-pointer"
          :class="activeEntity == district ? 'fill-[#3D2B1F]' : ''"
        />
        <!-- Деревни с tooltip -->
        <g v-for="village in villages" :key="village.ID">
          <circle
            r="6.2"
            stroke="black"
            stroke-width="2"
            :cx="getVillageCords(village)?.cx"
            :cy="getVillageCords(village)?.cy"
            @click="setActiveEntity(village)"
            @mouseover="showTooltip($event, village)"
            @mousemove="showTooltip($event, village)"
            @mouseleave="hideTooltip"
            class="fill-white duration-300 cursor-pointer hover:fill-slate-400"
          />
        </g>
      </svg>
    </div>

    <!-- Tooltip: позиционирован абсолютно относительно контейнера -->
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
      <br />
    </div>

    <div
      class="flex flex-col gap-[24px] xl:w-2/5 text-left pt-[24px] pb-[36px] md:py-0"
    >
      <h1>{{ activeEntity?.name }}</h1>
      <Markdown
        v-if="activeEntity.description"
        :content="activeEntity.description"
      />
      <div class="flex gap-[24px]" v-if="'ID' in activeEntity">
        <RouterLink
          :to="(isVillage ? '/villages/' : '/districts/') + activeEntity.ID"
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
  flex-direction: column;
  justify-content: center;
  min-height: calc(100vh - 80px);

  @media (min-width: 1280px) {
    align-items: center;
    flex-direction: row;
    justify-content: space-between;
  }
}
</style>
