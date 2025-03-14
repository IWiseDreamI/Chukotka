<script setup lang="ts">
import { useRoute, useRouter } from "vue-router";
import { InputText, Textarea, Button } from "primevue";
import { computed, ref } from "vue";
import Editor from "primevue/editor";
import { FieldInterface, getFormFields } from "@/helpers/form-fields";
import {
  DistrictForm,
  entityTypes,
  TermForm,
  VillageForm,
  MaterialForm,
  AboutForm,
} from "@/types";
import { entityGuard } from "@/helpers/entity-guard";

import { createDistrict } from "@/entities/api/district";
import { createVillage } from "@/entities/api/village";
import { createTerm } from "@/entities/api/term";
import { createMaterial } from "@/entities/api/material";
// Для страницы "О нас" используется функция обновления, так как запись уникальна
import { updateAbout } from "@/entities/api/about";

import { District } from "@/entities/interfaces/district";
import { Village } from "@/entities/interfaces/village";
import { Term } from "@/entities/interfaces/term";
import { Material } from "@/entities/interfaces/material";
import { About } from "@/entities/interfaces/about";

import { updateEntity } from "@/helpers/update-entity";

const route = useRoute();
const router = useRouter();

const entity = computed(() => {
  entityGuard(route.params.entity as string);
  return route.params.entity as entityTypes;
});

// Вычисляем поля формы в зависимости от сущности
const formFields = computed<FieldInterface[]>(() =>
  getFormFields(entity.value)
);

// Определяем отдельные модели для каждой сущности
const districtFormModel = ref<DistrictForm>({
  name: "",
  description: "",
  population: "",
  information: "",
});

const villageFormModel = ref<VillageForm>({
  name: "",
  description: "",
  contacts: "",
  nationality: "",
  population: "",
  organizations: "",
  transport: "",
  connection: "",
  activity: "",
  information: "",
  people: "",
  culture: "",
  socials: "",
  district_id: 0,
});

const termFormModel = ref<TermForm>({
  title: "",
  content: "",
});

const materialFormModel = ref<MaterialForm>({
  title: "",
  content: "",
  category: "",
  author: "",
  publication_date: "",
});

const aboutFormModel = ref<AboutForm>({
  content: "",
});

// computed, возвращающий нужный ref модели в зависимости от типа сущности
const formModelRef = computed(() => {
  switch (entity.value) {
    case "districts":
      return districtFormModel;
    case "villages":
      return villageFormModel;
    case "guidance":
      return termFormModel;
    case "materials":
      return materialFormModel;
    case "about":
      return aboutFormModel;
    default:
      return termFormModel;
  }
});

// Объединяем типы форм в один union
type FormModels =
  | DistrictForm
  | VillageForm
  | TermForm
  | MaterialForm
  | AboutForm;
type FormFieldsNames = keyof (DistrictForm &
  VillageForm &
  TermForm &
  MaterialForm &
  AboutForm);

// Маппинг обработчиков. Для about используется updateAbout, так как запись уже существует.
const entityHandlers = {
  districts: (data: FormModels) => createDistrict(data as District),
  villages: (data: FormModels) => createVillage(data as Village),
  guidance: (data: FormModels) => createTerm(data as Term),
  materials: (data: FormModels) => createMaterial(data as Material),
  about: (data: FormModels) => updateAbout(data as About),
};

// Заголовок формы и текст кнопки зависят от сущности
const formTitle = computed(() =>
  entity.value === "about"
    ? "Обновить страницу 'О нас'"
    : "Добавить экзэмпляр сущности"
);
const buttonText = computed(() =>
  entity.value === "about" ? "Обновить" : "Отправить"
);

const handleSubmit = async () => {
  const handler = entityHandlers[entity.value];
  if (!handler) {
    console.error("Unknown entity type:", entity.value);
    return;
  }
  try {
    // Приводим значение формы через unknown к объединённому типу
    const currentFormModel = formModelRef.value as unknown as FormModels;
    await handler(currentFormModel);
    let entityName = "";
    if (entity.value === "districts" || entity.value === "villages") {
      entityName = (formModelRef.value as unknown as { name: string }).name;
    } else if (entity.value === "guidance" || entity.value === "materials") {
      entityName = (formModelRef.value as unknown as { title: string }).title;
    } else if (entity.value === "about") {
      entityName = "Страница 'О проекте'";
    }
    alert(`${entityName} успешно обновлена`);

    updateEntity(entity.value);
    router.push(`/admin/${entity.value}`);
    window.location.reload();
  } catch (error) {
    console.error("Ошибка при отправке данных:", error);
  }
};
</script>

<template>
  <section class="add-entity py-[80px]">
    <h2>{{ formTitle }}</h2>
    <form class="flex flex-col gap-[12px] w-[800px]">
      <div
        class="flex flex-col gap-[8px]"
        v-for="field in formFields"
        :key="field.name"
      >
        <div v-if="field.field === 'input'">
          <label :for="field.name">{{ field.label }}</label>
          <InputText
            v-model="(formModelRef.value as Record<string, string | undefined>)[field.name as FormFieldsNames]"
            class="w-full"
            :id="field.name"
          />
        </div>
        <div v-if="field.field === 'textarea'">
          <label :for="field.name">{{ field.label }}</label>
          <Textarea
            v-model="(formModelRef.value as Record<string, string | undefined>)[field.name as FormFieldsNames]"
            class="w-full !py-[0.75em]"
            :id="field.name"
            rows="5"
            autoResize
          />
        </div>
        <div v-if="field.field === 'markdown'">
          <label :for="field.name">{{ field.label }}</label>
          <Editor
            v-model="(formModelRef.value as Record<string, string | undefined>)[field.name as FormFieldsNames]"
            class="w-full"
            :id="field.name"
            editorStyle="height: 220px"
          />
        </div>
      </div>
      <Button @click="handleSubmit">{{ buttonText }}</Button>
    </form>
  </section>
</template>

<style scoped lang="scss">
.add-entity {
  gap: 20px;
  display: flex;
  text-align: left;
  flex-direction: column;
  min-height: calc(100vh - 80px);
}
</style>
