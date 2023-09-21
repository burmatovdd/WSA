<template>
<Tile>
  <div class="certificates">
    <Card>
      <div class="card__content">
        <div class="content__title">Сертификаты, срок действия которых закончится в ближайшие 2 месяца <p class="content__title-num-cert">({{this.num}})</p></div>
        <div class="content__data">
          <div class="data__tables">
            <div class="data__name">Ресурс</div>
            <div class="data__date">Дата</div>
          </div>
          <div class="data__list">
            <ul>
              <li v-for="item in current">
                <div class="data__content">
                  <p class="data__content-resource">{{ item.resource }}</p>
                  <p class="data__content-date">{{ item.date }}</p>
                </div>
              </li>
              <li v-for="item in next">
                <div class="data__content">
                  <p class="data__content-resource">{{ item.resource }}</p>
                  <p class="data__content-date">{{ item.date }}</p>
                </div>
              </li>
            </ul>
          </div>
        </div>
      </div>
    </Card>
  </div>
</Tile>
</template>

<script>
import {defineComponent} from 'vue';
import * as httpClient from "../../../../httpClient.js";
import Card from "../../../Card/card.vue";
import Tile from "../../../Tile/tile.vue";
export default defineComponent({
  name: "certificate",
  components: {
    Card,
    Tile,
  },
  data: function (){
    return {
      current: [],
      next: [],
      num: null,
    }
  },
  mounted() {
    let sendUrl = "http://localhost:8080/api/certificates";

    httpClient.Get(sendUrl).then(response => {
      let resp = JSON.parse(response.data.body)
      this.$data.current = resp.current
      this.$data.next = resp.next
      this.num = resp.current.length + resp.next.length
    })
  }
})
</script>

<style lang="scss">
@use 'certificates';
</style>