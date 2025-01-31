import { api } from "./axios";
import { Admin } from "@/entities/interfaces/admin";

interface LoginResponse {
  token: string;
}

export const registerAdmin = async (adminData: Admin): Promise<{ message: string }> => {
  try {
    const response = await api.post<{ message: string }>('/register', adminData);
    return response.data;
  } catch (error) {
    console.error('Failed to register admin:', error);
    throw error;
  }
};

export const loginAdmin = async (credentials: Admin): Promise<LoginResponse> => {
  try {
    const response = await api.post<LoginResponse>('/login', credentials);
    return response.data;
  } catch (error) {
    console.error('Failed to login admin:', error);
    throw error;
  }
};
