<script setup>
import { RouterLink, RouterView } from 'vue-router'
import Footer from './components/common/Footer.vue';
</script>

<template>
  <RouterView />
  <Footer />
  <notifications />
</template>
