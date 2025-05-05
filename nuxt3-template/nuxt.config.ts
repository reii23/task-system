export default defineNuxtConfig({
  typescript: {
    strict: true
  },

  app: {
    head: {
      title: 'App Tareas',
      meta: [
        { name: 'viewport', content: 'width=device-width, initial-scale=1' },
        { name: 'description', content: 'Template de prueba para postulantes' }
      ]
    }
  },

  modules: ['@nuxtjs/tailwindcss']
})