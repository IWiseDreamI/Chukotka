<script setup lang="ts">
import { computed, onMounted } from "vue";
import { api } from "@/entities/api/axios";

const logged = computed(() => localStorage.getItem("token"));

const logout = () => {
  if (logged.value) {
    localStorage.removeItem("token");
  }
};

onMounted(() => {
  if (logged.value) {
    api
      .get("/is-admin", {
        headers: {
          Authorization: logged.value,
        },
      })
      .then((response) => {
        if (!response.data.valid) {
          localStorage.removeItem("token");
        }
      })
      .catch(() => {
        localStorage.removeItem("token");
      });
  }
});
</script>

<template>
  <footer
    class="w-[100vw] flex flex-col py-[20px] shadow-reverse lg:px-[80px] 2xl:px-[calc((100vw-1280px)/2)] box-border"
  >
    <div
      class="flex flex-col items-end w-full gap-[12px] sm:flex-row sm:justify-between sm:gap-0"
    >
      <div class="hidden md:flex gap-[36px]">
        <RouterLink :to="logged ? '/admin' : '/auth'" class="text-base">
          Панель
        </RouterLink>
        <p v-if="logged" @click="logout" class="text-base cursor-pointer">
          Выйти
        </p>
      </div>

      <p class="cursor-pointer">2024 © WiseDream</p>
    </div>
  </footer>
</template>

<style scoped></style>
