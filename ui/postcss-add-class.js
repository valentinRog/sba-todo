module.exports = (opts = {}) => {
  const { addClass } = opts;

  return {
    postcssPlugin: "postcss-add-class",
    Rule(rule) {
      rule.selectors = rule.selectors.map((selector) => {
        return `${selector}.${addClass}`;
      });
    },
  };
};

module.exports.postcss = true;
