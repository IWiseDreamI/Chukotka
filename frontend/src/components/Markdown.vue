<script setup lang="ts">
import { defineProps } from "vue";

defineProps<{
  content: string;
  class?: string;
}>();
</script>

<template>
  <div :class="class" class="w-full min-h-[400px] text-left" v-html="content" />
</template>
