<template>
  <header class="header">
    <Sidebar/>
  </header>
  <div class="main__container">
    <div class="main__statistic">
      <Block class="statistic__block statistic__block--resource">
        <Card>
          <div class="reports__content card__content">
            <div class="content__title">Ресурсы</div>
            <div class="content__data reports__content-data">
              <List :info1="this.all" :info2="this.error"/>
            </div>
          </div>
        </Card>
      </Block>
      <Block>
        <Card>
          <div class="reports__content card__content">
            <div class="content__title">Защищенные ресурсы</div>
            <div class="content__data reports__content-data">
              <List :info1="this.waf" :info2="this.kdp"/>
            </div>
          </div>
        </Card>
      </Block>
    </div>
  </div>
</template>

<script>
import {defineComponent} from "vue";
import Sidebar from "../../Sidebar/sidebar.vue";
import Block from "../../Tile/tile.vue";
import Card from "../../Card/card.vue";
import List from "./list/list.vue";
import * as httpClient from "../../../httpClient.js";
export default defineComponent( {
  name: "statistic",
  components: {
    Sidebar,
    Block,
    Card,
    List,
  },
  data: function (){
    return{
      all: {
        title: 'Все ресурсы',
        array: []
      },
      error: {
        title: 'Неактивные',
        array: []
      },
      waf: {
        title: 'За WAF',
        array: []
      },
      kdp: {
        title: 'За КДП',
        array: [],
      }
    }
  },
  // mounted() {
  //
  //   let sendUrl = "http://localhost:8080/api/statistic";
  //
  //   httpClient.Get(sendUrl).then(response => {
  //     let resp = JSON.parse(response.data.body)
  //     console.info(resp)
  //     this.all.array = resp.allURL
  //     this.error.array = resp.errURL
  //     this.waf.array = resp.wafURL
  //   })
  // }
})
</script>

<style lang="scss">
@use 'statistic';
</style>