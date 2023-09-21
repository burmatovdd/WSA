<script>
import {defineComponent} from "vue";
import Sidebar from "./components/Sidebar/sidebar.vue";

export default defineComponent({
  components: {
    Sidebar,
  }
})
</script>

<template>
  <router-view></router-view>
</template>

<style lang="scss">
@use 'app';
</style>
