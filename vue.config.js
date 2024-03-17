const { defineConfig } = require('@vue/cli-service')

module.exports = defineConfig({
  transpileDependencies: true,
  chainWebpack: (config) => {
    config.plugin('define').tap((args) => {
      const defineArgs = args[0];
      defineArgs['__VUE_PROD_HYDRATION_MISMATCH_DETAILS__'] = true;
      return [defineArgs];
    });
  },
});