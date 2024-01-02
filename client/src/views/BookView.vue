<template>
  <div id="books">
    <Breadcrumb paths="Books" />

    <div class="lg:flex justify-between items-center mb-6">
      <p class="text-2xl font-semibold mb-2 lg:mb-0">Welcome to book organizer!</p>
      <FunctionButton text="Logout" @click="logout()"/>
    </div>

    <div class="lg:flex justify-between items-center mb-6">
      <FunctionButton text="Create" @click="openModal()"/>
    </div>

    <div v-if="books.length === 0">
      <p>No books, it's time to add some</p>
    </div>
    <div v-else>
      <BookCard v-for="book in books" :book="book"/>
    </div>
  </div>
  <BookForm v-if="isShow"
            @closeModal="closeModal()"
            @fetchBooks="fetchBooks()" />
</template>

<script>
import Breadcrumb from '@/components/Breadcrumb.vue';
import BookCard from '@/components/BookCard.vue';
import FunctionButton from '@/components/FunctionButton.vue';
import BookForm from '@/components/BookForm.vue';
import bookApis from '../api/books';

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
    openModal() {
      this.isShow = true
    },
    closeModal() {
      this.isShow = false;
    }
  }
}
</script>
