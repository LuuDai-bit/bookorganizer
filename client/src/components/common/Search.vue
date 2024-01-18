<template>
  <label for="default-search" class="mb-2 text-sm font-medium text-gray-900 sr-only dark:text-white">Search</label>
  <div class="relative float-right">
    <div class="absolute inset-y-0 start-0 flex items-center ps-3 pointer-events-none">
      <svg class="w-4 h-4 text-gray-500 dark:text-gray-400" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 20 20">
          <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19 19-4-4m0-7A7 7 0 1 1 1 8a7 7 0 0 1 14 0Z"/>
      </svg>
    </div>
    <input type="search"
           class="block w-full p-2 ps-10 text-sm text-gray-900 border border-gray-300 rounded-lg bg-gray-50 focus:outline-none"
           placeholder="Book name"
           v-model="search">
    <button class="text-white absolute inset-y-0 end-0 w-16 flex items-center ps-3 bg-blue-500 hover:bg-blue-600 focus:outline-none rounded-e-lg text-white font-semibold shadow focus:ring-blue-300 font-medium rounded-e-lg text-sm dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
            @click="clickHandler()">Search</button>
  </div>
</template>

<script>
export default {
  name: "Search",
  data() {
    return {
      search: '',
    }
  },
  methods: {
    clickHandler() {
      this.$emit(`clickHandler`, this.search)
    }
  }
}
</script>
