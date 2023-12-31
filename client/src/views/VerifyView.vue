<template>
<div class="max-w-md mx-auto border max-w-sm mt-20 rounded">
  <form class="shadow-md px-4 py-6">
    <span>Please check email for verify code!</span>
    <br/><br/>
    <div class="flex justify-center gap-2 mb-6">
      <div class="mb-4">
        <input class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              id="verifyCode"
              maxlength="6"
              placeholder="000000"
              v-model="verifyCode"
              @input="validate()">
      </div>
    </div>
    <div class="flex items-center justify-center">
      <SubmitButton text="Submit"
                    :disable="disable"
                    @submit="onSubmit(verifyCode)"/>
      <a class="inline-block align-baseline font-bold text-sm text-teal-500 hover:text-teal-800 ml-4" href="#">
          Resend OTP
      </a>
    </div>
  </form>
</div>
</template>

<script>
import { useVuelidate } from '@vuelidate/core';
import { required, helpers } from "@vuelidate/validators";
import session_apis from "../api/sessions";
import router from "../router/index";
import SubmitButton from "../components/SubmitButton.vue";

const mustContainOnlyNumber = helpers.withParams(
  { type: "verifyCode" },
  (value) => !helpers.req(value) || /^\d+$/.test(value)
)

export default {
  name: "VerifyView",
  setup() {
    return { v$: useVuelidate() }
  },
  created () {
    session_apis.sendCode(this.$route.query.email)
  },
  mounted () {
    if(!this.$route.query.email) router.push({ path: '/login' })
  },
  components: {
    SubmitButton,
  },
  data () {
    return {
      verifyCode: null,
      disable: true,
    }
  },
  validations: {
    verifyCode: {
      required,
      mustContainOnlyNumber,
    }
  },
  methods: {
    validate () {
      this.v$.$validate()
      this.disable = this.v$.verifyCode.$invalid
    },
    onSubmit (verifyCode) {
      session_apis.verify(this.$route.query.email, verifyCode)
    },
  },
}
</script>
