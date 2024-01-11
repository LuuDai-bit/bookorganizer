<script setup>
import { RouterLink, RouterView } from 'vue-router'
import Footer from './components/common/Footer.vue';
import Loading from './components/common/Loading.vue';
</script>

<template>
  <RouterView />
  <Footer />
  <notifications />
  <Loading />
</template>
