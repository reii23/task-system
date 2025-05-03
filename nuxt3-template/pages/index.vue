<template>
  <div class="min-h-screen bg-gray-100 text-gray-900">
    <nav class="bg-blue-600 text-white px-6 py-4 flex justify-between items-center">
      <span class="font-bold text-lg">Mini Sistema de Tareas</span>
    </nav>
    <main class="p-6">
      <div class="max-w-lg mx-auto bg-white p-8 rounded shadow mb-8">
        <h2 class="text-2xl font-bold mb-6">Crear nueva tarea</h2>
        <form @submit.prevent="crearTarea" class="space-y-4">
          <div>
            <label class="block mb-1 font-semibold">Título</label>
            <input v-model="title" type="text" placeholder="Título" required class="border rounded w-full p-2" />
          </div>
          <div>
            <label class="block mb-1 font-semibold">Descripción</label>
            <textarea v-model="description" placeholder="Descripción" class="border rounded w-full p-2"></textarea>
          </div>
          <button type="submit" class="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700">Crear</button>
        </form>
        <div v-if="mensaje" class="mt-4 text-green-600">{{ mensaje }}</div>
      </div>
      <div class="max-w-2xl mx-auto bg-white p-8 rounded shadow">
        <h2 class="text-2xl font-bold mb-6">Lista de tareas</h2>
        <div v-if="tareas.length === 0" class="text-gray-500">No hay tareas registradas.</div>
        <ul>
          <li
            v-for="tarea in tareas"
            :key="tarea.id"
            class="flex items-center justify-between border-b py-3"
          >
            <div>
              <span :class="{ 'line-through text-gray-400': tarea.completed }">
                <strong>{{ tarea.title }}</strong>
                <span class="ml-2 text-sm text-gray-500">{{ tarea.description }}</span>
              </span>
            </div>
            <div>
              <button
                v-if="!tarea.completed"
                @click="completarTarea(tarea.id)"
                class="bg-green-500 text-white px-3 py-1 rounded hover:bg-green-600"
              >
                Completar
              </button>
              <span v-else class="text-green-600 font-bold">Completada</span>
            </div>
          </li>
        </ul>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const title = ref('')
const description = ref('')
const mensaje = ref('')
const tareas = ref([])

async function cargarTareas() {
  tareas.value = await $fetch('http://localhost:8080/tasks')
}

async function crearTarea() {
  if (!title.value.trim()) {
    mensaje.value = 'El título es obligatorio'
    return
  }
  await $fetch('http://localhost:8080/tasks', {
    method: 'POST',
    body: { title: title.value, description: description.value }
  })
  mensaje.value = 'Tarea creada correctamente'
  title.value = ''
  description.value = ''
  await cargarTareas() // Recarga la lista después de crear
}

async function completarTarea(id) {
  await $fetch(`http://localhost:8080/tasks/${id}/complete`, { method: 'PUT' })
  await cargarTareas() // Recarga la lista después de completar
}

onMounted(cargarTareas)
</script>