<template>
  <header class="header">
    <Sidebar/>
  </header>
  <div class="main__container">
    <div class="main__statistic">
      <Block>
        <Card>
          <div class="reports__content card__content">
            <div class="content__title">Ресурсы</div>
            <div class="content__data reports__content-data">
              <List :all-resources="this.allUrl"/>
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
      allUrl: [],
      errUrl: [],
      wafUrl: [],
    }
  },
  mounted() {
    let sendUrl = "http://localhost:8080/api/statistic";

    httpClient.Get(sendUrl).then(response => {
      let resp = JSON.parse(response.data.body)
      console.info(resp)
      this.allUrl = resp.allURL
      this.errUrl = resp.errURL
      this.wafUrl = resp.wafURL
    })
  }
})
</script>

<style lang="scss">
@use 'statistic';
</style>