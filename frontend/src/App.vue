<script setup lang="ts">
import Header from "./components/Header.vue";
import Footer from "./components/Footer.vue";
</script>

<template>
  <Header />
  <main class="w-full min-h-[calc(100vh-80px)] overflow-x-hidden">
    <RouterView />
  </main>
  <Footer />
</template>

<style scoped></style>
