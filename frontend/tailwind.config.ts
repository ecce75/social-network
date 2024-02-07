import type { Config } from "tailwindcss"

const config:Config = {
  darkMode: ["class"],
  content: [
    './src/pages/**/*.{ts,tsx}',
    './src/components/**/*.{ts,tsx}',
    './src/app/**/*.{ts,tsx}',
    './src/**/*.{ts,tsx}',
  ],
  prefix: "",
  theme: {
    container: {
      center: true,
      padding: "2rem",
      screens: {
        "2xl": "1400px",
      },
    },
    extend: {
      fontFamily: {
        rasa: ['Rasa'],
        sans: ['Rasa', 'sans-serif'],
      },
      
    },
  },
    daisyui: {
      themes: [
        {
          mytheme: {
            "primary": "#4d7c0f",
          },
        },
      ],
    },
  

    plugins: [require("daisyui")],
} 
export default config

