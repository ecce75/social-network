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
            "primary": "#4F7942",
            "secondary": "#355E3B",
            "accent": "#a21caf",
            "neutral": "#ffffff",
            "base-100": "#e5e7eb",
            "info": "#ffffff",
            "success": "#355E3B",
            "warning": "#ffffff",
            "error": "#ffffff",
          },
        },
      ],
    },
  

    plugins: [require("daisyui")],
} 
export default config

