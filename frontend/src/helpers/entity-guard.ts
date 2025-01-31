import { useRouter } from "vue-router";

export const entityGuard = (entity: string) => {
  const router = useRouter();

  if (!["districts", "villages", "guidance"].includes(entity)) {
    router.push("/admin");
  }
}