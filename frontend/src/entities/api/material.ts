// src/entities/api/material.ts
import { Material } from '@/entities/interfaces/material';
import { api } from './axios';

export const getMaterials = async (): Promise<Material[]> => {
  try {
    const response = await api.get('/materials');
    return response.data;
  } catch (error) {
    console.error('Failed to fetch materials:', error);
    throw error;
  }
};

export const getMaterialById = async (id: number): Promise<Material> => {
  try {
    const response = await api.get(`/materials/${id}`);
    return response.data;
  } catch (error) {
    console.error(`Failed to fetch material with id ${id}:`, error);
    throw error;
  }
};

export const createMaterial = async (newMaterial: Material): Promise<Material> => {
  try {
    const response = await api.post('/admin/materials', newMaterial, {
      headers: {
        Authorization: localStorage.getItem('token')
      }
    });
    return response.data;
  } catch (error) {
    console.error('Failed to create material:', error);
    throw error;
  }
};

export const updateMaterial = async (id: number, updatedMaterial: Material): Promise<Material> => {
  try {
    const response = await api.put(`/admin/materials/${id}`, updatedMaterial, {
      headers: {
        Authorization: localStorage.getItem('token')
      }
    });
    return response.data;
  } catch (error) {
    console.error('Failed to update material:', error);
    throw error;
  }
};

export const deleteMaterial = async (id: number): Promise<void> => {
  try {
    await api.delete(`/admin/materials/${id}`, {
      headers: {
        Authorization: localStorage.getItem('token')
      }
    });
  } catch (error) {
    console.error('Failed to delete material:', error);
    throw error;
  }
};
