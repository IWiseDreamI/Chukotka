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
  { name: "description", field: "textarea", label: "Описание поселения" },
  { name: "contacts", field: "markdown", label: "Контакты" },
  { name: "nationality", field: "textarea", label: "Национальности" },
  { name: "population", field: "textarea", label: "Население" },
  { name: "organizations", field: "markdown", label: "Организации" },
  { name: "transport", field: "textarea", label: "Транспорт" },
  { name: "connection", field: "textarea", label: "Связи" },
  { name: "activity", field: "textarea", label: "Активности" },
  { name: "information", field: "markdown", label: "Информация" },
  { name: "people", field: "markdown", label: "Люди" },
  { name: "culture", field: "markdown", label: "Культура" },
  { name: "socials", field: "markdown", label: "Соцсети" },
  { name: "districtId", field: "input", label: "Район" },
  { name: "coordinates", field: "textarea", label: "Координаты" },
];

const guidanceForm: FieldInterface[] = [
  { name: "title", field: "input", label: "Термин" },
  { name: "content", field: "markdown", label: "Содержание" },
];

export const getFormFields = (entity: string) => {
  return entity === "districts"
    ? districtForm
    : entity === "villages"
    ? villageForm
    : guidanceForm;
} 
