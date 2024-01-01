<template>
  <div id="home">
    <Breadcrumb paths="Books" />
  </div>
</template>

<script>
import Breadcrumb from '@/components/Breadcrumb.vue';

export default {
  name: "BookView",
  components: {
    Breadcrumb,
  },
  beforeMounted() {
    // TODO: get list books
  }
}
</script>
