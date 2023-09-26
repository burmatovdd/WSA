<template>
  <header class="header">
    <Sidebar/>
  </header>
  <div class="main__container">
    <main class="main__reports">
      <Block>
        <Card>
          <div class="reports__content card__content">
            <div class="content__title">Текущая неделя</div>
            <div class="content__data reports__content-data">
              <List :no-resolve="this.currentNoRes" :new-waf="this.currentNewWaf"/>
            </div>
          </div>
        </Card>
      </Block>
      <Block>
        <Card>
          <div class="card__content reports__content">
            <div class="content__title">Прошлая неделя</div>
            <div class="content__data reports__content-data">
              <List :no-resolve="this.lastNoRes" :new-waf="this.lastNewWaf"/>
            </div>
          </div>
        </Card>
      </Block>
    </main>
  </div>
</template>

<script>
import {defineComponent} from 'vue';
import * as httpClient from "../../../httpClient.js";
import Sidebar from "../../Sidebar/sidebar.vue";
import Block from "../../Tile/tile.vue";
import Card from "../../Card/card.vue";
import List from "./list/list.vue";
export default defineComponent({
  name: "reports",
  components: {
    Sidebar,
    Block,
    Card,
    List
  },
  data: function (){
    return{
      lastNoRes: [],
      lastNewWaf: [],
      currentNoRes: [],
      currentNewWaf: [],
    }
  },
  // mounted() {
  //
  //   let sendUrl = "http://localhost:8080/api/week-statistic";
  //
  //   httpClient.Get(sendUrl).then(response => {
  //     let resp = JSON.parse(response.data.body)
  //     console.info(resp)
  //     this.lastNoRes = resp.last.no_res_resource
  //     this.lastNewWaf = resp.last.new_waf_resource
  //     this.currentNoRes = resp.current.no_res_resource
  //     this.currentNewWaf = resp.current.new_waf_resource
  //   })
  // }
})
</script>

<style lang="scss">
@use 'reports';
</style>