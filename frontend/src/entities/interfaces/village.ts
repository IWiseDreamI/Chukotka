import { GormModel } from "./gorm";

export interface Village extends GormModel {
  name: string;
  description?: string;
  contacts?: string;
  nationality?: string;
  population?: string;
  organizations?: string;
  transport?: string;
  connection?: string;
  activity?: string;
  information?: string;
  people?: string;
  culture?: string;
  socials?: string;
  district_id: number;
  coordinates: string;
}