<template>
<div class="common">
  <div class="common__title">Общая информация</div>
  <div class="common__info">
    <div class="common__info--text">Всего ресурсов</div>
    <div class="common__info--text">{{resources}}</div>
    <div class="common__info--text">Всего за WAF</div>
    <div class="common__info--text">{{waf}}</div>
    <div class="common__info--text">Всего организаций</div>
    <div class="common__info--text">{{owners}}</div>
  </div>
</div>
</template>

<script>
import * as httpClient from "../../httpClient";
import { defineComponent } from 'vue';
export default defineComponent( {
  name: "commonInfo",
  data: () => {
    return{
      resources: null,
      waf: null,
      owners: null,
    }
  },
  mounted() {
    let sendUrl = "http://localhost:8080/api/general-statistic"
    httpClient.Get(sendUrl).then(response => {
      let resp = JSON.parse(response.data.body)
      this.resources = resp.resources
      this.owners = resp.owners
      this.waf = resp.waf
    })
  }
})
</script>

<style lang="scss">
@use 'commonInfo';
</style>
