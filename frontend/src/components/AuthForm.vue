<script lang="ts" setup>
import { ref } from "vue";
import { loginAdmin } from "@/entities/api/auth";
import { Card, InputText, Password, Button, FloatLabel } from "primevue";
import { useRouter } from "vue-router";

const router = useRouter();

const loginData = ref({
  username: "",
  password: "",
});

const handleLogin = async () => {
  try {
    const response = await loginAdmin(loginData.value);
    alert("Login successful!");
    localStorage.setItem("token", response.token);
    router.push("/admin");
    window.location.reload();
  } catch (error) {
    alert("Login failed. Please check your credentials.");
  }
};
</script>

<template>
  <Card class="login-container">
    <template #title>
      <h2>Вход</h2>
    </template>

    <template #content>
      <form @submit.prevent="handleLogin" class="login-form">
        <FloatLabel class="field">
          <label class="z-10" for="login-username">Username</label>
          <InputText
            id="login-username"
            v-model="loginData.username"
            required
          />
        </FloatLabel>
        <FloatLabel class="field">
          <label class="z-10" for="login-password">Password</label>
          <Password
            id="login-password"
            v-model="loginData.password"
            toggleMask
            required
          />
        </FloatLabel>
        <Button label="Войти" type="submit" class="p-button-primary" />
      </form>
    </template>
  </Card>
</template>

<style scoped lang="scss">
.login-container {
  width: 400px;
  margin-bottom: 40px;
  box-shadow: 0 0 12px #0000005b;
}

.login-form {
  display: flex;
  flex-direction: column;
  gap: 32px;

  .p-inputtext,
  .p-password {
    width: 100% !important;
  }
}
</style>
