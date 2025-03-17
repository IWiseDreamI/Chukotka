export type entityTypes = "districts" | "villages" | "guidance" | "materials" | "about";

export interface DistrictForm {
  name: string;
  description: string;
  population: string;
  information: string;
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
}

export interface TermForm {
  title: string;
  content: string;
}

export interface MaterialForm {
  title: string;
  description: string;
  content: string;
  category: string;
  author: string;
  publication_date: string;
}

export interface AboutForm {
  content: string;
}
