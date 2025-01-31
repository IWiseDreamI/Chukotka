<script setup lang="ts">
import { Button, Column, DataTable } from "primevue";

import { computed, ref, onMounted } from "vue";

import { useRoute } from "vue-router";
import { useTermStore } from "@/stores/terms";
import { useVillageStore } from "@/stores/villages";
import { useDistrictStore } from "@/stores/districts";

import { Term } from "@/entities/interfaces/term";
import { Village } from "@/entities/interfaces/village";
import { District } from "@/entities/interfaces/district";

import { deleteTerm } from "@/entities/api/term";
import { deleteVillage } from "@/entities/api/village";
import { deleteDistrict } from "@/entities/api/district";

import { entityTypes } from "@/types";
import { updateEntity } from "@/helpers/update-entity";

const terms = useTermStore();
const villages = useVillageStore();
const districts = useDistrictStore();

const route = useRoute();

const path = route.fullPath;
const entity = computed(() => route.params.entity as entityTypes);

const items = ref();

onMounted(async () => {
  items.value =
    entity.value === "districts"
      ? await districts.getDistricts()
      : entity.value === "villages"
      ? await villages.getVillages()
      : await terms.getTerms();
});

const deleteEntity = (id: number) => {
  const item = items.value.find((i: District | Village | Term) => i.ID == id);

  const entityName = (item as District | Village).name
    ? (item as District | Village).name
    : (item as Term).title;

  confirm("Вы действительно хотите удалить " + entityName);

  const del =
    entity.value === "districts"
      ? deleteDistrict(id)
      : entity.value === "villages"
      ? deleteVillage(id)
      : deleteTerm(id);

  del.then(() => {
    alert(entityName + " удалено");
    updateEntity(entity.value);
  });
};
</script>

<template>
  <section class="admin py-[80px]">
    <RouterLink to="/admin">
      <h1>Администраторская панель</h1>
    </RouterLink>

    <DataTable :value="items" tableStyle="min-width: 50rem">
      <Column header="Название">
        <template #body="item">
          <p>{{ item.data.title ? item.data.title : item.data.name }}</p>
        </template>
      </Column>
      <Column field="ID" header="Изменить" style="width: 60px">
        <template #body="item">
          <RouterLink :to="path + '/edit/' + item.data.ID">
            <i class="w-full pi pi-pencil text-center cursor-pointer" />
          </RouterLink>
        </template>
      </Column>
      <Column field="ID" header="Удалить" style="width: 60px">
        <template #body="item">
          <i
            class="w-full pi pi-trash text-center cursor-pointer"
            @click="
              () => {
                deleteEntity(item.data.ID);
              }
            "
          />
        </template>
      </Column>
    </DataTable>
    <RouterLink
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
