<template>
  <div id="home">

    <Breadcrumb paths="Dashboard" />

    <div class="lg:flex justify-between items-center mb-6">
      <p class="text-2xl font-semibold mb-2 lg:mb-0">Hi, {{ name }}!</p>
      <button class="bg-blue-500 hover:bg-blue-600 focus:outline-none rounded-lg px-6 py-2 text-white font-semibold shadow"
              @click="logout()">
        Logout</button>
    </div>

    <div class="flex flex-wrap -mx-3 mb-20">

      <div class="w-1/2 xl:w-1/4 px-3">
        <div class="w-full bg-white border text-blue-400 rounded-lg flex items-center p-6 mb-6 xl:mb-0">
          <svg class="w-16 h-16 fill-current mr-4 hidden lg:block" viewBox="0 0 20 20">
            <path d="M17.35,2.219h-5.934c-0.115,0-0.225,0.045-0.307,0.128l-8.762,8.762c-0.171,0.168-0.171,0.443,0,0.611l5.933,5.934c0.167,0.171,0.443,0.169,0.612,0l8.762-8.763c0.083-0.083,0.128-0.192,0.128-0.307V2.651C17.781,2.414,17.587,2.219,17.35,2.219M16.916,8.405l-8.332,8.332l-5.321-5.321l8.333-8.332h5.32V8.405z M13.891,4.367c-0.957,0-1.729,0.772-1.729,1.729c0,0.957,0.771,1.729,1.729,1.729s1.729-0.772,1.729-1.729C15.619,5.14,14.848,4.367,13.891,4.367 M14.502,6.708c-0.326,0.326-0.896,0.326-1.223,0c-0.338-0.342-0.338-0.882,0-1.224c0.342-0.337,0.881-0.337,1.223,0C14.84,5.826,14.84,6.366,14.502,6.708"></path>
          </svg>

          <div class="text-gray-700">
            <p class="font-semibold text-3xl">{{ numberOfBook }}</p>
            <p>Books Read</p>
          </div>
        </div>
      </div>

      <div class="w-1/2 xl:w-1/4 px-3">
        <div class="w-full bg-white border text-blue-400 rounded-lg flex items-center p-6 mb-6 xl:mb-0">
          <svg class="w-16 h-16 fill-current mr-4 hidden lg:block" viewBox="0 0 20 20">
            <path d="M17.684,7.925l-5.131-0.67L10.329,2.57c-0.131-0.275-0.527-0.275-0.658,0L7.447,7.255l-5.131,0.67C2.014,7.964,1.892,8.333,2.113,8.54l3.76,3.568L4.924,17.21c-0.056,0.297,0.261,0.525,0.533,0.379L10,15.109l4.543,2.479c0.273,0.153,0.587-0.089,0.533-0.379l-0.949-5.103l3.76-3.568C18.108,8.333,17.986,7.964,17.684,7.925 M13.481,11.723c-0.089,0.083-0.129,0.205-0.105,0.324l0.848,4.547l-4.047-2.208c-0.055-0.03-0.116-0.045-0.176-0.045s-0.122,0.015-0.176,0.045l-4.047,2.208l0.847-4.547c0.023-0.119-0.016-0.241-0.105-0.324L3.162,8.54L7.74,7.941c0.124-0.016,0.229-0.093,0.282-0.203L10,3.568l1.978,4.17c0.053,0.11,0.158,0.187,0.282,0.203l4.578,0.598L13.481,11.723z"></path>
          </svg>

          <div class="text-gray-700">
            <p class="font-semibold text-3xl">{{ numberOfReview }}</p>
            <p>Reviews</p>
          </div>
        </div>
      </div>

      <div class="w-1/2 xl:w-1/4 px-3">
        <div class="w-full bg-white border text-blue-400 rounded-lg flex items-center p-6">
          <svg class="w-16 h-16 fill-current mr-4 hidden lg:block" viewBox="0 0 20 20">
            <path d="M14.999,8.543c0,0.229-0.188,0.417-0.416,0.417H5.417C5.187,8.959,5,8.772,5,8.543s0.188-0.417,0.417-0.417h9.167C14.812,8.126,14.999,8.314,14.999,8.543 M12.037,10.213H5.417C5.187,10.213,5,10.4,5,10.63c0,0.229,0.188,0.416,0.417,0.416h6.621c0.229,0,0.416-0.188,0.416-0.416C12.453,10.4,12.266,10.213,12.037,10.213 M14.583,6.046H5.417C5.187,6.046,5,6.233,5,6.463c0,0.229,0.188,0.417,0.417,0.417h9.167c0.229,0,0.416-0.188,0.416-0.417C14.999,6.233,14.812,6.046,14.583,6.046 M17.916,3.542v10c0,0.229-0.188,0.417-0.417,0.417H9.373l-2.829,2.796c-0.117,0.116-0.71,0.297-0.71-0.296v-2.5H2.5c-0.229,0-0.417-0.188-0.417-0.417v-10c0-0.229,0.188-0.417,0.417-0.417h15C17.729,3.126,17.916,3.313,17.916,3.542 M17.083,3.959H2.917v9.167H6.25c0.229,0,0.417,0.187,0.417,0.416v1.919l2.242-2.215c0.079-0.077,0.184-0.12,0.294-0.12h7.881V3.959z"></path>
          </svg>

          <div class="text-gray-700">
            <p class="font-semibold text-3xl">{{ numberOfNewBook }}</p>
            <p>New Book</p>
          </div>
        </div>
      </div>

      <div class="w-1/2 xl:w-1/4 px-3">
        <div class="w-full bg-white border text-blue-400 rounded-lg flex items-center p-6">
          <svg class="w-16 h-16 fill-current mr-4 hidden lg:block" viewBox="0 0 20 20">
            <path d="M17.431,2.156h-3.715c-0.228,0-0.413,0.186-0.413,0.413v6.973h-2.89V6.687c0-0.229-0.186-0.413-0.413-0.413H6.285c-0.228,0-0.413,0.184-0.413,0.413v6.388H2.569c-0.227,0-0.413,0.187-0.413,0.413v3.942c0,0.228,0.186,0.413,0.413,0.413h14.862c0.228,0,0.413-0.186,0.413-0.413V2.569C17.844,2.342,17.658,2.156,17.431,2.156 M5.872,17.019h-2.89v-3.117h2.89V17.019zM9.587,17.019h-2.89V7.1h2.89V17.019z M13.303,17.019h-2.89v-6.651h2.89V17.019z M17.019,17.019h-2.891V2.982h2.891V17.019z"></path>
          </svg>

          <div class="text-gray-700">
            <p class="font-semibold text-3xl">{{ "" }}</p>
            <p>Guess what bitch</p>
          </div>

        </div>
      </div>

    </div>

    <div class="flex flex-wrap -mx-3">

      <div class="w-full xl:w-1/3 px-3">
        <MenuCard name="Settings" source="settings.png" dest="/settings" />
      </div>

      <div class="w-full xl:w-1/3 px-3">
        <MenuCard name="Books" source="book.png" dest="/books" />
      </div>

      <div class="w-full xl:w-1/3 px-3">
        <MenuCard name="Reviews" source="question-mark.png" dest="/reviews" />
      </div>

    </div>


  </div>
</template>

<script>
import session_apis from "../api/sessions";
import MenuCard from "@/components/home/MenuCard.vue";
import Breadcrumb from '@/components/common/Breadcrumb.vue';
import Footer from "@/components/common/Footer.vue";
import statisticApis from "@/api/statistic";

export default {
  name: 'DashboardHome',
  components: {
    MenuCard,
    Breadcrumb,
    Footer,
  },
  created() {
    this.fetchNumberOfNewBook()
    this.fetchFavoriteCategories()
  },
  data() {
    return {
      name: "Stranger",
      numberOfBook: 199,
      numberOfReview: 299,
      numberOfNewBook: 13,
      favoriteCategories: [],
    }
  },
  methods: {
    logout () {
      session_apis.logout();
    },
    fetchNumberOfNewBook() {
      const d = new Date()
      const self = this
      statisticApis.countBookByYear(d.getFullYear()).then(function(response) {
        self.numberOfBook = response.data.result[0].book_count
      })
    },
    fetchFavoriteCategories() {
      const self = this
      statisticApis.getFavoriteCategories().then(function(response) {
        self.favoriteCategories = response.data.result
      })
    }
  }
}
</script>
