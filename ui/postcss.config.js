module.exports = {
  plugins: [
    require("postcss-import"),
    require("./postcss-add-class")({ addClass: process.env.PREFIX }),
    require("postcss-nested"),
    require("postcss-minify"),
  ],
};
