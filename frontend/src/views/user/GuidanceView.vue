<script setup lang="ts">
import Column from "primevue/column";
import DataTable from "primevue/datatable";
import { useTermStore } from "@/stores/terms";

const terms = useTermStore();
const items = await terms.getTerms();
</script>

<template>
  <section class="flex flex-wrap gap-[24px] w-full pt-[40px] xl:pt-[120px]">
    <DataTable :value="items" tableStyle="min-width: 50rem">
      <Column field="name" header="Название">
        <template #body="item">
          <RouterLink class="w-full" :to="'/guidance/' + item.data.ID">
            {{ item.data.title }}
          </RouterLink>
        </template>
      </Column>
    </DataTable>
  </section>
</template>

<style scoped></style>
