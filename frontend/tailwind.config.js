/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      colors: {
        primary: '#a855f7', // Purple
        accent: '#f472b6',  // Pink
        dark: '#0a0a0a',
        darker: '#050505',
      },
      fontFamily: {
        heading: ['Orbitron', 'monospace'],
        body: ['Roboto', 'sans-serif'],
      },
    },
  },
  plugins: [],
}
