<template>
<Tile>
  <div class="reports">
    <Card>
      <div class="card__content">
        <div class="content__title">Прошлая неделя</div>
        <div class="content__deactivate">Неактивные ресурсы ({{this.lastWeekNoActive}})</div>
        <div class="content__waf">Новые WAF ресурсы ({{this.lastWeekWaf}})</div>
      </div>
    </Card>
    <Card>
      <div class="card__content">
        <div class="content__title">Текущая неделя</div>
        <div class="content__deactivate">Неактивные ресурсы ({{this.currentWeekNoActive}})</div>
        <div class="content__waf">Новые WAF ресурсы ({{this.currentWeekWaf}})</div>
      </div>
    </Card>
  </div>
</Tile>
</template>

<script>
import {defineComponent} from 'vue';
import Card from "../../../Card/card.vue";
import Tile from "../../../Tile/tile.vue";
import * as httpClient from "../../../../httpClient.js";
export default defineComponent({
  name: "reports",
  components: {
    Card,
    Tile,
  },
  data: function (){
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
})
</script>

<style lang="scss">
@use 'reports';
</style>