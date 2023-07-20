<template>
<div class="sidebar">
  <div class="sidebar-logo">
    <img src="../../img/logo.png" alt="logo" class="sidebar-logo_img">
  </div>
  <div class="sidebar-icons">
    <div class="sidebar-icon">
      <button class="sidebar-icon_button" @click="toDashboard">
        <img src="../../img/dashboard.svg" alt="dashboard-icon" class="sidebar-icon_button-img">
      </button>
      <p class="sidebar-icon_text">Дашборд</p>
    </div>
    <div class="sidebar-icon">
      <button class="sidebar-icon_button" @click="toOwners">
        <img src="../../img/owner.svg" alt="dashboard-icon" class="sidebar-icon_button-img">
      </button>
      <p class="sidebar-icon_text">Организации</p>
    </div>
    <div class="sidebar-icon">
      <button class="sidebar-icon_button" @click="exit">
        <img src="../../img/exit.svg" alt="dashboard-icon" class="sidebar-icon_button-img">
      </button>
      <p class="sidebar-icon_text">Выход</p>
    </div>
  </div>
</div>
</template>

<script>
import {defineComponent} from 'vue';
export default defineComponent( {
  name: "sidebar",
  data: function () {
    return {
    }
  },
  methods: {
    toDashboard(){
      this.$router.push('/dashboard')
    },
    toOwners(){
      console.log("in work")
    },
    exit(){
      this.$router.push('/')
    },
  }
})
</script>

<style lang="scss">
@use "sidabar";
</style>