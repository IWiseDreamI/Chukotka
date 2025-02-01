export type entityTypes = "districts" | "villages" | "guidance";

export interface DistrictForm {
  name: string;
  description: string;
  population: string;
  information: string;
  map: string;
}

export interface VillageForm {
  name: string;
  description: string;
  contacts: string;
  nationality: string;
  population: string;
  organizations: string;
  transport: string;
  connection: string;
  activity: string;
  information: string;
  people: string;
  culture: string;
  socials: string;
  district_id: number;
  coordinates: string;
}

export interface TermForm {
  title: string;
  content: string;
}