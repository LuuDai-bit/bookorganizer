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
                @openModal="openModal" />
    </div>
  </div>
  <BookForm v-if="isShow"
            :book="book"
            @closeModal="closeModal()"
            @fetchBooks="fetchBooks()" />
</template>

<script>
import Breadcrumb from '@/components/Breadcrumb.vue';
import BookCard from '@/components/BookCard.vue';
import FunctionButton from '@/components/FunctionButton.vue';
import BookForm from '@/components/BookForm.vue';
import bookApis from '@/api/books';
import session_apis from '@/api/sessions';

export default {
  name: "BookView",
  components: {
    Breadcrumb,
    BookCard,
    FunctionButton,
    BookForm,
  },
  data() {
    return {
      books: [],
      page: 1,
      isShow: false,
      book: null,
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
      })
    },
    openModal(book) {
      this.book = book
      this.isShow = true
    },
    closeModal() {
      this.isShow = false;
    },
    out() {
      session_apis.logout()
    }
  }
}
</script>
