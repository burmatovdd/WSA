<template>
  <header class="header" :inert="openedModalDialog">
    <Sidebar/>
  </header>
  <main class="main main--dashboard" :inert="openedModalDialog">
    <form class='main__forhead form form--basic'>
      <select type="date" class="form__dropDown form__dropDown--year">
        <option value="2023" selected>2023</option>
        <option value="2022">2022</option>
      </select>
      <select type="date" class="form__dropDown form__dropDown--month">
        <option value="july" selected>Июль</option>
        <option value="june">Июнь</option>
      </select>
      <input type='url' class="form__search" placeholder="Название ресурса">
      <button type="button" class="form__submit" @click="toggleModalDialog">Проверить ресурс</button>
    </form>
    <div class='chart chart--activeRes'/>
    <div class='chart chart--commonInfo'>
      <div class="chart__title">Общая информация</div>
      <div class="chart__info">
        <div class="chart__info-text">Всего организаций : </div>
        <div class="chart__info-text">Всего ресурсов : </div>
      </div>
    </div>
    <div class="chart chart--report">
      <h4 class="chart__title">Отчет</h4>
      <div class="chart__container">
        <div class="chart__week chart__week--last">
          <h5 class="chart__week-title">За прошлую неделю</h5>
          <p class="chart__week-text">Неактивные : {{ lastWeekNoActive }}</p>
          <p class="chart__week-text">Новые WAF : {{ lastWeekWaf }}</p>
        </div>
        <div class="chart__week--current">
          <h5 class="chart__week-title">За эту неделю</h5>
          <p class="chart__week-text">Неактивные : {{ currentWeekNoActive }}</p>
          <p class="chart__week-text">Новые WAF : {{ currentWeekWaf }}</p>
        </div>
      </div>
    </div>
    <div class="chart chart--SSL">
      <h4 class="chart__title">SSL OK</h4>
      <div class="chart__container">
        <SSLOk/>
      </div>
    </div>
    <div class="chart chart--WAF">
      <h4 class="chart__title">WAF</h4>
      <div class="chart__container">
        <Waf/>
      </div>
    </div>
  </main>
  <Modal
    :open="openedModalDialog"
    @close-modal-dialog="closeModalDialog"
  >
    <h4>Lorem ipsum dolor sit amet.</h4>
    <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Voluptatem, nesciunt!</p>
  </Modal>
</template>

<script>
import Sidebar from "../Sidebar/sidebar.vue";
import SSLOk from "../Charts/sslOkChart/sslOk.vue";
import Waf from "../Charts/wafChart/waf.vue";
import Modal from '../../components/Modal/modal.vue';
import * as httpClient from "../../httpClient";
import { defineComponent } from 'vue';
export default defineComponent({
  name: "dashboard.vue",
  components: {
    Sidebar,
    SSLOk,
    Waf,
    Modal
  },
  data: function () {
    return {
      openedModalDialog: false,
      lastWeekNoActive: null,
      lastWeekWaf: null,
      currentWeekNoActive: null,
      currentWeekWaf: null,
    }
  },
  mounted() {
    let sendUrl = "http://localhost:8080/api/get-week-stat";

    httpClient.Get(sendUrl).then(response => {
      let resp = JSON.parse(response.data.body)
      this.lastWeekNoActive = resp.lastWeek.noResolve
      this.lastWeekWaf = resp.lastWeek.newWaf
      this.currentWeekNoActive = resp.currentWeek.noResolve
      this.currentWeekWaf = resp.currentWeek.newWaf
    })
  },
  methods: {
    toggleModalDialog() {
      this.$data.openedModalDialog = !this.$data.openedModalDialog;
    },
    openModalDialog() {
      this.$data.openedModalDialog = true;
    },
    closeModalDialog() {
      this.$data.openedModalDialog = false;
    }
  }
});
</script>

<style lang="scss">
@use "dashboard";
</style>
