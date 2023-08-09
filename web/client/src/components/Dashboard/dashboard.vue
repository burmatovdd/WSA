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
      <Form class="form__check" v-slot="{ meta }">
        <ErrorMessage name="login" class="error-message"/>
        <Field name="search"
               type="url"
               class="form__search"
               placeholder="Название ресурса"
               v-model="resourceName"
               :rules="validateLogin"/>
        <button type="button" class="form__submit" @click="toggleModalDialog" :disabled="!meta.valid">Проверить ресурс</button>
      </Form>
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
      <div class="chart__container chart__container--week">
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
    @close-modal-dialog="closeModalDialog">
    <div v-if="!find" class="modal__content">
      <h4 class="modal__title">Проверка ресурса</h4>
      <AddResource :resource="resourceName"/>
    </div>
    <div v-else class="modal__content">
      <h4 class="modal__title">Проверка ресурса</h4>
      <ResInfo :resource="resourceInfo"/>
    </div>
  </Modal>
</template>

<script>
import Sidebar from "../Sidebar/sidebar.vue";
import SSLOk from "../Charts/sslOkChart/sslOk.vue";
import Waf from "../Charts/wafChart/waf.vue";
import Modal from '../../components/Modal/modal.vue';
import AddResource from "../AddResource/addResource.vue";
import ResInfo from "../ResInfo/resInfo.vue";
import {Form, Field, ErrorMessage} from 'vee-validate';
import * as httpClient from "../../httpClient";
import { defineComponent } from 'vue';
import {getData} from "../Charts/sslOkChart/PieConfig.js";
export default defineComponent({
  name: "dashboard.vue",
  components: {
    Sidebar,
    SSLOk,
    Waf,
    Modal,
    AddResource,
    ResInfo,
    Form,
    Field,
    ErrorMessage
  },
  data: function () {
    return {
      openedModalDialog: false,
      find: false,
      resourceName: null,
      resourceInfo: {
        resName: null,
        status: false,
        waf: false,
        ssl: false,
        date: null,
        user: null,
      },
      lastWeekNoActive: null,
      lastWeekWaf: null,
      currentWeekNoActive: null,
      currentWeekWaf: null,
      isActive: false,
    }
  },
  mounted() {
    let sendUrl = "http://localhost:8080/api/week-statistic";

    httpClient.Get(sendUrl).then(response => {
      let resp = JSON.parse(response.data.body)
      console.info(resp)
      this.lastWeekNoActive = resp.last.no_resolve
      this.lastWeekWaf = resp.last.new_waf
      this.currentWeekNoActive = resp.current.no_resolve
      this.currentWeekWaf = resp.current.new_waf
    })
  },
  methods: {
    validateLogin(value) {
      this.isActive = false
      // if the field is empty
      if (!value) {
        return null;
      }
      // All is good
      return true;
    },
    toggleModalDialog() {

      this.$data.openedModalDialog = !this.$data.openedModalDialog;
      let sendUrl = "http://localhost:8080/api/check-resource";

      // let today = Date.now()
      // console.info(today)
      // admin.parki.mosreg.ru
      // 2018.llo.zdravmo.mosreg.ru

      return httpClient.Post(sendUrl,this.$data.resourceName).then(response =>{
        let resp = JSON.parse(response.data.body)

        if (resp.DateEnd === "--------------------" || resp.DateEnd === "" ){
          this.resourceInfo.date = "undefined"
        }
        else{
          this.resourceInfo.date = resp.DateEnd
        }
        if (resp.Email === ""){
          this.resourceInfo.user = "undefined"
        }
        else{
          this.resourceInfo.user = resp.Email
        }
        this.resourceInfo.resName = resp.URL
        this.resourceInfo.status = resp.Status
        this.resourceInfo.waf = resp.WAF
        this.resourceInfo.ssl = resp.SSL
        this.$data.find = true

      }).catch(error => {
        if (error.response.data.code === 500){
          this.$data.find = false
        }
      })
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
