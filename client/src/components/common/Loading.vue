<template>
  <div id="loading"
        tabindex="-1"
        aria-hidden="true"
        class="backdrop-blur overflow-y-auto overflow-x-hidden fixed top-0 right-0 left-0 z-50 justify-center items-center w-full h-full md:inset-0 h-[calc(100%-1rem)] max-h-full flex hidden">
    <div class="m-auto relative p-4 w-full max-w-2xl max-h-2xl items-center">
      <img :src="source" alt="Loading..." class="m-auto">
    </div>
  </div>
</template>

<script>
export default {
  name: "Loading",
  data() {
    return {
      source: "book-loading.gif"
    }
  }
}
</script>
