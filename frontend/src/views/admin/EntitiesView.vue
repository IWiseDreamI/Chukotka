<script setup lang="ts">
import { Button, Column, DataTable } from "primevue";
import { computed, ref, onMounted } from "vue";
import { useRoute } from "vue-router";

import { useTermStore } from "@/stores/terms";
import { useVillageStore } from "@/stores/villages";
import { useDistrictStore } from "@/stores/districts";
import { useMaterialStore } from "@/stores/materials";
import { useAboutStore } from "@/stores/about";

import { Term } from "@/entities/interfaces/term";
import { Village } from "@/entities/interfaces/village";
import { District } from "@/entities/interfaces/district";
import { Material } from "@/entities/interfaces/material";
import { About } from "@/entities/interfaces/about";

import { deleteTerm } from "@/entities/api/term";
import { deleteVillage } from "@/entities/api/village";
import { deleteDistrict } from "@/entities/api/district";
import { deleteMaterial } from "@/entities/api/material";

import { entityTypes } from "@/types";
import { updateEntity } from "@/helpers/update-entity";

const districts = useDistrictStore();
const villages = useVillageStore();
const terms = useTermStore();
const materialStore = useMaterialStore();
const aboutStore = useAboutStore();

const route = useRoute();
const path = route.fullPath;
const entity = computed(() => route.params.entity as entityTypes);

const items = ref<Array<District | Village | Term | Material | About>>([]);

onMounted(async () => {
  if (entity.value === "districts") {
    items.value = await districts.getDistricts();
  } else if (entity.value === "villages") {
    items.value = await villages.getVillages();
  } else if (entity.value === "guidance") {
    items.value = await terms.getTerms();
  } else if (entity.value === "materials") {
    items.value = await materialStore.getMaterials();
  } else if (entity.value === "about") {
    // Для about запись уникальна; оборачиваем её в массив для DataTable
    const aboutItem = await aboutStore.getAboutPage();
    items.value = aboutItem ? [aboutItem] : [];
  }
});

const deleteEntity = (id: number) => {
  // Для about удаление недопустимо
  if (entity.value === "about") {
    alert("Удаление записи 'О нас' недопустимо");
    return;
  }
  const item = items.value.find(
    (i: District | Village | Term | Material | About) => i.ID === id
  );
  if (!item) return;

  let entityName = "";
  // Для guidance и materials используем поле title, для остальных – name
  if (entity.value === "guidance" || entity.value === "materials") {
    entityName = (item as Term | Material).title;
  } else {
    entityName = (item as District | Village).name;
  }

  if (!confirm("Вы действительно хотите удалить " + entityName + "?")) return;

  let delPromise;
  if (entity.value === "districts") {
    delPromise = deleteDistrict(id);
  } else if (entity.value === "villages") {
    delPromise = deleteVillage(id);
  } else if (entity.value === "guidance") {
    delPromise = deleteTerm(id);
  } else if (entity.value === "materials") {
    delPromise = deleteMaterial(id);
  }

  if (delPromise) {
    delPromise
      .then(() => {
        alert(entityName + " удалено");
        updateEntity(entity.value);
      })
      .catch((error) => console.error("Ошибка при удалении", error));
  }
};
</script>

<template>
  <section class="admin py-[80px]">
    <RouterLink to="/admin">
      <h1>Администраторская панель</h1>
    </RouterLink>

    <DataTable :value="items" tableStyle="min-width: 50rem">
      <Column header="Название">
        <template #body="slotProps">
          <p v-if="entity !== 'about'">
            {{
              slotProps.data?.title
                ? slotProps.data.title
                : slotProps.data?.name
            }}
          </p>
          <p v-else>О проекте</p>
        </template>
      </Column>
      <Column header="Изменить" style="width: 60px">
        <template #body="slotProps">
          <RouterLink :to="path + '/edit/' + slotProps.data.ID">
            <i class="w-full pi pi-pencil text-center cursor-pointer" />
          </RouterLink>
        </template>
      </Column>
      <!-- Кнопка удаления не отображается для about -->
      <Column v-if="entity !== 'about'" header="Удалить" style="width: 60px">
        <template #body="slotProps">
          <i
            class="w-full pi pi-trash text-center cursor-pointer"
            @click="() => deleteEntity(slotProps.data.ID)"
          />
        </template>
      </Column>
    </DataTable>
    <!-- Кнопка "Добавить" не отображается для about -->
    <RouterLink
      v-if="entity !== 'about'"
      :to="'/admin/' + entity + '/add'"
      class="flex gap-[24px] items-center"
    >
      <Button>
        <p>Добавить</p>
        <i class="pi pi-plus" />
      </Button>
    </RouterLink>
  </section>
</template>

<style scoped lang="scss">
.admin {
  gap: 32px;
  display: flex;
  align-items: start;
  flex-direction: column;
  min-height: calc(100vh - 80px);
}

.card {
  cursor: pointer;
  padding: 12px;
  box-shadow: 0 0 8px #0000002f;

  .pi {
    font-size: 80px;
  }
}
</style>
