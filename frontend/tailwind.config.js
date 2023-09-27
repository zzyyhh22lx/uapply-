/** @type {import('tailwindcss').Config} */
const { token } = require('./src/libs/token/index.js');

function buildApplyColors(name) {
  const levels = [10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 'Light', 'Hover', 'Active', 'Disabled'];
  const re = {};
  levels.forEach(level => {
    if (token[`applyColor${name}${level}`]) {
      re[`apply-${name.toLocaleLowerCase()}-${level}`] = token[`applyColor${name}${level}`];
    }
  });
  if (token[`applyColor${name}`]) {
    re[`apply-${name.toLocaleLowerCase()}`] = token[`applyColor${name}`];
  }
  return re;
}

module.exports = {
  content: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
  darkMode: 'media',
  theme: {
    extend: {
      colors: {
        ...buildApplyColors('Info'),
        ...buildApplyColors('Black'),
        ...buildApplyColors('Gray'),
        ...buildApplyColors('Error'),
        ...buildApplyColors('Success'),
        ...buildApplyColors('Warning'),
        ...buildApplyColors('Brand'),
      }
    },
  },
  plugins: [require("@tailwindcss/typography")],
};
