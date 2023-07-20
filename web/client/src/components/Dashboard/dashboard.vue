<template>
  <header><Sidebar/></header>
  <div class="container container-dashboard">
    <div class="bottomCharts">
      <div class="chart chart-report">
        <p class="chart-title">Отчет</p>
        <div class="chart-container">
          <div class="last-week">
            <p class="last-week_title">За прошлую неделю</p>
            <p class="last-week_text">Неактивные : {{lastWeekNoActive}}</p>
            <p class="last-week_text">Новые WAF : {{lastWeekWaf}}</p>
          </div>
          <div class="current-week">
            <p class="current-week_title">За эту неделю</p>
            <p class="current-week_text">Неактивные : {{currentWeekNoActive}}</p>
            <p class="current-week_text">Новые WAF : {{currentWeekWaf}}</p>
          </div>
        </div>
      </div>
      <div class="chart">
        <p class="chart-title">SSL OK</p>
        <div class="chart-container">
          <SSLOk/>
        </div>
      </div>
      <div class="chart">
        <p class="chart-title">WAF</p>
        <div class="chart-container">
          <Waf/>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Sidebar from "../Sidebar/sidebar.vue";
import SSLOk from "../Charts/sslOkChart/sslOk.vue";
import Waf from "../Charts/wafChart/waf.vue";
import *as httpClient from "../../httpClient";
import {defineComponent} from 'vue';
export default defineComponent( {
  name: "dashboard.vue",
  components: {
    Sidebar,
    SSLOk,
    Waf
  },
  data: function () {
    return {
      lastWeekNoActive: null,
      lastWeekWaf: null,
      currentWeekNoActive: null,
      currentWeekWaf: null,
    }
  },
  mounted() {
    let sendUrl = "http://localhost:8080/api/get-week-stat";

    httpClient.Get(sendUrl).then(response => {
      let resp = JSON.parse(response.data.body)
      this.lastWeekNoActive = resp.lastWeek.noResolve
      this.lastWeekWaf = resp.lastWeek.newWaf
      this.currentWeekNoActive = resp.currentWeek.noResolve
      this.currentWeekWaf = resp.currentWeek.newWaf
    })
  }
})
</script>

<style lang="scss">
@use "dashboard";
</style>