<script setup lang="ts">
import { useRoute, useRouter } from "vue-router";
import { InputText, Textarea, Button } from "primevue";
import { computed, ref, onMounted } from "vue";
import Editor from "primevue/editor";
import { FieldInterface, getFormFields } from "@/helpers/form-fields";
import {
  entityTypes,
  DistrictForm,
  VillageForm,
  TermForm,
  MaterialForm,
  AboutForm,
} from "@/types";
import { entityGuard } from "@/helpers/entity-guard";
import { updateEntity } from "@/helpers/update-entity";

import { updateDistrict } from "@/entities/api/district";
import { updateVillage } from "@/entities/api/village";
import { updateTerm } from "@/entities/api/term";
import { updateMaterial } from "@/entities/api/material";
import { updateAbout } from "@/entities/api/about";

import { District } from "@/entities/interfaces/district";
import { Village } from "@/entities/interfaces/village";
import { Term } from "@/entities/interfaces/term";
import { Material } from "@/entities/interfaces/material";
import { About } from "@/entities/interfaces/about";

import { useDistrictStore } from "@/stores/districts";
import { useVillageStore } from "@/stores/villages";
import { useTermStore } from "@/stores/terms";
import { useMaterialStore } from "@/stores/materials";
import { useAboutStore } from "@/stores/about";

// Получаем роут и роутер
const route = useRoute();
const router = useRouter();

// Определяем текущую сущность по параметру
const currentEntity = computed(() => {
  entityGuard(route.params.entity as string);
  return route.params.entity as entityTypes;
});

// Получаем конфигурацию полей формы для выбранной сущности
const formFields = computed<FieldInterface[]>(() =>
  getFormFields(currentEntity.value)
);

// Создаём экземпляры store для загрузки данных
const districtStore = useDistrictStore();
const villageStore = useVillageStore();
const termStore = useTermStore();
const materialStore = useMaterialStore();
const aboutStore = useAboutStore();

// Объявляем ref для модели формы (объединяем типы всех форм)
const formModel = ref<
  DistrictForm | VillageForm | TermForm | MaterialForm | AboutForm | null
>(null);

// Загружаем данные по id (если требуется) или для about – получаем уникальную запись
onMounted(async () => {
  const id = Number(route.params.id);
  if (currentEntity.value === "districts") {
    formModel.value = (await districtStore.getDistrictById(id)) as DistrictForm;
  } else if (currentEntity.value === "villages") {
    formModel.value = (await villageStore.getVillageById(id)) as VillageForm;
  } else if (currentEntity.value === "guidance") {
    formModel.value = (await termStore.getTermById(id)) as TermForm;
  } else if (currentEntity.value === "materials") {
    formModel.value = (await materialStore.getMaterialById(id)) as MaterialForm;
  } else if (currentEntity.value === "about") {
    // Для about id может быть не нужен, запись уникальна
    formModel.value = (await aboutStore.getAboutPage()) as AboutForm;
  }
  console.log("Loaded formModel:", formModel.value);
});

// Создадим вычисляемое свойство для типизированного доступа к форме
const typedFormModel = computed({
  get() {
    switch (currentEntity.value) {
      case "districts":
        return formModel.value as DistrictForm;
      case "villages":
        return formModel.value as VillageForm;
      case "guidance":
        return formModel.value as TermForm;
      case "materials":
        return formModel.value as MaterialForm;
      case "about":
        return formModel.value as AboutForm;
      default:
        return formModel.value;
    }
  },
  set(newValue) {
    formModel.value = newValue;
  },
});

// Объединяем интерфейсы сущностей для передачи в API-методы
type FormModels = District | Village | Term | Material | About;

// Маппинг обработчиков обновления по типу сущности
const entityHandlers = {
  districts: (data: FormModels) =>
    updateDistrict(Number(route.params.id), data as District),
  villages: (data: FormModels) =>
    updateVillage(Number(route.params.id), data as Village),
  guidance: (data: FormModels) =>
    updateTerm(Number(route.params.id), data as Term),
  materials: (data: FormModels) =>
    updateMaterial(Number(route.params.id), data as Material),
  about: (data: FormModels) => updateAbout(data as About),
};

const handleSubmit = async () => {
  if (!formModel.value) return;
  const handler = entityHandlers[currentEntity.value];
  if (!handler) {
    console.error("Unknown entity type:", currentEntity.value);
    return;
  }
  try {
    // Приводим значение формы через unknown к объединённому типу
    const currentFormModel = formModel.value as unknown as FormModels;
    await handler(currentFormModel);
    alert("Сущность изменена");

    updateEntity(currentEntity.value);
    router.push(`/admin/${currentEntity.value}`);
  } catch (error) {
    console.error("Ошибка при обновлении:", error);
  }
};
</script>

<template>
  <!-- Отображаем форму только если данные загружены -->
  <section v-if="formModel && typedFormModel" class="add-entity py-[80px]">
    <h2>Изменить экзэмпляр сущности</h2>
    <form
      class="flex flex-col gap-[12px] w-[800px]"
      @submit.prevent="handleSubmit"
    >
      <div
        v-for="field in formFields"
        :key="field.name"
        class="flex flex-col gap-[8px]"
      >
        <div v-if="field.field === 'input'">
          <label :for="field.name">{{ field.label }}</label>
          <InputText
            v-model="(typedFormModel as any)[field.name]"
            class="w-full"
            :id="field.name"
          />
        </div>
        <div v-if="field.field === 'textarea'">
          <label :for="field.name">{{ field.label }}</label>
          <Textarea
            v-model="(typedFormModel as any)[field.name]"
            class="w-full !py-[0.75em]"
            :id="field.name"
            rows="5"
            autoResize
          />
        </div>
        <div v-if="field.field === 'markdown'">
          <label :for="field.name">{{ field.label }}</label>
          <Editor
            v-model="(typedFormModel as any)[field.name]"
            class="w-full"
            :id="field.name"
            editorStyle="height: 220px"
          />
        </div>
      </div>
      <Button type="submit">Отправить</Button>
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
