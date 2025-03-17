import { GormModel } from "./gorm";

export interface Material extends GormModel {
    title: string;
    content: string;
    description: string;
    category: string;
    author?: string;
    publication_date: string;
  }
  