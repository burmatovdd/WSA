<template>
<div class="add">
  <Form class="add__form" name="form" v-slot="{ meta }">
    <Field name="resource-name"
           type="text"
           class="add__form-input add__form-input-resource"
           :placeholder="this.$props.resource"/>
    <Field name="resource-user"
           type="text"
           class="add__form-input add__form-input-user"
           placeholder="Ответсвтенное лицо"
           v-model="user"/>
    <Field name="resource-owner"
           type="text"
           class="add__form-input add__form-input-owner"
           placeholder="Организация"
           v-model="owner"/>
    <p class="add__form-text">Если данные не известны, оставьте поля пустыми</p>
    <div class="add__result" v-if="isOpen">
      <p class="add__result--text" v-if="statusOk === true">Ресурс добавлен!</p>
      <p class="add__result--text" v-else>Ошибка!</p>
    </div>
    <button class="add__form-button" :disabled="!meta.valid" @click="addResource">Добавить ресурс</button>
  </Form>
</div>
</template>

<script>
import {defineComponent} from "vue";
import {Form, Field} from 'vee-validate';
import * as httpClient from "../../../../httpClient.js";
export default defineComponent({
  name: "adder",
  props: {
    resource: null
  },
  data: function (){
    return{
      user: null,
      owner: null,
      isOpen: false,
      statusOk: null,
    }
  },
  components: {
    Form,
    Field,
  },
  methods: {
    addResource: function (event){
      event.preventDefault()
      let sendUrl = "http://localhost:8080/api/add-resource";

      return httpClient.Post(sendUrl,{
        url: this.$props.resource,
        email: this.$data.user,
        owner: this.$data.owner
      }).then(response =>{
        this.$data.isOpen = true
        this.$data.statusOk = true
      }).catch(error => {
        if (error.response.data.code === 500){
          this.$data.isOpen = true
          this.$data.statusOk = false
        }
      })
    }
  }
})
</script>

<style lang="scss">
@use 'adder';
</style>