<template>
  <!-- Main modal -->
  <div id="default-modal"
        tabindex="-1"
        aria-hidden="true"
        class="backdrop-blur overflow-y-auto overflow-x-hidden fixed top-0 right-0 left-0 z-10 justify-center items-center w-full h-full md:inset-0 h-[calc(100%-1rem)] max-h-full">
    <div class="mx-auto relative p-4 w-full max-w-2xl max-h-full">
      <!-- Modal content -->
      <div class="relative bg-white rounded-lg shadow dark:bg-gray-700">
        <!-- Modal header -->
        <div class="flex items-center justify-between p-4 md:p-5 border-b rounded-t dark:border-gray-600">
          <h3 class="text-xl font-semibold text-gray-900 dark:text-white">
            {{ title }}
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
                  maxlength="50"
                  v-model="currentBook.name">
            <div class="input-errors" v-for="error of v$.currentBook.name.$errors" :key="error.$uid">
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
                  maxlength="30"
                  v-model="currentBook.author">
            <div class="input-errors" v-for="error of v$.currentBook.author.$errors" :key="error.$uid">
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
                  v-model="currentBook.purchaseDate">
            <div class="input-errors" v-for="error of v$.currentBook.purchaseDate.$errors" :key="error.$uid">
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
                  v-model="currentBook.startReadAt">
          </div>

          <div class="mb-4">
            <label class="block text-gray-700 text-sm font-bold mb-2" for="startReadAt">
              Finish Read At
            </label>
            <input class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                  id="finishReadAt"
                  type="date"
                  v-model="currentBook.finishReadAt">
          </div>

          <div>
            <label class="block text-gray-700 text-sm font-bold mb-2" for="startReadAt">
              Type
            </label>
            <select class="block py-2.5 px-0 w-full text-sm text-gray-500 bg-transparent border-0 border-b-2 border-gray-200 appearance-none dark:text-gray-400 dark:border-gray-700 focus:outline-none focus:ring-0 focus:border-gray-200 peer"
                    v-model="currentBook.type">
              <option value="paperback">Paperback</option>
              <option value="ebook">Ebook</option>
            </select>
          </div>

          <div class="mb-3" v-if="currentBook.type==='ebook'">
            <label
              for="formFile"
              class="mb-2 inline-block text-neutral-700 dark:text-neutral-200"
              > {{ ebookTitle }}
            </label>
            <svg xmlns="http://www.w3.org/2000/svg"
                 fill="none"
                 viewBox="0 0 24 24"
                 stroke-width="1.5"
                 stroke="currentColor"
                 class="w-6 h-6 inline"
                 v-if="this.currentBook.key"
                 @mouseover="changeTooltip(true)"
                 @mouseout="changeTooltip(false)"
                 @click="resetFile()">
              <path stroke-linecap="round" stroke-linejoin="round" d="M16.023 9.348h4.992v-.001M2.985 19.644v-4.992m0 0h4.992m-4.993 0 3.181 3.183a8.25 8.25 0 0 0 13.803-3.7M4.031 9.865a8.25 8.25 0 0 1 13.803-3.7l3.181 3.182m0-4.991v4.99" />
            </svg>
            <div id="tooltip-reload-ebook"
                 class="absolute z-50 inline-block px-3 py-2 text-sm font-medium text-white bg-gray-900 rounded-lg shadow-sm opacity-1 dark:bg-gray-700"
                 v-if="tooltip">
              Reset
            </div>

            <input
              class="relative m-0 block w-full min-w-0 flex-auto rounded border border-solid border-neutral-300 bg-clip-padding px-3 py-[0.32rem] text-base font-normal text-neutral-700 transition duration-300 ease-in-out file:-mx-3 file:-my-[0.32rem] file:overflow-hidden file:rounded-none file:border-0 file:border-solid file:border-inherit file:bg-neutral-100 file:px-3 file:py-[0.32rem] file:text-neutral-700 file:transition file:duration-150 file:ease-in-out file:[border-inline-end-width:1px] file:[margin-inline-end:0.75rem] hover:file:bg-neutral-200 focus:border-primary focus:text-neutral-700 focus:shadow-te-primary focus:outline-none dark:border-neutral-600 dark:text-neutral-200 dark:file:bg-neutral-700 dark:file:text-neutral-100 dark:focus:border-primary"
              type="file"
              accept=".epub, .mobi, .azw3, .html, .htm"
              id="formFile"
              ref="fileInput"
              @change="uploadEbook" />
          </div>

          <div class="mb-4">
            <label class="block text-gray-700 text-sm font-bold mb-2" for="startReadAt">
              Categories
            </label>
            <div class="flex-wrap mb-1"
                 v-if="currentBook.categories.length > 0">
              <span class="inline-block bg-gray-100 text-gray-800 text-sm font-medium me-2 px-2.5 py-0.5 rounded dark:bg-gray-700 dark:text-gray-300"
                    v-for="(ct, index) in currentBook.categories">
                    {{ ct }}
                    <button type="button"
                            class="text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm w-8 h-4 ms-auto inline-flex justify-center items-center dark:hover:bg-gray-600 dark:hover:text-white"
                            @click="removeCategory(index)">
                      <svg class="w-3 h-3" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 14">
                        <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6"/>
                      </svg>
                      <span class="sr-only">Remove category</span>
                    </button>
              </span>
            </div>
            <div class="flex shadow-sm">
              <input class="shadow appearance-none border rounded w-full py-2 px-3 block text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    id="categories"
                    maxlength="20"
                    type="text"
                    v-model="category"
                    v-on:keyup.enter="addCategory(category)">

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
                    @submit="checkWarning()"/>
          <button type="button"
                  class="ms-3 text-gray-500 bg-white hover:bg-gray-100 focus:ring-4 focus:outline-none focus:ring-blue-300 rounded-lg border border-gray-200 text-sm font-medium px-5 py-2.5 hover:text-gray-900 focus:z-10 dark:bg-gray-700 dark:text-gray-300 dark:border-gray-500 dark:hover:text-white dark:hover:bg-gray-600 dark:focus:ring-gray-600"
                  @click="closeModal()">
                  Cancel
          </button>
        </div>
      </div>
    </div>
  </div>
  <WarningModal v-if="warning"
                message="The previous file will be destroy"
                @closeWarningModal="closeWarningModal()"
                @confirm="onSubmit(currentBook)" />
