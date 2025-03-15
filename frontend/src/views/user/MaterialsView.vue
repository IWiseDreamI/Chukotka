<script setup lang="ts">
import { useMaterialStore } from "@/stores/materials";
import { RouterLink } from "vue-router";

const materialStore = useMaterialStore();

// Загружаем материалы из стора
await materialStore.fetchMaterials();
const materials = await materialStore.getMaterials();

// Если нужно, можно добавить дополнительную обработку, например, сортировку
</script>

<template>
  <section class="flex flex-wrap gap-[24px] w-full pt-[40px] xl:pt-[120px]">
    <RouterLink
      v-for="material in materials"
      :key="material.ID"
      :to="`/materials/${material.ID}`"
      class="lg:w-[calc(33%-16px)] text-left border px-[24px] py-[12px] rounded-md hover:shadow-lg transition-shadow"
    >
      <article class="flex flex-col gap-[12px]">
        <h3 class="text-xl font-bold">{{ material.title }}</h3>
        <p class="text-sm text-gray-600">
          {{ material.category }}
          <span v-if="material.author"> — {{ material.author }}</span>
        </p>
        <!-- Можно добавить превью содержимого, обрезав текст -->
        <p class="text-base line-clamp-3">
          {{ material.content }}
        </p>
      </article>
    </RouterLink>
  </section>
</template>

<style scoped>
/* При необходимости добавьте стили для улучшения отображения карточек */
</style>
