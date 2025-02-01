<script setup lang="ts">
import { useRoute, useRouter } from "vue-router";
import { InputText, Textarea, Button } from "primevue";
import { computed, ref, onMounted } from "vue";
import Editor from "primevue/editor";
import { FieldInterface, getFormFields } from "@/helpers/form-fields";
import { entityTypes, DistrictForm, VillageForm, TermForm } from "@/types";
import { entityGuard } from "@/helpers/entity-guard";
import { updateEntity } from "@/helpers/update-entity";

import { updateDistrict } from "@/entities/api/district";
import { updateVillage } from "@/entities/api/village";
import { updateTerm } from "@/entities/api/term";

import { District } from "@/entities/interfaces/district";
import { Village } from "@/entities/interfaces/village";
import { Term } from "@/entities/interfaces/term";

import { useTermStore } from "@/stores/terms";
import { useVillageStore } from "@/stores/villages";
import { useDistrictStore } from "@/stores/districts";

const route = useRoute();
const router = useRouter();

const entity = computed(() => {
  entityGuard(route.params.entity as string);
  return route.params.entity as entityTypes;
});

const formFields = computed<FieldInterface[]>(() =>
  getFormFields(entity.value)
);

const terms = useTermStore();
const villages = useVillageStore();
const districts = useDistrictStore();

const districtFormModel = ref<DistrictForm>();
const villageFormModel = ref<VillageForm>();
const termFormModel = ref<TermForm>();

const formModel = ref(
  entity.value === "districts"
    ? districtFormModel.value
    : entity.value === "villages"
    ? villageFormModel.value
    : termFormModel.value
);

onMounted(async () => {
  const ID = Number(route.params.id);

  formModel.value =
    entity.value === "districts"
      ? ((await districts.getDistrictById(ID)) as DistrictForm)
      : entity.value === "villages"
      ? ((await villages.getVillageById(ID)) as VillageForm)
      : ((await terms.getTermById(ID)) as TermForm);

  console.log(formModel.value);
});

type FormFieldsNames = keyof (DistrictForm | VillageForm | TermForm);
type FormModels = District | Village | Term;

const entityHandlers = {
  districts: (data: FormModels) =>
    updateDistrict(Number(route.params.id), data as District),
  villages: (data: FormModels) =>
    updateVillage(Number(route.params.id), data as Village),
  guidance: (data: FormModels) =>
    updateTerm(Number(route.params.id), data as Term),
};

const handleSubmit = () => {
  const handler = entityHandlers[entity.value];
  if (handler) {
    handler(formModel.value as FormModels);
    alert("Сущность изменена");

    updateEntity(entity.value);
    router.push(`/admin/${entity.value}`);
  } else {
    console.error("Unknown entity type:", entity);
  }
};
</script>

<template>
  <section class="add-entity py-[80px]">
    <h2>Изменить экзэмпляр сущности</h2>
    <form class="flex flex-col gap-[12px] w-[800px]">
      <div class="flex flex-col gap-[8px]" v-for="field in formFields">
        <div v-if="field.field === 'input'">
          <label :for="field.name">{{ field.label }}</label>
          <InputText
            v-if="formModel"
            v-model="formModel[field.name as FormFieldsNames]"
            class="w-full"
            :id="field.name"
          />
        </div>

        <div v-if="field.field === 'textarea'">
          <label :for="field.name">{{ field.label }}</label>
          <Textarea
            v-if="formModel"
            v-model="formModel[field.name as FormFieldsNames]"
            class="w-full !py-[0.75em]"
            :id="field.name"
            rows="5"
            autoResize
          />
        </div>

        <div v-if="field.field === 'markdown'">
          <label :for="field.name">{{ field.label }}</label>
          <Editor
            v-if="formModel"
            v-model="formModel[field.name as FormFieldsNames]"
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
