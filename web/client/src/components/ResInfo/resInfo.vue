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
    <div class="user__content">
      <p class="user__content--text" v-if="!isEdit">{{this.$props.resource.user}}</p>
      <div class="user__edit--content" v-else>
        <div class="content__comment" v-if="isOk">
          <p class="comment__text" v-if="!statusEdit">Ошибка!</p>
        </div>
        <div class="content__main">
          <Form  v-slot="{ meta }">
            <Field name="search"
                   type="text"
                   class="content__input"
                   placeholder="Введите почту"
                   v-model="user"
                   :rules="validateInput"/>
            <button type="button" class="content__agree" @click="editUser" :disabled="!meta.valid">Изменить</button>
          </Form>
        </div>
      </div>
      <button class="user__content--edit" @click="openEditModal" v-if="!isEdit"><img src="../../img/edit.svg" alt="edit"></button>
    </div>

  </div>
  <div class="delete__result" v-if="isOpen">
    <p class="delete__result--text" v-if="statusOk">Ресурс удален!</p>
    <p class="delete__result--text" v-else>Ошибка!</p>
  </div>
<!--  todo: добавить кнопку редактирования-->
  <button class="res__delete" @click="deleteResource">Удалить</button>
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
    console.info(this.$props.resource)
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
