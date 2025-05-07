<template>
  <div class="min-h-screen flex flex-col" style="background-color: #f6f6f6;">
    <Navbar />
    <main class="flex-1 flex flex-col items-center justify-center">
      <div class="w-full max-w-md rounded shadow p-8" style="background-color: #fff;">
        <h1 class="text-xl font-bold mb-6" style="color: #22223b;">Agregar nueva tarea</h1>
        <form @submit.prevent="guardarTarea" class="space-y-5">
          <div>
            <input
              v-model="titulo"
              type="text"
              required
              placeholder="Título de la tarea"
              class="w-full px-3 py-2 rounded border"
              style="background-color: #fff; color: #22223b; border-color: #e0e0e0;"
            />
          </div>
          <div>
            <textarea
              v-model="descripcion"
              rows="3"
              placeholder="Descripción (opcional)"
              class="w-full px-3 py-2 rounded border"
              style="background-color: #fff; color: #22223b; border-color: #e0e0e0;"
            ></textarea>
          </div>
          <div class="flex justify-end gap-2">
            <NuxtLink
              to="/"
              class="px-4 py-2 rounded font-semibold"
              style="background-color: #e0e0e0; color: #22223b;"
            >
              Cancelar
            </NuxtLink>
            <button
              type="submit"
              :disabled="isSubmitting || !titulo.trim()"
              class="px-4 py-2 rounded font-semibold shadow-sm"
              :style="`background-color: #fa541c; color: #fff; opacity: ${isSubmitting || !titulo.trim() ? 0.5 : 1}`"
            >
              {{ isSubmitting ? 'Guardando...' : 'Guardar' }}
            </button>
          </div>
          <div v-if="mensaje" class="text-green-600 mt-2">{{ mensaje }}</div>
        </form>
      </div>
    </main>
    <Footer />
  </div>
</template>

<script setup>
import Navbar from '~/components/Navbar.vue'
import Footer from '~/components/Footer.vue'
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const titulo = ref('')
const descripcion = ref('')
const isSubmitting = ref(false)
const mensaje = ref('')

const guardarTarea = async () => {
  if (!titulo.value.trim()) return
  isSubmitting.value = true
  const baseUrl = process.server
    ? 'http://backend:8080'
    : (process.env.NUXT_PUBLIC_API_URL || 'http://localhost:8080');
  await $fetch(`${baseUrl}/tasks`, {
    method: 'POST',
    body: { title: titulo.value, description: descripcion.value }
  })
  mensaje.value = 'Tarea guardada correctamente'
  titulo.value = ''
  descripcion.value = ''
  isSubmitting.value = false
  setTimeout(() => router.push('/'), 800)
}
</script>