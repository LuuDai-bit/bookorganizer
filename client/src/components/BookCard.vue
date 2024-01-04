<template>
  <div class="inline-block mr-5 mb-5 w-full max-w-sm p-4 bg-white border border-gray-200 rounded-lg shadow sm:p-8 dark:bg-gray-800 dark:border-gray-700">
    <h5 class="mb-4 text-xl font-medium text-gray-500 dark:text-gray-400">{{ book.name }}</h5>

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
        <span class="text-base font-normal leading-tight text-gray-500 dark:text-gray-400 ms-3">Categories: {{ book.categories.join(', ') || '~' }}</span>
      </li>
    </ul>
    <button type="button"
            class="text-white bg-blue-600 hover:bg-blue-700 focus:ring-4 focus:outline-none focus:ring-blue-200 dark:focus:ring-blue-900 font-medium rounded-lg text-sm px-5 py-2.5 inline-flex justify-center w-full text-center"
            @click="openModal(book)">Update</button>
  </div>
</template>

<script>
import formatDate from '../mixins/format_date';

export default {
  name: "BookCard",
  props: ["book"],
  methods: {
    formatDate,
    openModal(book) {
      this.$emit("openModal", book)
    }
  }
}
</script>
