<template>
  <main>
    <LoginInput></LoginInput>
  </main>
</template>

<script>
import LoginInput from '@/components/login/LoginInput.vue';

export default {
  name: "LoginView",
  components: {
    LoginInput,
}
}
</script>

<style scoped>
main {
  margin: auto;
}
</style>
