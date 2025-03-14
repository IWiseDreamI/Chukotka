import { GormModel } from "./gorm";

export interface About extends GormModel {
    content: string;
  }
  