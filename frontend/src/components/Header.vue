<script setup lang="ts"></script>

<template>
  <header
    class="w-[100vw] flex items-center h-[80px] 2xl:px-[calc((100vw-1280px)/2)] lg:px-[80px] shadow-lg justify-between"
  >
    <RouterLink to="/">
      <div class="flex gap-[12px] text-[32px] items-center font-bold">
        <img
          src="/assets/icons/logo.svg"
          alt="Логотип"
          class="w-[48px] select-none"
        />
        <h3 class="text-2xl">Чукотка: цифровая энциклопедия</h3>
      </div>
    </RouterLink>
    <nav class="flex gap-[48px]">
      <RouterLink to="/">Главная</RouterLink>
      <RouterLink to="/districts">Районы</RouterLink>
      <RouterLink to="/guidance">Энциклопедия</RouterLink>
      <RouterLink to="/about">О нас</RouterLink>
    </nav>
  </header>
</template>

<style scoped></style>
