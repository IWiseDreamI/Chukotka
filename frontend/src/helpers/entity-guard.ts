import { useRouter } from "vue-router";

export const entityGuard = (entity: string) => {
  const router = useRouter();

  if (!["districts", "villages", "guidance", "materials", "about"].includes(entity)) {
    router.push("/admin");
  }
}
