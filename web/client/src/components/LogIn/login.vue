<template>
<div class="container container-login">
    <h1 class="login-title">Welcome to Waf Analytics</h1>
    <Form class="login-form" @submit="onSubmit" v-slot="{ meta }">
      <ErrorMessage name="login" class="error-message"/>
      <p class="error-message" :class="{noError: !isActive}">Неправильный логин или пароль</p>
      <Field name="login"
             type="text"
             class="login-form_input login-form_input--login"
             placeholder="login"
             :rules="validateLogin"/>

      <Field name="password"
             type="password"
             class="login-form_input"
             placeholder="password"
             :rules="validatePassword"/>
      <button class="button active"
              id = "button"
              @click="login"
              :disabled="!meta.valid">
        login
      </button>
    </Form>
  </div>
</template>

<script>
import {defineComponent} from 'vue';
import {Form, Field, ErrorMessage} from 'vee-validate';
import *as httpClient from "../../httpClient";
export default defineComponent( {
  name: "login",
  components: {
    Form,
    Field,
    ErrorMessage
  },
  data: function () {
    const user = {
      login: null,
      password: null
    }
    return {
      user,
      resp : null,
      errored: false,
      isActive: false,
    }
  },
  methods: {
    onSubmit(values) {
      this.user.login = values.login;
      this.user.password = values.password;
    },
    validateLogin(value) {
      this.isActive = false
      // if the field is empty
      if (!value) {
        return 'Поля должны быть заполнеными';
      }
      // if the field is not a valid email
      const regex = /^[A-Za-z0-9]/i;
      if (!regex.test(value)) {
        return 'логин может содержать только латинские буквы и цифры';
      }
      this.user.login = value;
      // All is good
      return true;
    },
    validatePassword(value) {
      this.isActive = false
      if (!value) {
        return 'Поля должны быть заполнеными';
      }
      this.user.password = value;
      return true;
    },
    login: async function () {

      let sendUrl = "http://localhost:8080/api/login";

      let postInfo = httpClient.Post(sendUrl, this.user)

      postInfo.then(response => {
        this.$router.push('/dashboard');
      })
      .catch(err => {
        this.errored = true;
        this.isActive = !this.isActive;
      });
    },
  }
})
</script>

<style lang="scss">
@use "login";
</style>
