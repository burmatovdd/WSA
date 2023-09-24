<template>
<Tile>
  <div class="info">
    <Card>
      <div class="content__title">Общее количество ресурсов - <p class="content__title-num">{{this.genStats.stat.resources}}</p></div>
      <div class="content__time">Данные актуальны на {{this.now}}</div>
    </Card>
    <Card>
      <div class="content__title">Общее количество неактивных ресурсов - <p class="content__title-num">{{this.genStats.stat.deactivateResource}}</p></div>
      <div class="content__time">Данные актуальны на {{this.now}}</div>
    </Card>
    <Card>
      <div class="content__title">Ресурсов за WAF - <p class="content__title-num">{{this.genStats.stat.waf}}</p></div>
      <div class="content__time">Данные актуальны на {{this.now}}</div>
    </Card>
    <Card>
      <div class="content__title">Клиентов - <p class="content__title-num">{{this.genStats.stat.owners}}</p></div>
      <div class="content__time">Данные актуальны на {{this.now}}</div>
    </Card>
  </div>
</Tile>
</template>

<script>
import {defineComponent} from 'vue';
import Card from "../../../Card/card.vue";
import Tile from "../../../Tile/tile.vue";
import * as httpClient from "../../../../httpClient.js";
import moment from 'moment'
export default defineComponent({
  name: "info",
  components: {
    Card,
    Tile,
  },
  data: function (){
    var now = moment().format("DD-MM-YYYY HH:mm");
    return{
      genStats:{
        stat: {
          resources: null,
          deactivateResource: null,
          waf: null,
          owners: null,
        },
        allUrl: [],
        errUrl: [],
        wafUrl: [],
        owners: [],
      },
      now,
    }
  },
  mounted() {
    let sendUrl = "http://localhost:8080/api/general-statistic"
    httpClient.Get(sendUrl).then(response => {
      let resp = JSON.parse(response.data.body)
      this.genStats.stat = resp.genStats
      this.genStats.allUrl = resp.allURL
      this.genStats.errUrl = resp.errURL
      this.genStats.wafUrl = resp.wafURL
      this.genStats.owners = resp.owners
    })
  }
})
</script>

<style lang="scss">
@use 'info';
</style>