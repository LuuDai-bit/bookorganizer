<template>
  <Breadcrumb paths="Reviews" />
  <div id="reviews">
    <div v-for="book in books">
      <div v-for="review in book.reviews">
        <div class="flex flex-col gap-3 mt-14">
          <ReviewItem :book="book" :review="review" />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import bookApis from '@/api/books';
import Breadcrumb from '@/components/common/Breadcrumb.vue';
import ReviewItem from '@/components/reviews/ReviewItem.vue';

export default {
  name: "Review",
  components: {
    Breadcrumb,
    ReviewItem,
  },
  data() {
    return {
      books: [],
      total: 0,
      page: 1,
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
        self.books = self.books.filter((book) => book.reviews)
        self.total = response.data.total;
      });
    },
  },
  components: { Breadcrumb, ReviewItem }
}
</script>
