<template>
  <div class="sidebar">
    <div class='sidebar__row'>
      <a href="/dashboard" class="sidebar__logo">
        <img src="../../img/logo.png" alt="logo" width="97" height="146" class="sidebar__logo-img">
      </a>
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
@use "sidabar";
</style>
