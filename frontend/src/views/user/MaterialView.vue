<script setup lang="ts">
import { onMounted, ref } from "vue";
import { useRoute } from "vue-router";
import { useMaterialStore } from "@/stores/materials";
import Markdown from "@/components/Markdown.vue";
import type { Material } from "@/entities/interfaces/material";

const route = useRoute();
const materialStore = useMaterialStore();
const material = ref<Material | null>(null);

onMounted(async () => {
  const id = Number(route.params.id);
  material.value = await materialStore.getMaterialById(id);
  console.log("Loaded material:", material.value);
});
</script>

<template>
  <section class="material-detail" v-if="material">
    <h1 class="title">{{ material.title }}</h1>
    <div class="meta">
      <span class="category">{{ material.category }}</span>
      <span v-if="material.author" class="author">
        – {{ material.author }}</span
      >
      <span v-if="material.publication_date" class="date">
        ({{ material.publication_date }})</span
      >
    </div>
    <Markdown :content="material.content" />
  </section>
  <div v-else>Загрузка...</div>
</template>

<style scoped>
.material-detail {
  margin-top: 40px;
  gap: 12px;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  padding: 20px;
}
.title {
  font-size: 2rem;
  margin-bottom: 10px;
}
.meta {
  color: #666;
  margin-bottom: 20px;
  font-size: 0.9rem;
}
</style>
