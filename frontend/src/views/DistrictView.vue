<script setup lang="ts">
import { computed, ref } from "vue";
import { useRoute } from "vue-router";
import Menubar from "primevue/menubar";

import { District } from "@/types";
import data from "@/assets/data.json";
import Markdown from "@/components/Markdown.vue";

const route = useRoute();
const district = computed(() => route.params.name as District);
const districtData = computed(() => data.data[district.value]);

const activeItem = ref<"about" | "villages" | "population">("about");

const items = ref([
  {
    label: "О районе",
    icon: "pi pi-info-circle",
    command: () => {
      activeItem.value = "about";
    },
  },
  {
    label: "Сёла",
    icon: "pi pi-home",
    command: () => {
      activeItem.value = "villages";
    },
  },
  {
    label: "Население",
    icon: "pi pi-users",
    command: () => {
      activeItem.value = "population";
    },
  },
]);

type villageType = keyof typeof districtData.value.villages;

const villages = computed(() => {
  return Object.keys(districtData.value.villages) as Array<villageType>;
});
const activeVillage = ref(villages.value[0]);

const villageBar = computed(() =>
  villages.value.map((village: villageType) => {
    return {
      label: village,
      command: () => {
        activeVillage.value = village;
      },
    };
  })
);
</script>

<template>
  <section class="flex pt-[120px] justify-between">
    <div class="flex flex-col w-[40%] text-left gap-[24px] items-center">
      <img
        :src="'/assets/images/emblems/' + district + '.png'"
        :alt="'Герб' + districtData.name"
        class="max-h-[200px] w-auto"
      />
      <h1 class="w-full">{{ districtData.name }}</h1>
      <p class="w-full">{{ districtData.description }}</p>
    </div>

    <div class="flex flex-col w-[50%] text-left gap-[24px]">
      <Menubar :model="items" />
      <Markdown v-if="activeItem == 'about'" :source="districtData.info" />
      <Markdown
        v-if="activeItem == 'population'"
        :source="districtData.population"
      />
      <div v-if="activeItem == 'villages'">
        <Menubar :model="villageBar" />
        <Markdown
          class="pt-[24px]"
          :source="districtData.villages[activeVillage]"
        />
      </div>
    </div>
  </section>
</template>

<style scoped></style>
