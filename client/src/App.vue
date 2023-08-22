<script setup>
import HelloWorld from './components/HelloWorld.vue'
import TheWelcome from './components/TheWelcome.vue'

import { useRouter } from 'vue-router';

const router = useRouter();
</script>

<template>
  <header>
    <div>
      header
    </div>
  </header>

  <main>
    <main class="main"><router-view /></main>
  </main>
</template>

<style scoped>
</style>
