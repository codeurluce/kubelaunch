import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      fontFamily: {
        sans: ["Syne", "sans-serif"],
        mono: ["DM Mono", "monospace"],
      },
      colors: {
        bg: "#0a0e17",
        surface: "#111827",
        surface2: "#1a2234",
        border: "#1e2d47",
        accent: "#3b82f6",
        accent2: "#06b6d4",
        "k-green": "#10b981",
        "k-amber": "#f59e0b",
        "k-red": "#ef4444",
        "k-muted": "#64748b",
        "k-text": "#e2e8f0",
      },
      keyframes: {
        pulse2: {
          "0%, 100%": { opacity: "1" },
          "50%": { opacity: "0.3" },
        },
        fadeIn: {
          from: { opacity: "0", transform: "translateY(8px)" },
          to: { opacity: "1", transform: "translateY(0)" },
        },
        slideIn: {
          from: { opacity: "0", transform: "translateX(-8px)" },
          to: { opacity: "1", transform: "translateX(0)" },
        },
      },
      animation: {
        pulse2: "pulse2 1.5s infinite",
        fadeIn: "fadeIn 0.3s ease forwards",
        slideIn: "slideIn 0.2s ease forwards",
      },
    },
  },
  plugins: [],
};

export default config;
