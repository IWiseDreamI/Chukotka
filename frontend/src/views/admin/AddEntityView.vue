<script setup lang="ts">
import { useRoute, useRouter } from "vue-router";
import { InputText, Textarea, Button } from "primevue";
import { computed, ref } from "vue";
import Editor from "primevue/editor";
import { FieldInterface, getFormFields } from "@/helpers/form-fields";
import { DistrictForm, entityTypes, TermForm, VillageForm } from "@/types";
import { entityGuard } from "@/helpers/entity-guard";

import { createDistrict } from "@/entities/api/district";
import { createVillage } from "@/entities/api/village";
import { createTerm } from "@/entities/api/term";

import { District } from "@/entities/interfaces/district";
import { Village } from "@/entities/interfaces/village";
import { Term } from "@/entities/interfaces/term";

import { updateEntity } from "@/helpers/update-entity";

const route = useRoute();
const router = useRouter();

const entity = computed(() => {
  entityGuard(route.params.entity as string);
  return route.params.entity as entityTypes;
});

const formFields = computed<FieldInterface[]>(() =>
  getFormFields(entity.value)
);

const districtFormModel = ref<DistrictForm>({
  name: "",
  description: "",
  population: "",
  information: "",
  map: "",
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
  coordinates: "",
});

const termFormModel = ref<TermForm>({
  title: "",
  content: "",
});

const formModel = ref(
  entity.value === "districts"
    ? districtFormModel
    : entity.value === "villages"
    ? villageFormModel
    : termFormModel
);

type FormFieldsNames = keyof (DistrictForm & VillageForm & TermForm);
type FormModels = District | Village | Term;

const entityHandlers = {
  districts: (data: FormModels) => createDistrict(data as District),
  villages: (data: FormModels) => createVillage(data as Village),
  guidance: (data: FormModels) => createTerm(data as Term),
};

const handleSubmit = () => {
  const handler = entityHandlers[entity.value];
  if (handler) {
    handler(formModel.value as FormModels).then(() => {
      const entityName =
        (formModel.value as Village | District).name ||
        (formModel.value as Term).title;

      alert(entityName + " добавлено");

      updateEntity(entity.value);
      router.push(`/admin/${entity.value}`);
      window.location.reload();
    });
  } else {
    console.error("Unknown entity type:", entity);
  }
};
</script>

<template>
  <section class="add-entity py-[80px]">
    <h2>Добавить экзэмпляр сущности</h2>
    <form class="flex flex-col gap-[12px] w-[800px]">
      <div class="flex flex-col gap-[8px]" v-for="field in formFields">
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
      <Button @click="handleSubmit">Отправить</Button>
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
