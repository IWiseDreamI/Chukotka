<script setup lang="ts">
import { Card } from "primevue";

const entities = [
  { link: "/admin/districts", text: "Районы", icon: "pi-map-marker" },
  { link: "/admin/villages", text: "Поселения", icon: "pi-building" },
  { link: "/admin/guidance", text: "Энциклопедия", icon: "pi-book" },
  { link: "/admin/materials", text: "Материалы", icon: "pi-file" },
  { link: "/admin/about", text: "О нас", icon: "pi-info-circle" },
];
</script>

<template>
  <section class="admin py-[80px]">
    <RouterLink to="/admin">
      <h1>Администраторская панель</h1>
    </RouterLink>

    <div class="flex flex-col items-center w-full gap-[32px] md:flex-row">
      <RouterLink
        v-for="(item, index) in entities"
        :key="index"
        :to="item.link"
      >
        <Card class="card">
          <template #content>
            <div class="flex flex-col gap-[24px] w-[300px]">
              <i :class="'pi ' + item.icon" />
              <h3>{{ item.text }}</h3>
            </div>
          </template>
        </Card>
      </RouterLink>
    </div>
  </section>
</template>

<style scoped lang="scss">
.admin {
  gap: 32px;
  display: flex;
  align-items: start;
  flex-direction: column;
  min-height: calc(100vh - 80px);
}

.card {
  cursor: pointer;
  padding: 12px;
  box-shadow: 0 0 8px #0000002f;

  .pi {
    font-size: 80px;
  }
}
</style>
