<script setup lang="ts">
import { ref } from "vue";

const logged = ref(localStorage.getItem("token"));

const logout = () => {
  if (logged.value) {
    localStorage.removeItem("token");
    window.location.reload();
  }
};
</script>

<template>
  <footer
    class="w-[100vw] flex flex-col py-[20px] shadow-reverse lg:px-[80px] 2xl:px-[calc((100vw-1280px)/2)] box-border"
  >
    <div
      class="flex flex-col items-end w-full gap-[12px] sm:flex-row sm:justify-between sm:gap-0"
    >
      <div class="flex gap-[36px]">
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
