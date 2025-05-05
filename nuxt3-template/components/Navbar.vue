<template>
  <nav class="w-full shadow mb-8" style="background-color: #000009;">
    <div class="max-w-2xl mx-auto flex justify-between items-center py-4 px-6">
      <h1 class="text-xl font-bold" style="color: #fa541c;">Gestion de tareas</h1>
      <NuxtLink
        to="/crear-tarea"
        class="px-4 py-2 rounded font-semibold shadow-sm transition-colors"
        style="background-color: #fa541c; color: #fff;"
      >
        + Nueva tarea
      </NuxtLink>
    </div>
  </nav>
</template>