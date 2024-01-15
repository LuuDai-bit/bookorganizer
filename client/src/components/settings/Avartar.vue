<template>
  <div class="inline-block mr-5 w-[30%]">
    <img :src="url" alt="Avatar..." width="30%" height="30%">
  </div>
</template>

<script>
  export default {
    name: "Avatar",
    props: ["url"]
  }
</script>
