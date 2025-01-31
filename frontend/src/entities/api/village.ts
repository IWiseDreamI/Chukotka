import { api } from './axios';
import { Village } from '@/entities/interfaces/village';

export const createVillage = async (newVillage: Village) => {
  try {
    const response = await api.post('/admin/villages', newVillage, {
      headers: {
        "Authorization": localStorage.getItem('token')
      }
    });
    return response.data;
  } catch (error) {
    console.error('Failed to create village:', error);
    throw error;
  }
};

export const updateVillage = async (id: number, updatedVillage: Village) => {
  try {
    const response = await api.put(`/admin/villages/${id}`, updatedVillage, {
      headers: {
        "Authorization": localStorage.getItem('token')
      }
    });
    return response.data;
  } catch (error) {
    console.error('Failed to update village:', error);
    throw error;
  }
};

export const deleteVillage = async (id: number) => {
  try {
    await api.delete(`/admin/villages/${id}`, {
      headers: {
        "Authorization": localStorage.getItem('token')
      }
    });
  } catch (error) {
    console.error('Failed to delete village:', error);
    throw error;
  }
};
