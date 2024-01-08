<template>
  <div id="books">
    <Breadcrumb paths="Books" />

    <div class="lg:flex justify-between items-center mb-6">
      <p class="text-2xl font-semibold mb-2 lg:mb-0">Welcome to book organizer!</p>
      <FunctionButton text="Logout" @click="out()"/>
    </div>

    <div class="lg:flex justify-between items-center mb-6">
      <FunctionButton text="Create" @click="openModal(null)"/>
    </div>

    <div v-if="books.length === 0">
      <p>No books, it's time to add some</p>
    </div>
    <div v-else>
      <BookCard v-for="book in books"
                :book="book"
                @openModal="openModal"
                @openReviewModal="openReviewModal" />
    </div>

    <Pagination :currentPage="page"
                :total="total"
                :perPage="perPage"
                @goToPage="goToPage" />
  </div>
  <BookForm v-if="isShow"
            :book="book"
            @closeModal="closeModal()"
            @fetchBooks="fetchBooks()" />
  <ReviewForm v-if="isReview"
              :book="book"
              @closeModal="closeModal('review')" />
</template>

<script>
import Breadcrumb from '@/components/common/Breadcrumb.vue';
import BookCard from '@/components/books/BookCard.vue';
import FunctionButton from '@/components/common/FunctionButton.vue';
import BookForm from '@/components/books/BookForm.vue';
import Pagination from '@/components/common/Pagination.vue';
import ReviewForm from '@/components/reviews/ReviewForm.vue';
import bookApis from '@/api/books';
import session_apis from '@/api/sessions';

export default {
  name: "BookView",
  components: {
    Breadcrumb,
    BookCard,
    FunctionButton,
    BookForm,
    Pagination,
    ReviewForm,
  },
  data() {
    return {
      books: [],
      page: 1,
      isShow: false,
      isReview: false,
      book: null,
      total: 0,
      perPage: 20,
    }
  },
  created() {
    this.fetchBooks()
  },
  methods: {
    fetchBooks() {
      let self = this
      bookApis.getBooks(this.page).then(function (response) {
        self.books = response.data.books || []
        self.total = response.data.total
      })
    },
    openModal(book) {
      this.book = book
      this.isShow = true
    },
    openReviewModal(book) {
      this.book = book
      this.isReview = true
    },
    closeModal(type) {
      if(type == 'review') {
        this.isReview = false;
      } else {
        this.isShow = false;
      }
    },
    out() {
      session_apis.logout()
    },
    goToPage(page) {
      this.page = page
      this.fetchBooks()
    }
  }
}
</script>
