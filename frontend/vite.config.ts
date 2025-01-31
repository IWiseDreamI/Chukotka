import * as path from 'path';
import {defineConfig} from 'vite';
import vue from '@vitejs/plugin-vue';

// https://vitejs.dev/config/
export default defineConfig({
  server: {
    port: 4200
  },
  resolve: {
    alias: [
      { find: '@', replacement: path.resolve(__dirname, './src') },
    ],
  },
  plugins: [vue()],
  build: {
    target: 'es2022',
    sourcemap: true
  },
});
