import { GormModel } from "./gorm";
import { Village } from "./village";

export interface District extends GormModel {
  name: string;
  description?: string;
  population?: string;
  information?: string;
  villages?: Village[];
}
