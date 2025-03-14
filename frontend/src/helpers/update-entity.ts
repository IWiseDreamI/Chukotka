import { useDistrictStore } from "@/stores/districts";
import { useTermStore } from "@/stores/terms";
import { useVillageStore } from "@/stores/villages";
import { useMaterialStore } from "@/stores/materials";
import { useAboutStore } from "@/stores/about";
import { entityTypes } from "@/types";

export const updateEntity = (entity: entityTypes) => {
  console.log('[updateEntity] Обновление сущности:', entity);
  const districts = useDistrictStore();
  const villages = useVillageStore();
  const terms = useTermStore();
  const materials = useMaterialStore();
  const about = useAboutStore();

  if (entity === "districts") {
    districts.fetchDistricts();
    console.log('[updateEntity] districts.fetchDistricts() вызван');
  } else if (entity === "villages") {
    villages.fetchVillages();
    console.log('[updateEntity] villages.fetchVillages() вызван');
  } else if (entity === "guidance") {
    terms.fetchTerms();
    console.log('[updateEntity] terms.fetchTerms() вызван');
  } else if (entity === "materials") {
    materials.fetchMaterials();
    console.log('[updateEntity] materials.fetchMaterials() вызван');
  } else if (entity === "about") {
    about.fetchAbout();
    console.log('[updateEntity] about.fetchAboutPage() вызван');
  }
};

