export interface FieldInterface {
  name: string;
  field: string;
  label: string;
}

const districtForm: FieldInterface[] = [
  { name: "name", field: "input", label: "Название района" },
  { name: "description", field: "textarea", label: "Описание района" },
  { name: "population", field: "markdown", label: "Население" },
  { name: "information", field: "markdown", label: "Информация" },
  { name: "map", field: "textarea", label: "Очертания района" },
];

const villageForm: FieldInterface[] = [
  { name: "name", field: "input", label: "Название поселения" },
  { name: "description", field: "markdown", label: "Описание поселения" },
  { name: "contacts", field: "markdown", label: "Контакты" },
  { name: "nationality", field: "markdown", label: "Национальности" },
  { name: "population", field: "markdown", label: "Население" },
  { name: "organizations", field: "markdown", label: "Организации" },
  { name: "transport", field: "markdown", label: "Транспорт" },
  { name: "connection", field: "markdown", label: "Связи" },
  { name: "activity", field: "markdown", label: "Активности" },
  { name: "information", field: "markdown", label: "Информация" },
  { name: "people", field: "markdown", label: "Люди" },
  { name: "culture", field: "markdown", label: "Культура" },
  { name: "socials", field: "markdown", label: "Соцсети" },
  { name: "district_id", field: "input", label: "Район" },
  { name: "coordinates", field: "textarea", label: "Координаты" },
];

const guidanceForm: FieldInterface[] = [
  { name: "title", field: "input", label: "Термин" },
  { name: "content", field: "markdown", label: "Содержание" },
];

const materialForm: FieldInterface[] = [
  { name: "title", field: "input", label: "Название материала" },
  { name: "description", field: "textarea", label: "Описание" },
  { name: "content", field: "markdown", label: "Содержание" },
  { name: "category", field: "input", label: "Категория" },
  { name: "author", field: "input", label: "Автор" },
  { name: "publication_date", field: "input", label: "Дата публикации" },
];

const aboutForm: FieldInterface[] = [
  { name: "content", field: "markdown", label: "Контент страницы 'О проекте'" },
];

export const getFormFields = (entity: string): FieldInterface[] => {
  if (entity === "districts") return districtForm;
  else if (entity === "villages") return villageForm;
  else if (entity === "guidance") return guidanceForm;
  else if (entity === "materials") return materialForm;
  else if (entity === "about") return aboutForm;
  else return [];
};
