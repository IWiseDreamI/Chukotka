import { defineStore } from 'pinia';
import { ref } from 'vue';
import { api } from '@/entities/api/axios';
import { Village } from '@/entities/interfaces/village';

export const useVillageStore = defineStore('village', () => {
  const villages = ref<Village[]>([]);

  const fetchVillages = async () => {
    try {
      const response = await api.get('/villages');
      villages.value = response.data;
    } catch (error) {
      console.error('Failed to fetch villages:', error);
    }
  };

  const getVillages = async () => {
    if (!villages.value?.length) await fetchVillages();
    return villages.value as Village[]
  };

  const getVillageById = async (id: number) => {
    if (!villages.value?.length) await fetchVillages();
    return villages.value.find((item) => item.ID === id) as Village
  };

  const getDistrictVillages = async (id: number) => {
    if (!villages.value?.length) await fetchVillages();
    console.log(villages.value)
    return villages.value.filter((item) => Number(item.districtId) === id) as Village[]
  };

  return { getVillages, fetchVillages, getVillageById, getDistrictVillages };
});
