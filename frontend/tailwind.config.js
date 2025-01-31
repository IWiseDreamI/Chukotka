/** @type {import('tailwindcss').Config} */
export default {
  purge: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  content: [],
  theme: {
    extend: {
      boxShadow: {
        'reverse': '0 10px 15px 3px rgb(0 0 0 / 0.1), 0 4px 6px 4px rgb(0 0 0 / 0.1);',
      },
      screens: {
        'mc': '380px'
      },
    },
  },
  plugins: [],
}