</template>

<script>
import { useVuelidate } from '@vuelidate/core';
import { required } from "@vuelidate/validators";
import { notifyError, notifySuccess } from '@/api/notification';
import SubmitButton from '@/components/common/SubmitButton.vue';
import { formatDateWithFormat } from '@/mixins/format_date';
import bookApis from '@/api/books';
import fileApis from '@/api/file';
import WarningModal from '../common/WarningModal.vue';

export default {
  name: "BookForm",
  setup() {
    return { v$: useVuelidate() }
  },
  props: ["book"],
  components: {
    SubmitButton,
    WarningModal
  },
  data() {
    return {
      currentBook: {
        id: this.book ? this.book.id : '',
        name: this.book ? this.book.name : '',
        author: this.book ? this.book.author : '',
        purchaseDate: this.book ? formatDateWithFormat(this.book.purchase_date, 'YYYY-MM-DD') : '',
        startReadAt: this.book ? formatDateWithFormat(this.book.start_read_at, 'YYYY-MM-DD') : '',
        finishReadAt: this.book ? formatDateWithFormat(this.book.finish_read_at, 'YYYY-MM-DD') : '',
        key: '',
        fileName: this.book ? this.book.file_name : '',
        type: this.book ? this.book.type : 'paperback',
        categories: this.book ? this.book.categories : [],
      },
      category: '',
      nameKey: {},
      warning: false,
      tooltip: false,
    }
  },
  computed: {
    title() {
      let t = this.book ? "Update book?" : "Reading a new book?"
      return t
    },
    ebookTitle() {
      if(this.currentBook.fileName) {
        return this.currentBook.fileName
      } else if(this.book && this.book.file_name) {
        return this.book.file_name
      } else {
        return "Wanna store your ebook it here?"
      }
    }
  },
  validations: {
    currentBook: {
      name: { required },
      author: { required },
      purchaseDate: { required },
    },
  },
  methods: {
    closeModal() {
      this.$emit('closeModal')
    },
    onSubmit(currentBook) {
      this.v$.$validate();
      if (this.v$.currentBook.name.$invalid ||
          this.v$.currentBook.author.$invalid ||
          this.v$.currentBook.purchaseDate.$invalid) return

      let self = this
      if(this.book) {
        bookApis.updateBook(currentBook)
          .then(function (response) {
            if(response.data.message) notifySuccess(response.data.message)
            self.$emit('closeModal')
            self.$emit('fetchBooks')
          }).catch(function(error) {
            if(error.response.data.message) notifyError(error.response.data.message)
          })
      } else {
        bookApis.createBook(currentBook)
          .then(function (response) {
            if(response.data.message) notifySuccess(response.data.message)
            self.$emit('closeModal')
            self.$emit('fetchBooks')
          }).catch(function (error) {
            if(error.response.data.message) notifyError(error.response.data.message)
          })
      }
    },
    addCategory(category) {
      if(!category || this.currentBook.categories.includes(category)) {
        notifyError('Empty or duplicate category')
        return
      }

      this.currentBook.categories.push(category)
      this.category = ''
    },
    removeCategory(index) {
      this.currentBook.categories.splice(index, 1)
    },
    uploadEbook() {
      this.currentBook.fileName = this.$refs.fileInput.value.split(/(\\|\/)/g).pop()
      if(this.nameKey[this.$refs.fileInput.files[0]]) {
        this.currentBook.key = this.nameKey[this.$refs.fileInput.files[0]]

        return
      }

      let self = this
      fileApis.uploadSingleFile(this.$refs.fileInput.files[0]).then(function(response) {
        self.currentBook.key = response.data.key
        self.nameKey[self.$refs.fileInput.value] = self.currentBook.key
        self.disable = false
      }).catch(function(error) {
        if(error.response.data.message) notifyError(error.response.data.message)
      })
    },
    checkWarning() {
      if(this.currentBook.key && !this.book) {
        this.warning = true
        return
      }

      this.onSubmit(this.currentBook)
    },
    closeWarningModal() {
      this.warning = false
    },
    changeTooltip(value) {
      this.tooltip = value
    },
    resetFile() {
      this.currentBook.key = ""
      this.currentBook.fileName = this.book ? this.book.file_name : ""
      this.$refs.fileInput.value = null

      this.tooltip = false
    }
  }
}
</script>
