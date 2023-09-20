<link rel="stylesheet" href="resInfo.scss">
<template>
<div class="res__content">
  <div class="res__name">
    <div class="res__name--content">
      <h4 class="name__title">Название:</h4>
      <p class="name__text">{{this.$props.resource.resName}}</p>
    </div>
    <div class="res__name--content">
      <h4 class="name__title">IP:</h4>
      <p class="name__text">{{this.$props.resource.ip}}</p>
    </div>
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
    <div class="user__content">
      <div class="user__content--container" v-if="!isEdit">
        <p class="user__content--text">{{this.$props.resource.email}}</p>
        <p class="user__content--text user__content--text-fio">{{this.$props.resource.fio}}</p>
      </div>
      <div class="user__edit--content" v-else>
        <div class="content__comment" v-if="isOk">
          <p class="comment__text" v-if="!statusEdit">Ошибка!</p>
        </div>
        <div class="content__main">
          <Form>
            <Field name="search"
                   type="text"
                   class="content__input"
                   placeholder="Введите почту"
                   v-model="user"/>
            <button type="button" class="content__agree" @click="editUser">Изменить</button>
          </Form>
        </div>
      </div>
    </div>

  </div>
  <div class="delete__result" v-if="isOpen">
    <p class="delete__result--text" v-if="statusOk">Ресурс удален!</p>
    <p class="delete__result--text" v-else>Ошибка!</p>
  </div>
  <div class="buttons__container">
    <button class="res__button" @click="openEditModal" v-if="!isEdit">Изменить</button>
    <button class="res__button" @click="closeEditModal" v-if="isEdit">Назад</button>
    <button class="res__button res__button--delete" @click="deleteResource">Удалить</button>
  </div>
</div>
</template>

<script>
import { defineComponent } from 'vue';
import * as httpClient from "../../httpClient.js";
import {Form, Field, ErrorMessage} from 'vee-validate';
export default defineComponent({
  name: "resInfo.vue",
  components: {
    Form,
    Field,
    ErrorMessage,
  },
  props: {
    resource: Object
  },
  data: function (){
    return {
      isEdit: false,
      isOpen: false,
      isOk: false,
      statusOk: false,
      statusEdit: false,
      user: null
    }
  },
  methods:{
    validateInput(value) {
      this.isActive = false
      // if the field is empty
      if (!value) {
        return null;
      }
      // All is good
      return true;
    },
    deleteResource: function(){
      let sendUrl = "http://localhost:8080/api/delete-resource";

      return httpClient.Post(sendUrl,this.$props.resource.resName).then(response =>{
        this.$data.isOpen = true
        this.$data.statusOk = true
      })
        .catch(error => {
        if (error.response.data.code === 500){
          this.$data.isOpen = true
          this.$data.statusOk = false
        }
      })
    },
    openEditModal: function (){
      this.$data.isEdit = true
    },
    closeEditModal: function (){this.$data.isEdit = false},
    editUser: function (){
      let sendUrl = "http://localhost:8080/api/update-resource";

      return httpClient.Post(sendUrl,{
        url: this.$props.resource.resName,
        email: this.$data.user
      }).then(response =>{
        this.$data.isEdit = false
        this.$data.statusEdit = true
        this.$props.resource.user = this.$data.user
      })
        .catch(error => {
          if (error.response.data.code === 500){
            this.$data.isEdit = true
            this.$data.isOk = true
            this.$data.statusEdit = false
          }
        })
    }
  }
});
</script>

<style lang="scss">
@use "resInfo";
</style>
