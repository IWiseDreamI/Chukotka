<script setup lang="ts">
import Markdown from "@/components/Markdown.vue";
import { useAboutStore } from "@/stores/about";
import { onMounted, ref } from "vue";

const about = useAboutStore();
const data = ref<string>("");

onMounted(async () => {
  data.value = (await about.getAboutPage()).content;
});
</script>

<template>
  <div class="about">
    <Markdown :content="data" />
  </div>
</template>

<style scoped>
.about {
  margin-top: 40px;
}
</style>
