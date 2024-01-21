const withMT = require("@material-tailwind/html/utils/withMT");

module.exports = withMT({
  content: ["./views/**/*.templ", "./views/**/*.html", "./views/**/*.go"],
  theme: {
    extend: {},
  },
  plugins: [],
});