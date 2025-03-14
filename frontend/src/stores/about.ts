// src/stores/about.ts
import { defineStore } from 'pinia';
import { ref } from 'vue';
import { api } from '@/entities/api/axios';
import { About } from '@/entities/interfaces/about';

export const useAboutStore = defineStore('about', () => {
  const about = ref<About | null>(null);

  const fetchAbout = async () => {
    try {
      const response = await api.get('/about');
      about.value = response.data;
    } catch (error) {
      console.error('Failed to fetch about page:', error);
    }
  };

  const getAboutPage = async () => {
    if (!about.value) await fetchAbout();
    return about.value as About;
  };

  return { about, fetchAbout, getAboutPage };
});
