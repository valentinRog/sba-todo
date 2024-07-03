module.exports = (opts = {}) => {
  const { addClass } = opts;

  return {
    postcssPlugin: "postcss-add-class",
    Rule(rule) {
      rule.selectors = rule.selectors.map((selector) => {
        if (addClass === "" || selector.endsWith(`.${addClass}`)) {
          return selector;
        }
        return `${selector}.${addClass}`;
      });
    },
  };
};

module.exports.postcss = true;
