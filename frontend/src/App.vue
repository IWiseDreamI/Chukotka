<script setup lang="ts">
import Header from "./components/Header.vue";
import Footer from "./components/Footer.vue";
</script>

<template>
  <Header />
  <main
    class="min-h-[calc(100vh-80px)] relative overflow-x-hidden lg:overflow-x-auto"
  >
    <Suspense>
      <RouterView />
    </Suspense>
  </main>
  <Footer />
</template>

<style scoped></style>
