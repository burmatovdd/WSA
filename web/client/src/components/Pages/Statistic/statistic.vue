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
              <List :no-resolve="this.currentNoRes" :new-waf="this.currentNewWaf"/>
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
export default defineComponent( {
  name: "statistic",
  components: {
    Sidebar,
    Block,
    Card,
    List,
  }
})
</script>

<style lang="scss">
@use 'statistic';
</style>