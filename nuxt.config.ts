// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2025-05-15',
  devtools: { enabled: true },
  modules: ['@nuxt/ui', '@nuxt/icon', '@nuxt/fonts'],
  runtimeConfig: {
    kubernetes: {
      labelSelector: process.env.LABEL_SELECTOR ?? 'ok8.sh/dashboard=true',
    }
  }
})