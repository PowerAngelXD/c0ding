// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2025-07-15',
  devtools: { enabled: true },
  modules: ['@nuxt/ui'],
  fonts: {
    providers: {
      google: false,
      bunny: false,
      googleicons: false,
      'google-icons': false
    },
    families: [
      { name: 'Inter', provider: 'none' },
      { name: 'Material Symbols', provider: 'none' },
      { name: 'Material Symbols Outlined', provider: 'none' },
      { name: 'Material Symbols Rounded', provider: 'none' },
      { name: 'Material Symbols Sharp', provider: 'none' }
    ]
  }
})
