<template>
<div class="res__content">
  <p class="content__text">Хотите добавить {{this.$props.resource}} ?</p>
  <Form class="add__form" @submit="onSubmit" v-slot="{ meta }">
    <p class="form__text">Ответственное лицо</p>
    <Field name="user"
           type="text"
           class="form__input"
           placeholder="Введите почту"
           v-model="user"
           :rules="validateInput"/>
    <p class="form__text">Организация</p>
    <Field name="owner"
           type="text"
           class="form__input"
           placeholder="ExampleInc"
           v-model="owner"
           :rules="validateInput"/>
    <p class="form__text">Если данные не известны, поставьте прочерк(-)</p>
    <button type="button" class="res__button" @click="onSubmit" :disabled="!meta.valid">Добавить</button>
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
      owner: null
    }
  },
  methods: {
    onSubmit(){
      let sendUrl = "http://localhost:8080/api/add-resource";

      console.info(JSON.stringify({
        url: this.$props.resource,
        email: this.$data.user,
        owner: this.$data.owner
      }))

      return httpClient.Post(sendUrl,{
        url: this.$props.resource,
        email: this.$data.user,
        owner: this.$data.owner
      }).then(response =>{
        let resp = JSON.parse(response.data.body)
        console.info(resp)
      }).catch(error => {
        if (error.response.data.code === 500){
          console.log("something is going wrong")
        }
      })
    },
    validateInput(value) {
      this.isActive = false
      // if the field is empty
      if (!value) {
        return null;
      }
      // All is good
      return true;
    },
  }
});

</script>

<style lang="scss">
@use 'addResource';
</style>
