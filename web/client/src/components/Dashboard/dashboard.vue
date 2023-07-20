<template>
  <header class="header">
    <Sidebar/>
  </header>
  <main class="main main--dashboard">
    <select type="date" class="dropDown dropDown--date dropDown--year">
      <option value="2023" selected>2023</option>
      <option value="2022">2022</option>
    </select>
    <select type="date" class="dropDown dropDown--date dropDown--month">
      <option value="july" selected>Июль</option>
      <option value="june">Июнь</option>
    </select>
    <div class='chart chart--activeRes'/>
    <div class='chart chart--commonInfo'/>
    <div class="chart chart--report">
      <h4 class="chart__title">Отчет</h4>
      <div class="chart__container">
        <div class="chart__week chart__week--last">
          <h5 class="chart__week-title">За прошлую неделю</h5>
          <p class="chart__week-text">Неактивные : {{ lastWeekNoActive }}</p>
          <p class="chart__week-text">Новые WAF : {{ lastWeekWaf }}</p>
        </div>
        <div class="chart__week--current">
          <h5 class="chart__week-title">За эту неделю</h5>
          <p class="chart__week-text">Неактивные : {{ currentWeekNoActive }}</p>
          <p class="chart__week-text">Новые WAF : {{ currentWeekWaf }}</p>
        </div>
      </div>
    </div>
    <div class="chart chart--SSL">
      <h4 class="chart__title">SSL OK</h4>
      <div class="chart__container">
        <SSLOk/>
      </div>
    </div>
    <div class="chart chart--WAF">
      <h4 class="chart__title">WAF</h4>
      <div class="chart__container">
        <Waf/>
      </div>
    </div>
  </main>
</template>

<script>
import Sidebar from "../Sidebar/sidebar.vue";
import SSLOk from "../Charts/sslOkChart/sslOk.vue";
import Waf from "../Charts/wafChart/waf.vue";
import *as httpClient from "../../httpClient";
import { defineComponent } from 'vue';
export default defineComponent({
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
