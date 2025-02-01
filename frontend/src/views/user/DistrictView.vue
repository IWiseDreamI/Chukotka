<script setup lang="ts">
import { computed, ref } from "vue";
import { useRoute } from "vue-router";
import Menubar from "primevue/menubar";

import Markdown from "@/components/Markdown.vue";
import { useVillageStore } from "@/stores/villages";
import { useDistrictStore } from "@/stores/districts";

const route = useRoute();
const districtId = computed(() => Number(route.params.id));

const villageStore = useVillageStore();

const districtsStore = useDistrictStore();
const district = await districtsStore.getDistrictById(districtId.value);

const villages = await villageStore.getDistrictVillages(district?.ID as number);

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

const activeVillage = ref(villages[0]);

const villageBar = computed(() => {
  console.log(villages);

  return villages.map((village) => {
    return {
      label: village.name,
      command: () => {
        activeVillage.value = village;
      },
    };
  });
});
</script>

<template>
  <section class="flex justify-between pt-[40px] xl:pt-[120px]">
    <div class="flex flex-col w-[40%] text-left gap-[24px] items-center">
      <img
        :src="'/assets/images/emblems/' + district?.name + '.png'"
        :alt="'Герб' + district?.name"
        class="max-h-[200px] w-auto"
      />
      <h1 class="w-full">{{ district?.name }}</h1>
      <p class="w-full">{{ district?.description }}</p>
    </div>

    <div class="flex flex-col w-[50%] text-left gap-[24px]">
      <Menubar :model="items" />
      <Markdown
        v-if="activeItem == 'about'"
        :content="String(district?.information)"
      />
      <Markdown
        v-if="activeItem == 'population'"
        :content="String(district?.population)"
      />
      <div v-if="activeItem == 'villages'">
        <Menubar :model="villageBar" />
        <RouterLink :to="'/villages/' + activeVillage?.ID"
          >Подробнее</RouterLink
        >
        <Markdown
          class="pt-[24px]"
          :content="String(activeVillage?.information)"
        />
      </div>
    </div>
  </section>
</template>

<style scoped></style>
