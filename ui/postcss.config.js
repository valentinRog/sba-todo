const path = require("path");

module.exports = {
  plugins: [
    require("postcss-import")({
      path: [path.resolve(__dirname, "style")],
    }),
    require("postcss-mixins"),
    require("postcss-nested"),
    require("postcss-simple-vars"),
    require("./postcss-add-class")({ addClass: process.env.PREFIX || "" }),
    require("postcss-minify"),
  ],
};
