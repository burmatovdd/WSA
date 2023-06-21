<template>
<div class="container">
  <div class="login">
    <h1 class="login-title">Welcome to Waf Analytics</h1>
    <Form class="login-form" @submit="onSubmit" v-slot="{ meta }">
      <Field name="login"
             type="text"
             class="login-from_input"
             placeholder="login"
             :rules="validateLogin"/>
      <ErrorMessage name="login" class="error-message"/>
      <Field name="password"
             type="password"
             class="input-password"
             placeholder="password"
             :rules="validatePassword"/>
      <ErrorMessage name="password" class="error-message"/>
      <button class="button"
              :disabled="!meta.valid"
              @click="login">
        login
      </button>
    </Form>
  </div>
</div>
</template>

<script>
import {defineComponent} from 'vue';
import {Form, Field, ErrorMessage} from 'vee-validate';
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
      user
    }
  },
  methods: {
    onSubmit(values) {
      this.user.login = values.login;
      this.user.password = values.password;
    },
    validateLogin(value) {
      // if the field is empty
      if (!value) {
        return 'This field is required';
      }
      // if the field is not a valid email
      const regex = /^[A-Za-z0-9]/i;
      if (!regex.test(value)) {
        return 'login can contain only letters and numbers';
      }
      this.user.login = value;
      // All is good
      return true;
    },
    validatePassword(value) {
      if (!value) {
        return 'This field is required';
      }
      this.user.password = value;
      return true;
    },
    login(){
      console.log("test")
    }
  }
})
</script>

<style lang="scss">
@use "login";
</style>