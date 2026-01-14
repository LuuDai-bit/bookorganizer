<template>
  <!-- Start card -->
  <div class="flex flex-col bg-white rounded-lg shadow-md w-full m-6 overflow-hidden sm:w-52">
    <svg class="w-6 h-6 text-gray-800" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 18 20">
        <path d="M16 0H4a2 2 0 0 0-2 2v1H1a1 1 0 0 0 0 2h1v2H1a1 1 0 0 0 0 2h1v2H1a1 1 0 0 0 0 2h1v2H1a1 1 0 0 0 0 2h1v1a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V2a2 2 0 0 0-2-2Zm-5.5 4.5a3 3 0 1 1 0 6 3 3 0 0 1 0-6ZM13.929 17H7.071a.5.5 0 0 1-.5-.5 3.935 3.935 0 1 1 7.858 0 .5.5 0 0 1-.5.5Z"/>
    </svg>
    <img :src="source" alt="" class="h-20 m-6">

    <h2 class="text-center px-2 pb-5 text-gray-800">{{ name }}</h2>

    <a href="#"
       class="bg-blue-500 text-white p-3 text-center hover:bg-blue-800 transition-all duration-500"
       @click="redirectTo(dest)">Go</a>
  </div>
  <!-- End card -->
</template>

<script>
import router from '@/router/index';

export default {
  name: "MenuCard",
  props: ["name", "source", "dest"],
  methods: {
    redirectTo (dest) {
      router.push({ path: dest })
    }
  }
}
</script>
