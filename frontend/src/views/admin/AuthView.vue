<script setup lang="ts">
import AuthForm from "@/components/AuthForm.vue";
</script>

<template>
  <div class="auth">
    <AuthForm />
  </div>
</template>

<style scoped lang="scss">
.auth {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: calc(100vh - 80px);
}
</style>
