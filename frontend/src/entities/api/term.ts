import { Term } from '@/entities/interfaces/term';
import { api } from './axios';

export const createTerm = async (newTerm: Term) => {
  try {
    const response = await api.post('/admin/terms', newTerm, {
      headers: {
        "Authorization": localStorage.getItem('token')
      }
    });
    return response.data;
  } catch (error) {
    console.error('Failed to create term:', error);
    throw error;
  }
};

export const updateTerm = async (id: number, updatedTerm: Term) => {
  try {
    const response = await api.put(`/admin/terms/${id}`, updatedTerm, {
      headers: {
        "Authorization": localStorage.getItem('token')
      }
    });
    return response.data;
  } catch (error) {
    console.error('Failed to update term:', error);
    throw error;
  }
};

export const deleteTerm = async (id: number) => {
  try {
    await api.delete(`/admin/terms/${id}`, {
      headers: {
        "Authorization": localStorage.getItem('token')
      }
    });
  } catch (error) {
    console.error('Failed to delete term:', error);
    throw error;
  }
};