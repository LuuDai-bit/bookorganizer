<template>
  <!-- Main modal -->
  <div id="default-modal"
        tabindex="-1"
        aria-hidden="true"
        class="backdrop-blur overflow-y-auto overflow-x-hidden fixed top-0 right-0 left-0 z-50 justify-center items-center w-full md:inset-0 h-[calc(100%-1rem)] max-h-full">
    <div class="mx-auto relative p-4 w-full max-w-2xl max-h-full">
      <!-- Modal content -->
      <div class="relative bg-white rounded-lg shadow dark:bg-gray-700">
        <!-- Modal header -->
        <div class="flex items-center justify-between p-4 md:p-5 border-b rounded-t dark:border-gray-600">
          <h3 class="text-xl font-semibold text-gray-900 dark:text-white">
            Reading a new book?
          </h3>
          <button type="button"
                  class="text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm w-8 h-8 ms-auto inline-flex justify-center items-center dark:hover:bg-gray-600 dark:hover:text-white"
                  @click="closeModal()">
            <svg class="w-3 h-3" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 14">
              <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6"/>
            </svg>
            <span class="sr-only">Close modal</span>
          </button>
        </div>

        <!-- Modal body -->
        <div class="p-4 md:p-5 space-y-4">
          <div class="mb-4">
            <label class="block text-gray-700 text-sm font-bold mb-2" for="name">
              Name
            </label>
            <input class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                  id="name"
                  type="text"
                  placeholder="Joe Doe"
                  v-model="name">
            <div class="input-errors" v-for="error of v$.name.$errors" :key="error.$uid">
              <div class="error-msg">{{ error.$message }}</div>
            </div>
          </div>

          <div class="mb-4">
            <label class="block text-gray-700 text-sm font-bold mb-2" for="author">
              Author
            </label>
            <input class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                  id="author"
                  type="text"
                  placeholder="Robert"
                  v-model="author">
            <div class="input-errors" v-for="error of v$.author.$errors" :key="error.$uid">
              <div class="error-msg">{{ error.$message }}</div>
            </div>
          </div>

          <div class="mb-4">
            <label class="block text-gray-700 text-sm font-bold mb-2" for="purchaseDate">
              Purchase Date
            </label>
            <input class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                  id="purchaseDate"
                  type="date"
                  v-model="purchaseDate">
            <div class="input-errors" v-for="error of v$.purchaseDate.$errors" :key="error.$uid">
              <div class="error-msg">{{ error.$message }}</div>
            </div>
          </div>

          <div class="mb-4">
            <label class="block text-gray-700 text-sm font-bold mb-2" for="startReadAt">
              Start Read At
            </label>
            <input class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                  id="startReadAt"
                  type="date"
                  v-model="startReadAt">
          </div>

          <div class="mb-4">
            <label class="block text-gray-700 text-sm font-bold mb-2" for="startReadAt">
              Finish Read At
            </label>
            <input class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                  id="finishReadAt"
                  type="date"
                  v-model="finishReadAt">
          </div>

          <div class="mb-4">
            <label class="block text-gray-700 text-sm font-bold mb-2" for="startReadAt">
              Categories
            </label>
            <div class="flex-wrap mb-1"
                 v-if="categories.length > 0">
              <span class="inline-block bg-gray-100 text-gray-800 text-sm font-medium me-2 px-2.5 py-0.5 rounded dark:bg-gray-700 dark:text-gray-300"
                    v-for="ct in categories">{{ ct }}</span>
            </div>
            <div class="flex shadow-sm">
              <input class="shadow appearance-none border rounded w-full py-2 px-3 block text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    id="categories"
                    maxlength="20"
                    type="text"
                    v-model="category">

              <button type="button"
                      class="w-[2.875rem] h-[1.25] flex-shrink-0 inline-flex justify-center items-center gap-x-2 text-sm font-semibold rounded-e-md border border-transparent text-white hover:bg-gray-100 disabled:opacity-50 disabled:pointer-events-none dark:focus:outline-none dark:focus:ring-1 dark:focus:ring-gray-600"
                      @click="addCategory(category)">
                <svg class="w-6 h-6 text-gray-800" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 18 18">
                  <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 1v16M1 9h16"/>
                </svg>
              </button>
            </div>
          </div>

        </div>

        <!-- Modal footer -->
        <div class="flex items-center p-4 md:p-5 border-t border-gray-200 rounded-b dark:border-gray-600">
          <SubmitButton text="Submit"
                    :disable=false
                    @submit="onSubmit(name, author, purchaseDate, startReadAt, finishReadAt, categories)"/>
          <button type="button"
                  class="ms-3 text-gray-500 bg-white hover:bg-gray-100 focus:ring-4 focus:outline-none focus:ring-blue-300 rounded-lg border border-gray-200 text-sm font-medium px-5 py-2.5 hover:text-gray-900 focus:z-10 dark:bg-gray-700 dark:text-gray-300 dark:border-gray-500 dark:hover:text-white dark:hover:bg-gray-600 dark:focus:ring-gray-600"
                  @click="closeModal()">
                  Cancel
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { useVuelidate } from '@vuelidate/core';
import { required } from "@vuelidate/validators";
import bookApis from '@/api/books';
import { notifyError, notifySuccess } from '../api/notification';
import SubmitButton from './SubmitButton.vue';

export default {
  name: "BookForm",
  setup() {
    return { v$: useVuelidate() }
  },
  components: {
    SubmitButton,
  },
  data() {
    return {
      name: '',
      author: '',
      purchaseDate: '',
      startReadAt: '',
      finishReadAt: '',
      categories: [],
      category: '',
    }
  },
  validations: {
    name: { required },
    author: { required },
    purchaseDate: { required },
  },
  methods: {
    closeModal() {
      this.$emit('closeModal')
    },
    onSubmit(name, author, purchaseDate, startReadAt, finishReadAt, categories) {
      this.v$.$validate();
      if (this.v$.name.$invalid || this.v$.author.$invalid || this.v$.purchaseDate.$invalid) return

      let self = this
      bookApis.createBook(name, author, purchaseDate, startReadAt, finishReadAt, categories)
        .then(function (response) {
          if(response.data.message) notifySuccess(response.data.message)
          self.$emit('closeModal')
          self.$emit('fetchBooks')
        }).catch(function (error) {
          if(error.response.data.message) notifyError(error.response.data.message)
        })
    },
    addCategory(category) {
      this.categories.push(category)
      this.category = ''
    }
  }
}
</script>
