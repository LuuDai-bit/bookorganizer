<template>
  <form class="bg-white shadow-md rounded px-8 pt-6 pb-8 my-auto w-80">
    <div class="mb-4">
      <label class="block text-gray-700 text-sm font-bold mb-2" for="email">
        Email
      </label>
      <input class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
             id="email"
             type="email"
             placeholder="Email"
             v-model="email"
             @input="validate()">
      <div class="input-errors" v-for="error of v$.email.$errors" :key="error.$uid">
        <div class="error-msg">{{ error.$message }}</div>
      </div>
    </div>
    <div class="mb-6">
      <label class="block text-gray-700 text-sm font-bold mb-2" for="password">
        Password
      </label>
      <input class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline"
             id="password"
             type="password"
             placeholder="******************"
             v-model="password"
             @input="validate()">
    </div>
    <div class="flex items-center justify-between">
      <SubmitButton text="Sign In"
                    :disable="disable"
                    @submit="onSubmit(email, password)"/>
      <a class="inline-block align-baseline font-bold text-sm text-blue-500 hover:text-blue-800"
          @click="redirectToSignup()">
        Signup
      </a>
    </div>
  </form>
</template>

<script>
import { useVuelidate } from '@vuelidate/core';
import { required, email } from "@vuelidate/validators";
import session_apis from "../api/sessions";
import router from "../router/index";
import SubmitButton from "../components/SubmitButton.vue"

export default {
  name: "LoginInput",
  setup() {
    return { v$: useVuelidate() }
  },
  components: {
    SubmitButton,
  },
  data() {
    return {
      email: "",
      password: null,
      disable: true,
    }
  },
  validations: {
    email: {
      required,
      email
    },
    password: { required },
  },
  methods: {
    onSubmit (email, password) {
      session_apis.login(email, password);
    },
    validate () {
      this.v$.$validate();
      this.disable = this.v$.email.$invalid || this.v$.password.$invalid
    },
    redirectToSignup () {
      router.push({ path: '/signup' });
    },
  }
}
</script>
