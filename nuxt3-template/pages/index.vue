<template>
  <div class="min-h-screen flex flex-col" style="background-color: #f6f6f6;">
    <Navbar />
    <main class="flex-1 flex flex-col items-center">
      <TaskList />
    </main>
    <Footer />
  </div>
</template>

<script setup>
import Navbar from '~/components/Navbar.vue'
import TaskList from '~/components/TaskList.vue'
import Footer from '~/components/Footer.vue'
</script>