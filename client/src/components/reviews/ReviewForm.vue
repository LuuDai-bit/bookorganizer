<template>
  <!-- Main modal -->
  <div id="default-modal"
        tabindex="-1"
        aria-hidden="true"
        class="backdrop-blur overflow-y-auto overflow-x-hidden fixed top-0 right-0 left-0 z-50 justify-center items-center w-full h-full md:inset-0 h-[calc(100%-1rem)] max-h-full">
    <div class="mx-auto relative p-4 w-full max-w-2xl max-h-full">
      <!-- Modal content -->
      <div class="relative bg-white rounded-lg shadow dark:bg-gray-700">
        <!-- Modal header -->
        <div class="flex items-center justify-between p-4 md:p-5 border-b rounded-t dark:border-gray-600">
          <h3 class="text-xl font-semibold text-gray-900 dark:text-white">
            Add Review for {{ book.name }}
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
          <Vue3StarRatings v-model="point" />
        </div>

        <div class="p-4 md:p-5 space-y-4">
          <textarea class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                id="review"
                type="text"
                placeholder="What do you think about this book?"
                v-model="review"></textarea>
        </div>
        <!-- Modal footer -->
        <div class="flex items-center p-4 md:p-5 border-t border-gray-200 rounded-b dark:border-gray-600">
          <SubmitButton text="Submit"
                        :disable=false
                        @submit="onSubmit(point, review)"/>
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
import Vue3StarRatings from "vue3-star-ratings";
import SubmitButton from "../common/SubmitButton.vue";
import reviewApis from "@/api/reviews";
import { notifyError, notifySuccess } from "@/api/notification";

export default {
  name: "ReviewForm",
  props: ["book"],
  components: {
    Vue3StarRatings,
    SubmitButton,
  },
  data() {
    return {
      point: 0,
      review: '',
    }
  },
  methods: {
    closeModal() {
      this.$emit('closeModal')
    },
    onSubmit(point, review) {
      let self = this
      reviewApis.createReview(this.book.id, point, review)
        .then(function(response) {
          self.closeModal()
          if(response.data.message) notifySuccess(response.data.message)
          self.$emit(`fetchBooks`)
        }).catch(function(error) {
          if(error.response.data.message) notifyError(error.response.data.message)
        })
    }
  }
}
</script>
