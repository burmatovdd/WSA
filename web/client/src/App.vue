<script>
import {defineComponent} from "vue";
import router from "./router.js";
import * as storage from "./storage.js";

export default defineComponent({
  mounted() {
    let token = storage.get("token");
    if (token == null){
      router.push("/");
    }
  }
})
</script>

<template>
  <router-view></router-view>
</template>

<style lang="scss">
@use 'app';
</style>
