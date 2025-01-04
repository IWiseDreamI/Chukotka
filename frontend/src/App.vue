<script setup lang="ts">
import Header from "./components/Header.vue";
import Footer from "./components/Footer.vue";
</script>

<template>
  <Header />
  <main class="w-full 2xl:px-[calc((100vw-1280px)/2)] min-h-[calc(100vh-80px)]">
    <RouterView />
  </main>
  <Footer />
</template>

<style scoped></style>
