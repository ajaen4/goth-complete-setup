/** @type {import('tailwindcss').Config} */
export default {
  content: ["./internal/components/**/*.templ"],
  theme: {
    extend: {
      colors: require("tailwindcss/colors"),
    },
  },
  plugins: [],
};
