<template>
  <button class="text-white font-bold py-2 px-4 rounded focus:outline-none"
          :class="{ 'bg-blue-500 hover:bg-blue-700': !disable,
                    'bg-gray-300 hover:bg-gray-300': disable }"
          type="button"
          @click="onSubmit()"
          :disabled="disable">
    {{ text }}
  </button>
</template>

<script>
  export default {
    name: "SubmitButton",
    props: ["text", "disable"],
    methods: {
      onSubmit () {
        this.$emit('submit')
      }
    }
  }
</script>
