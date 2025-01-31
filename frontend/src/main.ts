import { createApp } from 'vue'
import { createPinia } from 'pinia'

import router from './index'

import App from './App.vue'

import './style.scss'
import 'primeicons/primeicons.css';
import PrimeVue from 'primevue/config';
import Lara from '@primevue/themes/aura';
import { definePreset } from '@primevue/themes';

const preset = definePreset(Lara, {
  semantic: {
      primary: {
          50: '#79553D',
          100: '#79553D',
          200: '#79553D',
          300: '#79553D',
          400: '#79553D',
          500: '#3D2B1F',
          600: '#3D2B1F',
          700: '#3D2B1F',
          800: '#3D2B1F',
          900: '#3D2B1F',
          950: '#3D2B1F'
      }
  }
})

const app = createApp(App);
const pinia = createPinia()

app.use(PrimeVue, {
  theme: {
    preset: preset,
    options: {
      darkModeSelector: false || 'none',
    }
  }
});
app.use(router);
app.use(pinia);

app.mount('#app')
