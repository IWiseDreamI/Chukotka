<script setup lang="ts">
import { ref } from "vue";
import { Button, Drawer } from "primevue";
const menu = ref(false);
</script>

<template>
  <header
    class="w-[100vw] flex items-center h-[80px] 2xl:px-[calc((100vw-1280px)/2)] sm:px-[80px] px-[12px] shadow-lg justify-between"
  >
    <RouterLink to="/">
      <div class="flex gap-[12px] items-center font-bold">
        <img
          src="/assets/icons/logo.svg"
          alt="Логотип"
          class="w-[32px] sm:w-[48px] select-none"
        />
        <h3 class="text-xl sm:text-2xl">
          Чукотка <b class="hidden sm:visible">: цифровая энциклопедия</b>
        </h3>
      </div>
    </RouterLink>
    <nav class="hidden gap-[48px] lg:flex">
      <RouterLink to="/">Главная</RouterLink>
      <RouterLink to="/villages">Сёла</RouterLink>
      <RouterLink to="/materials">Материалы</RouterLink>
      <RouterLink to="/guidance">Энциклопедия</RouterLink>
      <RouterLink to="/about">О проекте</RouterLink>
    </nav>
    <nav class="flex gap-[48px] lg:hidden">
      <Drawer v-model:visible="menu" header="Навигация" position="right">
        <div class="w-full flex flex-col gap-[12px]">
          <RouterLink to="/">Главная</RouterLink>
          <RouterLink to="/villages">Сёла</RouterLink>
          <RouterLink to="/materials">Материалы</RouterLink>
          <RouterLink to="/guidance">Энциклопедия</RouterLink>
          <RouterLink to="/about">О проекте</RouterLink>
        </div>
      </Drawer>
      <Button icon="pi pi-bars" @click="menu = true" />
    </nav>
  </header>
</template>

<style scoped></style>
