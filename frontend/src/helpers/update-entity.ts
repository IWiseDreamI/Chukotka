import { useDistrictStore } from "@/stores/districts";
import { useTermStore } from "@/stores/terms";
import { useVillageStore } from "@/stores/villages";
import { entityTypes } from "@/types"

export const updateEntity = (entity: entityTypes ) => {
  const districts = useDistrictStore();
  const villages = useVillageStore();
  const terms = useTermStore()

  if (entity === "districts") districts.fetchDistricts();
  if (entity === "villages") villages.fetchVillages();
  if (entity === "guidance") terms.fetchTerms();
}