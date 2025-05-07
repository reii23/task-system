<template>
    <div class="w-full max-w-2xl rounded shadow p-8 bg-white">
      <!-- pendientes -->
      <div v-if="tareasPendientes.length > 0" class="mb-8">
        <h2 class="text-lg font-semibold mb-3" style="color: #22223b;">Pendientes</h2>
        <ul>
          <li v-for="tarea in tareasPendientes" :key="tarea.id" class="flex items-center justify-between border-b py-3" style="border-color: #f0f0f0;">
            <div>
              <span class="font-medium" style="color: #22223b;">{{ tarea.title }}</span>
              <span v-if="tarea.description" class="ml-2 text-sm" style="color: #6c757d;">{{ tarea.description }}</span>
            </div>
            <button
              @click="toggleCompletada(tarea.id)"
              class="ml-4 px-4 py-1 rounded font-semibold shadow-sm"
              style="background-color: #fa541c; color: #fff;"
            >
              Completar
            </button>
          </li>
        </ul>
      </div>
  
      <!-- completadas -->
      <div v-if="tareasCompletadas.length > 0">
        <h2 class="text-lg font-semibold mb-3" style="color: #6c757d;">Completadas</h2>
        <ul>
          <li v-for="tarea in tareasCompletadas" :key="tarea.id" class="flex items-center justify-between border-b py-3" style="border-color: #f0f0f0;">
            <div>
              <span class="line-through" style="color: #bdbdbd;">{{ tarea.title }}</span>
              <span v-if="tarea.description" class="ml-2 text-sm line-through" style="color: #bdbdbd;">{{ tarea.description }}</span>
            </div>
            <span class="font-bold" style="color: #fa541c;">✓</span>
          </li>
        </ul>
      </div>
  
      <!-- mensaje cuando no hay tareas -->
      <div v-if="tareas.length === 0" class="text-center py-10">
        <h2 class="text-lg font-semibold mb-2" style="color: #22223b;">No hay tareas</h2>
        <NuxtLink
          to="/crear-tarea"
          class="px-4 py-2 rounded font-semibold shadow-sm"
          style="background-color: #fa541c; color: #fff;"
        >
          Crear tarea
        </NuxtLink>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref, computed, onMounted } from 'vue'
  
  const tareas = ref([])
  
  async function cargarTareas() {
    const baseUrl = process.server
      ? 'http://backend:8080'
      : (process.env.NUXT_PUBLIC_API_URL || 'http://localhost:8080');
    tareas.value = await $fetch(`${baseUrl}/tasks`)
  }
  
  onMounted(cargarTareas)
  
  const tareasPendientes = computed(() => (tareas.value || []).filter(t => !t.completed))
  const tareasCompletadas = computed(() => (tareas.value || []).filter(t => t.completed))
  
  async function toggleCompletada(id) {
    const baseUrl = process.server
      ? 'http://backend:8080'
      : (process.env.NUXT_PUBLIC_API_URL || 'http://localhost:8080');
    await $fetch(`${baseUrl}/tasks/${id}/complete`, { method: 'PUT' })
    await cargarTareas()
  }
  </script>