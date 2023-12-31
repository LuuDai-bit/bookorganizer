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

    <div class="mb-4">
      <label class="block text-gray-700 text-sm font-bold mb-2" for="email">
        Name
      </label>
      <input class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
             id="name"
             type="name"
             placeholder="Name"
             v-model="name"
             @focusout="validate()">
      <div class="input-errors" v-for="error of v$.name.$errors" :key="error.$uid">
        <div class="error-msg">{{ error.$message }}</div>
      </div>
    </div>

    <div class="mb-6">
      <label class="block text-gray-700 text-sm font-bold mb-2" for="password">
        Password
      </label>
      <input class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
             id="password"
             type="password"
             placeholder="******************"
             v-model="password"
             @focusout="validate()">
      <div class="input-errors" v-for="error of v$.password.$errors" :key="error.$uid">
        <div class="error-msg">{{ error.$message }}</div>
      </div>
    </div>

    <div class="mb-6">
      <label class="block text-gray-700 text-sm font-bold mb-2" for="password_confirmation">
        Password Confirmation
      </label>
      <input class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline"
             id="password_confirmation"
             type="password"
             placeholder="******************"
             v-model="passwordConfirmation"
             @focusout="validate()">
      <div class="input-errors" v-for="error of v$.passwordConfirmation.$errors" :key="error.$uid">
        <div class="error-msg">{{ error.$message }}</div>
      </div>
    </div>
    <div class="flex items-center justify-between">
      <button class="text-white font-bold py-2 px-4 rounded focus:outline-none"
              :class="{ 'bg-blue-500 hover:bg-blue-700': !invalid(),
                        'bg-gray-300 hover:bg-gray-300': invalid() }"
              type="button"
              @click="onSubmit(email, name, password)"
              :disabled="invalid()">
        Signup
      </button>
      <a class="inline-block align-baseline font-bold text-sm text-blue-500 hover:text-blue-800"
         @click="redirectToSignin()">
        Signin
      </a>
    </div>
  </form>
</template>

<script>
import { useVuelidate } from '@vuelidate/core';
import { required, email, helpers } from "@vuelidate/validators";
import session_apis from "../api/sessions";
import router from "../router/index";

const passwordMustMatch = helpers.withParams(
  { type: 'passwordConfirmation' },
  (value, siblings) => !helpers.req(value) || value == siblings.password
)

export default {
  name: "Signup",
  setup() {
    return { v$: useVuelidate() }
  },
  data() {
    return {
      email: null,
      name: null,
      password: null,
      passwordConfirmation: null,
    }
  },
  validations: {
    email: {
      required,
      email
    },
    name: { required},
    password: { required },
    passwordConfirmation: {
      required,
      passwordMustMatch: helpers.withMessage('Password not match', passwordMustMatch),
    }
  },
  methods: {
    onSubmit (email, name, password) {
      session_apis.signup(email, name, password);
    },
    validate () {
      this.v$.$validate()
    },
    invalid () {
      return (
        this.v$.email.$invalid ||
        this.v$.name.$invalid ||
        this.v$.password.$invalid ||
        this.v$.passwordConfirmation.$invalid
      )
    },
    redirectToSignin () {
      router.push({ path: '/login' })
    }
  }
}
</script>
