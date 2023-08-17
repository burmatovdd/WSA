<template>
<div class="report">
  <h4 class="report__title">Отчет</h4>
  <div class="report__container">
    <div class="report__week report__week--last">
      <h5 class="report__week-title">За прошлую неделю</h5>
      <p class="report__week-text report__week-text-noActive">Неактивные : {{ lastWeekNoActive }}</p>
      <p class="report__week-text">Новые WAF : {{ lastWeekWaf }}</p>
    </div>
    <div class="report__week--current">
      <h5 class="report__week-title">За эту неделю</h5>
      <p class="report__week-text report__week-text-noActive">Неактивные : {{ currentWeekNoActive }}</p>
      <p class="report__week-text">Новые WAF : {{ currentWeekWaf }}</p>
    </div>
</div>
</div>
</template>

<script>
import * as httpClient from "../../httpClient";
export default {
  name: "report",
  data: () => {
    return{
      lastWeekNoActive: null,
      lastWeekWaf: null,
      currentWeekNoActive: null,
      currentWeekWaf: null,
    }
  },
  mounted() {
    let sendUrl = "http://localhost:8080/api/week-statistic";

    httpClient.Get(sendUrl).then(response => {
      let resp = JSON.parse(response.data.body)
      this.lastWeekNoActive = resp.last.no_resolve
      this.lastWeekWaf = resp.last.new_waf
      this.currentWeekNoActive = resp.current.no_resolve
      this.currentWeekWaf = resp.current.new_waf
    })
  }
}
</script>

<style lang="scss">
@use 'report';
</style>
