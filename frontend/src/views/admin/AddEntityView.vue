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

// Получаем конфигурацию полей формы для выбранной сущности
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

// Используем старый подход: присваиваем formModel равным нужной модели
const formModel = ref(
  entity.value === "districts"
    ? districtFormModel.value
    : entity.value === "villages"
    ? villageFormModel.value
    : entity.value === "guidance"
    ? termFormModel.value
    : entity.value === "materials"
    ? materialFormModel.value
    : entity.value === "about"
    ? aboutFormModel.value
    : termFormModel.value
);

type FormFieldsNames = keyof (DistrictForm &
  VillageForm &
  TermForm &
  MaterialForm &
  AboutForm);
type FormModels = District | Village | Term | Material | About;

const entityHandlers = {
  districts: (data: FormModels) => createDistrict(data as District),
  villages: (data: FormModels) => createVillage(data as Village),
  guidance: (data: FormModels) => createTerm(data as Term),
  materials: (data: FormModels) => createMaterial(data as Material),
  about: (data: FormModels) => updateAbout(data as About),
};

const formTitle = computed(() =>
  entity.value === "about"
    ? "Обновить страницу 'О нас'"
    : "Добавить экзэмпляр сущности"
);
const buttonText = computed(() =>
  entity.value === "about" ? "Обновить" : "Отправить"
);

const handleSubmit = () => {
  const handler = entityHandlers[entity.value];
  if (!handler) {
    console.error("Unknown entity type:", entity.value);
    return;
  }
  // Передаем formModel.value напрямую – оно является обычным объектом, а не реактивной обёрткой
  handler(formModel.value as FormModels)
    .then(() => {
      const entityName =
        (formModel.value as { name?: string }).name ||
        (formModel.value as { title?: string }).title ||
        (entity.value === "about" ? "Страница 'О нас'" : "");
      alert(entityName + " добавлено");

      updateEntity(entity.value);
      router.push(`/admin/${entity.value}`);
      window.location.reload();
    })
    .catch((error) => {
      console.error("Ошибка при отправке данных:", error);
    });
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
            v-model="(formModel as Record<string, string | undefined>)[field.name as FormFieldsNames]"
            class="w-full"
            :id="field.name"
          />
        </div>
        <div v-if="field.field === 'textarea'">
          <label :for="field.name">{{ field.label }}</label>
          <Textarea
            v-model="(formModel as Record<string, string | undefined>)[field.name as FormFieldsNames]"
            class="w-full !py-[0.75em]"
            :id="field.name"
            rows="5"
            autoResize
          />
        </div>
        <div v-if="field.field === 'markdown'">
          <label :for="field.name">{{ field.label }}</label>
          <Editor
            v-model="(formModel as Record<string, string | undefined>)[field.name as FormFieldsNames]"
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
