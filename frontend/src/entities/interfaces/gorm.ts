export interface GormModel {
  ID: number;
  createdAt: string; // Date in string format
  updatedAt: string; // Date in string format
  deletedAt?: string | null; // Optional field
}