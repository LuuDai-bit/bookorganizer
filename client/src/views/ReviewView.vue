<template>
  <div id="reviews">
    <Breadcrumb paths="Reviews" />

    <div class="lg:flex justify-between items-center mb-6">
      <p class="text-2xl font-semibold mb-2 lg:mb-0">Welcome to book reviews!</p>
      <FunctionButton text="Logout" @clickHandler="out()"/>
    </div>

    <div v-for="book in books">
      <div class="lg:flex justify-between items-center my-10">
        <h3 class="text-3xl font-bold dark:text-white inline">
          <span>
            <svg class="w-6 h-6 text-gray-800 dark:text-white inline" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 16 20">
              <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M1 17V2a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1H3a2 2 0 0 0-2 2Zm0 0a2 2 0 0 0 2 2h12M5 15V1m8 18v-4"/>
            </svg>
          </span>
          {{ book.name }}
        </h3>
        <FunctionButton text="Add Review"
                        @clickHandler="openReviewModal(book)"/>
      </div>

      <div v-for="review in book.reviews">
        <div class="flex flex-col gap-3 mt-5">
          <ReviewItem :book="book" :review="review" />
        </div>
      </div>
      <div v-if="!book.reviews">
        No reviews, please add one
      </div>
    </div>

    <Pagination :currentPage="page"
                :total="total"
                :perPage="perPage"
                @goToPage="goToPage" />
  </div>
  <ReviewForm v-if="isReview"
              :book="book"
              @closeModal="closeModal()"
              @fetchBooks="fetchBooks()" />
</template>

<script>
import bookApis from '@/api/books';
import Breadcrumb from '@/components/common/Breadcrumb.vue';
import ReviewItem from '@/components/reviews/ReviewItem.vue';
import FunctionButton from '@/components/common/FunctionButton.vue';
import ReviewForm from '@/components/reviews/ReviewForm.vue';
import Pagination from '@/components/common/Pagination.vue';

export default {
  name: "Review",
  components: {
    Breadcrumb,
    ReviewItem,
    FunctionButton,
    ReviewForm,
    Pagination,
  },
  data() {
    return {
      books: [],
      total: 0,
      page: 1,
      perPage: 20,
      book: null,
      isReview: false,
    };
  },
  created() {
    this.fetchBooks();
  },
  methods: {
    fetchBooks() {
      let self = this;
      bookApis.getBooks(this.page).then(function (response) {
        self.books = response.data.books || [];
        self.total = response.data.total;
      });
    },
    openReviewModal(book) {
      this.book = book
      this.isReview = true
    },
    closeModal() {
      this.isReview = false
    },
    goToPage(page) {
      this.page = page
      this.fetchBooks()
    },
  },
  components: { Breadcrumb, ReviewItem, FunctionButton, ReviewForm, Pagination }
}
</script>
