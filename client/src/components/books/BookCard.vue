<template>
  <div class="inline-block mr-5 mb-5 w-full max-w-sm p-4 bg-white border border-gray-200 rounded-lg shadow sm:p-8 dark:bg-gray-800 dark:border-gray-700">
    <h5 class="mb-4 text-xl font-medium text-gray-500 dark:text-gray-400 inline">{{ book.name }}</h5>
    <svg xmlns="http://www.w3.org/2000/svg"
         fill="none"
         viewBox="0 0 24 24"
         stroke-width="1.5"
         stroke="currentColor"
         class="w-6 h-6 inline float-right cursor-pointer"
         v-if="book.file_name"
         @click="downloadFile()">
      <path stroke-linecap="round" stroke-linejoin="round" d="m9 12.75 3 3m0 0 3-3m-3 3v-7.5M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z" />
    </svg>

    <ul role="list" class="space-y-5 my-7">
      <li class="flex items-center">
        <span class="text-base font-normal leading-tight text-gray-500 dark:text-gray-400 ms-3">Author: {{ book.author }}</span>
      </li>
      <li class="flex items-center">
        <span class="text-base font-normal leading-tight text-gray-500 dark:text-gray-400 ms-3">Purchase date: {{ formatDate(book.purchase_date) }}</span>
      </li>
      <li class="flex items-center">
        <span class="text-base font-normal leading-tight text-gray-500 dark:text-gray-400 ms-3">Start read at: {{ formatDate(book.start_read_at) || '~' }}</span>
      </li>
      <li class="flex items-center">
        <span class="text-base font-normal leading-tight text-gray-500 dark:text-gray-400 ms-3">Finish read at: {{ formatDate(book.finish_read_at) || '~' }}</span>
      </li>
      <li class="flex items-center">
        <span class="text-base font-normal leading-tight text-gray-500 dark:text-gray-400 ms-3">Type: {{ book.type || '~' }}</span>
      </li>
      <li class="flex items-center">
        <span class="text-base font-normal leading-tight text-gray-500 dark:text-gray-400 ms-3">Categories: {{ book.categories.join(', ') || '~' }}</span>
      </li>
    </ul>

    <div class="flex justify-evenly">
      <button type="button"
              class="text-white bg-blue-600 hover:bg-blue-700 focus:ring-4 focus:outline-none focus:ring-blue-200 dark:focus:ring-blue-900 font-medium rounded-lg text-sm px-5 py-2.5 inline-flex justify-center w-2/5 text-center"
              @click="openModal(book)">Update</button>
      <button type="button"
              class="text-white bg-blue-600 hover:bg-blue-700 focus:ring-4 focus:outline-none focus:ring-blue-200 dark:focus:ring-blue-900 font-medium rounded-lg text-sm px-5 py-2.5 inline-flex justify-center w-2/5 text-center"
              @click="openReviewModal(book)">Add Review</button>
    </div>
  </div>
</template>

<script>
import bookApis from '@/api/books';
import formatDate from '@/mixins/format_date';

export default {
  name: "BookCard",
  props: ["book"],
  methods: {
    formatDate,
    openModal(book) {
      this.$emit("openModal", book)
    },
    openReviewModal(book) {
      this.$emit("openReviewModal", book)
    },
    downloadFile() {
      const self = this;
      bookApis.download(this.book.id).then(function(response) {
        const link = document.createElement("a")
        link.href = response.data.url
        link.download = self.book.name.split(' ').join('_')
        link.click()
        URL.revokeObjectURL(link.href)
      })
    }
  }
}
</script>
