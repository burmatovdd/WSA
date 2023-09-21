<template>
<header class="header">
  <Sidebar/>
</header>
  <div class="main__container">
    <main class="main__analytics">
      <Block>
        <Card>
        </Card>
      </Block>
        <Block>
          <Card>
            <div class="chart chart--WAF">
              <h4 class="chart__title">WAF</h4>
              <div class="chart__container chart__container-waf">
                <chart/>
              </div>
            </div>
          </Card>
        </Block>
    </main>
  </div>

</template>

<script>
import {defineComponent} from "vue";
import Sidebar from "../../Sidebar/sidebar.vue";
import Block from "../../Tile/tile.vue";
import Card from "../../Card/card.vue";
import Chart from "./Charts/chart.vue";
export default defineComponent({
  name: "analytics",
  components: {
    Sidebar,
    Block,
    Card,
    Chart,
  }
})
</script>

<style lang="scss">
@use "analytics";
</style>