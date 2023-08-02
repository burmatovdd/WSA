<template>
<div class="res__content">
  <div class="res__name">
    <h4 class="name__title">Название</h4>
    <p class="name__text">{{this.$props.resource.resName}}</p>
  </div>
  <div class="res__info">
    <div class="res__info--item">
      <p class="item__name">Статус</p>
     <div class="item__img" v-if="this.$props.resource.status">
       <img src="../../img/ok.svg" alt="ok">
     </div>
     <div class="item__img" v-else>
       <img src="../../img/bad.svg" alt="bad">
     </div>
    </div>
    <div class="res__info--item">
      <p class="item__name">WAF</p>
      <div class="item__img" v-if="this.$props.resource.waf">
        <img src="../../img/ok.svg" alt="ok">
      </div>
      <div class="item__img" v-else>
        <img src="../../img/bad.svg" alt="bad">
      </div>
    </div>
    <div class="res__info--item">
      <p class="item__name">SSL</p>
      <div class="item__img" v-if="this.$props.resource.ssl">
        <img src="../../img/ok.svg" alt="ok">
      </div>
      <div class="item__img" v-else>
        <img src="../../img/bad.svg" alt="bad">
      </div>
    </div>
  </div>
  <div class="res__date">
    <h4 class="date__title">Дата окончания сертифката</h4>
    <p class="date__content">{{this.$props.resource.date}}</p>
  </div>
  <div class="res__user">
    <h4 class="user__title">Ответсвенное лицо</h4>
    <p class="user__text">{{this.$props.resource.user}}</p>
  </div>
<!--  todo: добавить кнопку редактирования-->
  <button class="res__delete">Удалить</button>
</div>
</template>

<script>
import { defineComponent } from 'vue';
export default defineComponent({
  name: "resInfo.vue",
  props: {
    resource: Object
  },
  data: function (){
    return {
      resStatus: false,
      resWaf: false,
      resSsl: false,
    }
  },
});
</script>

<style lang="scss">
@use "resInfo";
</style>
