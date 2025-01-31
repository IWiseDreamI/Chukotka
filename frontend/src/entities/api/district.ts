import { District } from '@/entities/interfaces/district';
import { api } from './axios';

export const createDistrict = async (newDistrict: District) => {
  try {
    const response = await api.post('/admin/districts', newDistrict, {
      headers: {
        "Authorization": localStorage.getItem('token')
      }
    });
    return response.data;
  } catch (error) {
    console.error('Failed to create district:', error);
    throw error;
  }
};

export const updateDistrict = async (id: number, updatedDistrict: District) => {
  try {
    const response = await api.put(`/admin/districts/${id}`, updatedDistrict, {
      headers: {
        "Authorization": localStorage.getItem('token')
      }
    });
    return response.data;
  } catch (error) {
    console.error('Failed to update district:', error);
    throw error;
  }
};

export const deleteDistrict = async (id: number) => {
  try {
    await api.delete(`/admin/districts/${id}`, {
      headers: {
        "Authorization": localStorage.getItem('token')
      }
    });
  } catch (error) {
    console.error('Failed to delete district:', error);
    throw error;
  }
};
