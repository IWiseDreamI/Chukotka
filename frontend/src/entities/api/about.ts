import { About } from '@/entities/interfaces/about';
import { api } from './axios';

export const updateAbout = async (aboutContent: About) => {
  try {
    const response = await api.put('/admin/about', aboutContent, {
      headers: {
        "Authorization": localStorage.getItem('token')
      }
    });
    return response.data;
  } catch (error) {
    console.error('Failed to update about page:', error);
    throw error;
  }
};
