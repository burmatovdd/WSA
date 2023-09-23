<template>
  <header class="header" :inert="openedModalDialog">
    <Sidebar/>
  </header>
  <div class="main__container">
    <main class="main__dashboard" :inert="openedModalDialog">
      <Checker @checkerResult="getResult"/>
      <Info/>
      <Reports/>
      <Certificates/>
    </main>
  </div>
  <DialogComp :open="openedModalDialog" @close-modal-dialog="closeModalDialog">
    <Card>
      <h4 class="card__title">
        <template v-if="!this.$data.resourceInfo.found">Ресурс не найден</template>
        <template v-else>Общая информация</template>
      </h4>
      <Adder v-if="!this.$data.resourceInfo.found" :resource="this.$data.resourceInfo.resName"/>
      <ResourceInfo v-else :resource="this.$data.resourceInfo"/>
    </Card>
  </DialogComp>
</template>

<script>
import Sidebar from "../../Sidebar/sidebar.vue";
import Checker from "./Check-Resource/checker.vue";
import Info from "./General-Info/info.vue";
import Reports from "./General-Reports/reports.vue";
import Certificates from "./Certificates/certificates.vue";
import DialogComp from "./Modal/dialog.vue";
import Adder from "./Add-Resource/adder.vue";
import ResourceInfo from "./Resource-Info/resource.vue";
import Card from "../../Card/card.vue";
import {defineComponent} from 'vue';
export default defineComponent({
  name: "main-page",
  components: {
    Sidebar,
    Checker,
    Info,
    Reports,
    Certificates,
    DialogComp,
    Adder,
    ResourceInfo,
    Card
  },
  data: function (){
    return{
      openedModalDialog: false,
      resourceInfo: Object,
    }
  },
  methods: {
    getResult: function (properties){
      this.$data.resourceInfo = properties
      this.$data.openedModalDialog = !this.$data.openedModalDialog;
    },
    openModalDialog: function () {
      this.$data.openedModalDialog = true;
    },
    closeModalDialog: function () {
      this.$data.openedModalDialog = false;
    }
  },
})
</script>

<style lang="scss">
@use 'main';
</style>