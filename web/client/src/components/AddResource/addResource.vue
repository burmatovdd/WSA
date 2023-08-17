<template>
<div class="res__content">
  <p class="content__text">Хотите добавить {{this.$props.resource}} ?</p>
  <Form class="add__form" @submit="onSubmit" >
    <p class="form__text">Ответственное лицо</p>
    <Field name="user"
           type="text"
           class="form__input"
           placeholder="Введите почту"
           v-model="user"/>
    <p class="form__text">Организация</p>
    <Field name="owner"
           type="text"
           class="form__input"
           placeholder="Введите организацию"
           v-model="owner"/>
    <p class="form__text">Если данные не известны, оставьте поля пустыми</p>
    <div class="add__result" v-if="isOpen">
      <p class="add__result--text" v-if="!statusOk">Ресурс добавлен!</p>
      <p class="add__result--text" v-else>Ошибка!</p>
    </div>
    <button type="button" class="res__button" @click="onSubmit">Добавить</button>
  </Form>
</div>
</template>

<script>
import { defineComponent } from 'vue';
import * as httpClient from "../../httpClient.js";
import {Form, Field, ErrorMessage} from 'vee-validate';
export default defineComponent({
  name: "addResource.vue",
  components: {
    Form,
    Field,
    ErrorMessage
  },
  props: {
    resource: {
      type: String,
      default: "fallback content"
    }
  },
  data: function (){
    return {
      user: null,
      owner: null,
      isOpen: false,
      statusOk: false,
    }
  },
  methods: {
    onSubmit(){
      let sendUrl = "http://localhost:8080/api/add-resource";

      return httpClient.Post(sendUrl,{
        url: this.$props.resource,
        email: this.$data.user,
        owner: this.$data.owner
      }).then(response =>{
        this.$data.isOpen = true
      })
        .catch(error => {
        if (error.response.data.code === 500){
          this.$data.statusOk = false
        }
      })
    },
  }
});

</script>

<style lang="scss">
@use 'addResource';
</style>
