<template>
  <div class="flex items-center justify-center">
    <nav aria-label="Page navigation example">
      <ul class="inline-flex -space-x-px text-base h-10">
        <li>
          <a class="flex items-center justify-center px-4 h-10 leading-tight border border-gray-300 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400"
              :class="(currentPage == 1) ? `text-gray-600 bg-gray-50 hover:bg-gray-50 hover:text-gray-600`:`text-gray-500 bg-white hover:bg-gray-100 hover:text-gray-700 dark:hover:bg-gray-700 dark:hover:text-white`"
              :disabled="currentPage == 1"
              @click="goPreviousPage()">Previous
            </a>
        </li>
        <li v-for="page of listPage">
          <a class="flex items-center justify-center px-4 h-10 leading-tight border border-gray-300 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400"
              :class="(page == currentPage) ? `text-blue-600 bg-blue-50 hover:bg-blue-50 hover:text-blue-600`:`text-gray-500 bg-white hover:bg-gray-100 hover:text-gray-700 dark:hover:bg-gray-700 dark:hover:text-white`"
              :disabled="page == currentPage"
              @click="goChoosenPage(page)">{{ page }}</a>
        </li>
        <li>
          <a class="flex items-center justify-center px-4 h-10 leading-tight border border-gray-300 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400"
              :class="(currentPage == numberOfPage) ? `text-gray-600 bg-gray-50 hover:bg-gray-50 hover:text-gray-600`:`text-gray-500 bg-white hover:bg-gray-100 hover:text-gray-700 dark:hover:bg-gray-700 dark:hover:text-white`"
              :disabled="currentPage == numberOfPage"
             @click="goNextPage()">Next</a>
        </li>
      </ul>
    </nav>
  </div>
</template>

<script>
export default {
  name: "Pagination",
  props: ["currentPage", "total", "perPage"],
  data() {
    return {
      numberOfPage: 0,
      listPage: [],
      disablePrevious: false,
      disableNext: false,
    }
  },
  beforeUpdate() {
    this.numberOfPage = Math.ceil(this.total / this.perPage)
    this.listPage = Array.from({length: this.numberOfPage}, (_, i) => i + 1)
    this.disablePrevious = (this.currentPage == 1)
    this.disableNext = (this.currentPage == 1)
  },
  methods: {
    goPreviousPage() {
      if(this.currentPage == 1) return

      this.scrollToTop()
      this.$emit("goToPage", this.currentPage - 1)
    },
    goNextPage() {
      if(this.currentPage == this.numberOfPage) return

      this.scrollToTop()
      this.$emit("goToPage", this.currentPage + 1)
    },
    goChoosenPage(page) {
      if(this.currentPage == page) return

      this.scrollToTop()
      this.$emit("goToPage", page)
    },
    scrollToTop() {
      window.scrollTo(0, 0)
    },
  }
}
</script>
