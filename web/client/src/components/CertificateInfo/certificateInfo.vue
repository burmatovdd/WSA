<template>
  <div class="certificate">
    <h4 class="certificate__title">Сертификаты, срок действия которых закончится в ближайшие 2 месяца</h4>
    <div class="certificate__content">
        <div class="current__month">
          <h4 class="current__month--title">Текущий месяц</h4>
          <div class="current__month--data">
            <div class="data__info">
              <div class="data__info--resource">Ресурс</div>
              <div class="data__info--date">Дата</div>
            </div>
            <ul>
              <li v-for="item in current">
                <div class="data__content">
                  <p class="data__content--resource">{{ item.resource }}</p>
                  <p class="data__content--date">{{ item.date }}</p>
                </div>
              </li>
            </ul>
          </div>
        </div>
        <div class="next__month">
          <h4 class="next__month--title">Следующий месяц</h4>
          <div class="next__month--data">
            <div class="data__info">
              <div class="data__info--resource">Ресурс</div>
              <div class="data__info--date">Дата</div>
            </div>
            <ul>
              <li v-for="item in next">
                <div class="data__content">
                  <p class="data__content--resource">{{ item.resource }}</p>
                  <p class="data__content--date">{{ item.date }}</p>
                </div>
              </li>
            </ul>
          </div>
        </div>
    </div>
  </div>
</template>

<script>
import * as httpClient from "../../httpClient";
export default {
  name: 'certificate',
  data:  () => {
    return {
      current: [],
      next: [],
    }
  },
  mounted() {
    let sendUrl = "http://localhost:8080/api/certificates";

    httpClient.Get(sendUrl).then(response => {
      let resp = JSON.parse(response.data.body)
      this.$data.current = resp.current
      this.$data.next = resp.next
    })
  }
}
</script>
<style lang="scss">
@use 'certificateInfo';
</style>
