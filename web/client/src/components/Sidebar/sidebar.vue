<template>
  <div class="sidebar">
    <div class='sidebar__row'>
      <div class="sidebar__title">
        <a href="/dashboard" class="sidebar__logo">
          <img src="../../images/logo.svg" alt="logo" class="sidebar__logo-img">
        </a>
        <div class="sidebar__title-text">
          <p class="title__place">МОСКОВСКАЯ ОБЛАСТЬ</p>
          <p class="title__name">WSA</p>
        </div>
      </div>
      <Nav/>
    </div>
  </div>
</template>

<script>
import Nav from "../Nav/nav.vue";
import { defineComponent } from 'vue';
export default defineComponent({
  name: "custom-sidebar",
  components: {
    Nav
  }
})
</script>

<style lang="scss">
@use 'sidebar';
</style>