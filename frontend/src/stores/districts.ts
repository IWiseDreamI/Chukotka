import { defineStore } from 'pinia';
import { ref } from 'vue';
import { api } from '@/entities/api/axios';
import { District } from '@/entities/interfaces/district';

export const useDistrictStore = defineStore('district', () => {
  const districts = ref<District[]>();

  const fetchDistricts = async () => {
    try {
      const response = await api.get('/districts');
      districts.value = response.data;
    } catch (error) {
      console.error('Failed to fetch districts:', error);
    }
  };

  const getDistricts = async () => {
    if (!districts.value?.length) await fetchDistricts()
    return districts.value as District[]
  }

  const getDistrictById = async (id: number) => {
    if (!districts.value?.length) await fetchDistricts() 
    return districts.value?.find((item) => item.ID === id) as District
  };


  return { getDistricts, fetchDistricts, getDistrictById };
});