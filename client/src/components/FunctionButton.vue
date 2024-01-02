<template>
  <button class="bg-blue-500 hover:bg-blue-600 focus:outline-none rounded-lg px-6 py-2 text-white font-semibold shadow"
          @click="clickHandler()">
    {{ text }}</button>
</template>

<script>
  export default {
    name: "SubmitButton",
    props: ["text"],
    methods: {
      clickHandler () {
        this.$emit('click')
      }
    }
  }
</script>
