import { defineStore } from 'pinia';
import { ref } from 'vue';
import { api } from '@/entities/api/axios';
import { Term } from '@/entities/interfaces/term';

export const useTermStore = defineStore('term', () => {
  const terms = ref<Term[]>([]);

  const fetchTerms = async () => {
    try {
      const response = await api.get('/terms');
      terms.value = response.data;
    } catch (error) {
      console.error('Failed to fetch terms:', error);
    }
  };

  const getTerms = async () => {
    if (!terms.value?.length) await fetchTerms();
    return terms.value as Term[]
  };

  const getTermById = async (id: number) => {
    if (!terms.value?.length) await fetchTerms();
    return terms.value.find((item) => item.ID === id) as Term
  };

  return { getTerms, fetchTerms, getTermById };
});