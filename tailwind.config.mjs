/** @type {import('tailwindcss').Config} */
export default {
  darkMode: 'class',
  content: [
    './src/**/*.{astro,html,js,jsx,ts,tsx,svelte}',
  ],
  theme: {
    extend: {
      colors: {
        'obsidian': '#09090B',
        'alabaster': '#FAFAFA',
        'violet': '#7C3AED',
        'mint': '#10B981',
      },
    },
  },
  plugins: [],
}