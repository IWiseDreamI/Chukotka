import { GormModel } from "./gorm";

export interface Term extends GormModel {
  title: string;
  content: string;
}