<script setup lang="ts">
import { defineProps } from "vue";

defineProps<{
  content: string;
  class?: string;
}>();
</script>

<template>
  <div :class="class" class="w-full text-left py-[12px]" v-html="content" />
</template>
