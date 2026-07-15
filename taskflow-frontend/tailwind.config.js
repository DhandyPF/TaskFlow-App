/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{js,ts,jsx,tsx}"],
  theme: {
    extend: {
      colors: {
        paper: "#F5F7FA",
        ink: "#16302B",
        "ink-soft": "#3F5750",
        line: "#DCE3E0",
        marigold: "#E8A33D",
        coral: "#D64545",
        flow: "#3D5AFE",
      },
      fontFamily: {
        display: [
          "Iowan Old Style",
          "Palatino Linotype",
          "Georgia",
          "serif",
        ],
        body: [
          "-apple-system",
          "BlinkMacSystemFont",
          "Segoe UI",
          "Helvetica Neue",
          "Arial",
          "sans-serif",
        ],
        mono: [
          "ui-monospace",
          "SF Mono",
          "Cascadia Code",
          "Roboto Mono",
          "monospace",
        ],
      },
    },
  },
  plugins: [],
};
