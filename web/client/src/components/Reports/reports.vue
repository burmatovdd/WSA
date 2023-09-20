<template>
<div class="reports">
  <h4 class="reports__title">Отчёты</h4>
  <div class="reports__content">
    <div class="current__week">
      <h4 class="current__week--title">Текущая неделя</h4>
      <div class="current__week--content">
        <div class="current__week--noActive">
          <p class="noActive__title">Неактивные</p>
          <div class="noActive__data">
            <div class="data__info">
              <div class="data__info--resource">Ресурс</div>
              <div class="data__info--date">Дата</div>
            </div>
            <ul>
              <li v-for="item in currentNoRes">
                <div class="data__content">
                  <p class="data__content--resource">{{ item.resource }}</p>
                  <p class="data__content--date">{{ item.date }}</p>
                </div>
              </li>
            </ul>
          </div>
        </div>
        <div class="current__week--Waf">
          <p class="Waf__title">Новые WAF</p>
          <div class="Waf__data">
            <div class="data__info">
              <div class="data__info--resource">Ресурс</div>
              <div class="data__info--date">Дата</div>
            </div>
            <ul>
              <li v-for="item in currentNewWaf">
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
    <div class="last__week">
      <h4 class="last__week--title">Прошлая неделя</h4>
      <div class="last__week--content">
        <div class="last__week--WAF">
          <p class="WAF__title">Неактивные</p>
          <div class="WAF__data">
            <div class="data__info">
              <div class="data__info--resource">Ресурс</div>
              <div class="data__info--date">Дата</div>
            </div>
            <ul>
              <li v-for="item in lastNoRes">
                <div class="data__content">
                  <p class="data__content--resource">{{ item.resource }}</p>
                  <p class="data__content--date">{{ item.date }}</p>
                </div>
              </li>
            </ul>
          </div>
        </div>
        <div class="last__week--Waf">
          <p class="Waf__title">Новые WAF</p>
          <div class="Waf__data">
            <div class="data__info">
              <div class="data__info--resource">Ресурс</div>
              <div class="data__info--date">Дата</div>
            </div>
            <ul>
              <li v-for="item in lastNewWaf">
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
  </div>
</div>
</template>

<script>
import * as httpClient from "../../httpClient";

export default {
  name: "reports",
  data: () => {
    return{
      lastNoRes: [],
      lastNewWaf: [],
      currentNoRes: [],
      currentNewWaf: [],
    }
  },
  mounted() {
    let sendUrl = "http://localhost:8080/api/week-statistic";

    httpClient.Get(sendUrl).then(response => {
      let resp = JSON.parse(response.data.body)
      this.lastNoRes = resp.last.no_res_resource
      this.lastNewWaf = resp.last.new_waf_resource
      this.currentNoRes = resp.current.no_res_resource
      this.currentNewWaf = resp.current.new_waf_resource
    })
  }
}
</script>

<style lang="scss">
@use "reports";
</style>
