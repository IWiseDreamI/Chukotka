// src/stores/materials.ts
import { defineStore } from 'pinia';
import { ref } from 'vue';
import { api } from '@/entities/api/axios';
import { Material } from '@/entities/interfaces/material';

export const useMaterialStore = defineStore('material', () => {
  const materials = ref<Material[]>([]);

  const fetchMaterials = async () => {
    try {
      const response = await api.get('/materials');
      materials.value = response.data;
    } catch (error) {
      console.error('Failed to fetch materials:', error);
    }
  };

  const getMaterials = async () => {
    if (!materials.value?.length) await fetchMaterials();
    return materials.value;
  };

  const getMaterialById = async (id: number) => {
    if (!materials.value?.length) await fetchMaterials();
    return materials.value.find((item) => item.ID === id) as Material;
  };

  return { materials, fetchMaterials, getMaterials, getMaterialById };
});
