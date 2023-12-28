<template>
  <form class="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4">
    <div class="mb-4">
      <label class="block text-gray-700 text-sm font-bold mb-2" for="email">
        Email
      </label>
      <input class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
             id="email"
             type="email"
             placeholder="Email"
             v-model="email"
             @focusout="validate()">
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
             @focusout="validate()">
    </div>
    <div class="flex items-center justify-between">
      <button class="text-white font-bold py-2 px-4 rounded focus:outline-none"
              :class="{ 'bg-blue-500 hover:bg-blue-700': !invalid(),
                        'bg-gray-300 hover:bg-gray-300': invalid() }"
              type="button"
              @click="onSubmit(email, password)"
              :disabled="invalid()">
        Sign In
      </button>
      <a class="inline-block align-baseline font-bold text-sm text-blue-500 hover:text-blue-800" href="#">
        Forgot Password?
      </a>
    </div>
  </form>
</template>

<script>
import { useVuelidate } from '@vuelidate/core'
import { required, email } from "@vuelidate/validators";
import session_apis from "../api/sessions";

export default {
  name: "LoginInput",
  setup() {
    return { v$: useVuelidate() }
  },
  data() {
    return {
      email: "",
      password: null
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
      // call api here
      session_apis.login(email, password);
    },
    validate () {
      this.v$.$validate()
    },
    invalid () {
      return (this.v$.email.$invalid || this.v$.password.$invalid)
    }
  }
}
</script>
