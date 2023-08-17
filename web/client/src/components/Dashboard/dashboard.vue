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
    <div class='chart chart--activeRes'>
      Free content
    </div>
    <div class='chart chart--commonInfo'>
      <CommonInfo/>
    </div>
    <div class="chart chart--SSL">
      Free content
    </div>
    <div class="chart chart--WAF">
      <h4 class="chart__title">WAF</h4>
      <div class="chart__container">
        <Waf/>
      </div>
    </div>
    <div class="chart chart--report">
      <Report/>
    </div>
    <div class="chart chart--reports">
      <Reports/>
    </div>
    <div class="chart chart--certificates">
        <Certificate/>
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
import Certificate from "../CertificateInfo/certificateInfo.vue";
import Waf from "../wafChart/waf.vue";
import Modal from '../../components/Modal/modal.vue';
import AddResource from "../AddResource/addResource.vue";
import ResInfo from "../ResInfo/resInfo.vue";
import Report from "../ReportGeneral/report.vue";
import Reports from "../Reports/reports.vue";
import CommonInfo from "../CommonInfo/commonInfo.vue";
import {Form, Field, ErrorMessage} from 'vee-validate';
import * as httpClient from "../../httpClient";
import { defineComponent } from 'vue';
export default defineComponent({
  name: "dashboard.vue",
  components: {
    Sidebar,
    Certificate,
    Waf,
    Modal,
    AddResource,
    ResInfo,
    CommonInfo,
    Report,
    Reports,
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
        ip: null,
        status: false,
        waf: false,
        ssl: false,
        date: null,
        email: null,
        fio: null,

      },
      isActive: false,
      isOpen: false,
      isOpenMore: false,

    }
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

      return httpClient.Post(sendUrl,this.$data.resourceName).then(response =>{
        let resp = JSON.parse(response.data.body)

        if (resp.IP === "" || resp.IP === "----------"){
          this.resourceInfo.ip = "undefined"
        }else{
          this.resourceInfo.ip = resp.IP
        }

        if (resp.DateEnd === "--------------------" || resp.DateEnd === "" ){
          this.resourceInfo.date = "undefined"
        }
        else{
          this.resourceInfo.date = resp.DateEnd
        }
        if (resp.Email === ""){
          this.resourceInfo.email = "undefined"
          this.resourceInfo.fio = "undefined"
        }
        else{
          this.resourceInfo.email = resp.Email
          this.resourceInfo.fio = resp.FIO
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
