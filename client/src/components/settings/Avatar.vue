<template>
  <div class="inline-block mr-5 w-[30%] items-center">
    <img :src="url" alt="Avatar..." class="rounded w-36 h-36 m-auto">
  </div>
</template>

<script>
  export default {
    name: "Avatar",
    props: ["url"]
  }
</script>
