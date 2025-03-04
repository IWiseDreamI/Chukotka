<script setup lang="ts">
import { computed, ref, onMounted, onBeforeUnmount, watch } from "vue";
import { useRoute } from "vue-router";
import Menubar from "primevue/menubar";
import { Button } from "primevue";
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
  return villages.map((village) => {
    return {
      label: village.name,
      command: () => {
        activeVillage.value = village;
      },
    };
  });
});

// Ссылки на компоненты Menubar
const mainMenubar = ref();
const villageMenubar = ref();

// Функция удаления класса p-menubar-mobile
function removeMobileClass() {
  if (mainMenubar.value && mainMenubar.value.$el) {
    mainMenubar.value.$el.classList.remove("p-menubar-mobile");
  }
  if (villageMenubar.value && villageMenubar.value.$el) {
    villageMenubar.value.$el.classList.remove("p-menubar-mobile");
  }
}

watch(activeItem, () => {
  setTimeout(() => {
    removeMobileClass();
  }, 100);
});

// Обработчик resize, который удаляет класс
function handleResize() {
  removeMobileClass();
}

onMounted(() => {
  // Задержка перед удалением класса
  setTimeout(() => {
    removeMobileClass();
  }, 300);
  window.addEventListener("resize", handleResize);
});

onBeforeUnmount(() => {
  window.removeEventListener("resize", handleResize);
});
</script>

<template>
  <section
    class="flex flex-col justify-between pt-[40px] xl:pt-[120px] xl:flex-row"
  >
    <div
      class="flex flex-col w-full mb-[40px] xl:mb-0 xl:w-[40%] text-left gap-[24px] items-center"
    >
      <img
        :src="'/assets/images/emblems/' + district?.name + '.png'"
        :alt="'Герб' + district?.name"
        class="max-h-[200px] w-auto"
      />
      <h1 class="w-full">{{ district?.name }}</h1>
      <p class="w-full">{{ district?.description }}</p>
    </div>

    <div class="flex flex-col w-full xl:w-[50%] text-left gap-[24px]">
      <!-- Основной Menubar с ref -->
      <Menubar ref="mainMenubar" :model="items" />
      <Markdown
        v-if="activeItem == 'about'"
        :content="String(district?.information)"
      />
      <Markdown
        v-if="activeItem == 'population'"
        :content="String(district?.population)"
      />
      <div v-if="activeItem == 'villages'">
        <!-- Menubar для деревень с ref -->
        <Menubar ref="villageMenubar" :model="villageBar" />
        <Markdown
          class="pt-[24px]"
          :content="String(activeVillage?.description)"
        />
        <RouterLink :to="'/villages/' + activeVillage?.ID">
          <Button>Подробнее</Button>
        </RouterLink>
      </div>
    </div>
  </section>
</template>

<style scoped>
/* Дополнительные стили при необходимости */
</style>
