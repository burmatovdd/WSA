<template>
<Block>
  <div class="checker">
    <div class="checker__title">Получить информацию о ресурсе</div>
    <Form class="checker__form" name="form" v-slot="{ meta }">
      <Field name="resource-name"
             type="text"
             class="checker__form-input checker__form-input-login"
             placeholder="Ресурс"
             v-model="resource"
             :rules="validateResource"/>
      <button class="checker__form-button" :disabled="!meta.valid" @click="checkResource">Проверить ресурс</button>
    </Form>
  </div>
</Block>
</template>

<script>
import {defineComponent} from 'vue';
import * as httpClient from "../../../../httpClient.js";
import {Form, Field} from 'vee-validate';
import Block from "../../../Tile/tile.vue";

export default defineComponent({
  name: "checker",
  components:{
    Form,
    Field,
    Block
  },
  data: function (){
    return{
      resource: null,
      resourceInfo: {
        found: false,
        resName: null,
        ip: null,
        status: false,
        waf: false,
        ssl: {
          ok: true,
          commonName: null,
          issuer: null,
          date: null,
        },
        email: null,
        fio: null,
        owner: null,
      },
    }
  },
  methods: {
    validateResource: function(value) {
      this.isActive = false
      // if the field is empty
      if (!value) {
        return null;
      }
      // All is good
      return true;
    },
    checkResource: function (event) {
      event.preventDefault()
      let sendUrl = "http://localhost:8080/api/check-resource";

      return httpClient.Post(sendUrl,this.$data.resource).then(response =>{
        let resp = JSON.parse(response.data.body)
        if (resp.IP === "" || resp.IP === "----------"){
          this.resourceInfo.ip = "undefined"
        }else{
          this.resourceInfo.ip = resp.IP
        }

        if (resp.SSL.date_cert === "--------------------" || resp.SSL.date_cert === "" ){
          this.resourceInfo.ssl.ok = false
        }
        else{
          this.resourceInfo.date = resp.DateEnd
        }
        if (resp.Email === ""){
          this.resourceInfo.email = "нет"
          this.resourceInfo.fio = "нет"
        }
        if (resp.Owner === ""){
          this.resourceInfo.owner = "нет"
        }
        else{
          this.resourceInfo.email = resp.Email
          this.resourceInfo.fio = resp.FIO
        }
        this.resourceInfo.resName = resp.URL
        this.resourceInfo.status = resp.Status
        this.resourceInfo.waf = resp.WAF
        this.resourceInfo.ssl.commonName = resp.SSL.common_name
        this.resourceInfo.ssl.issuer = resp.SSL.issuer
        this.resourceInfo.ssl.date = resp.SSL.date_cert
        this.resourceInfo.found = true
        this.$emit('checkerResult',this.$data.resourceInfo)

      }).catch(error => {
        if (error.response.data.code === 500){
          this.resourceInfo.found = false
          this.resourceInfo.resName = this.$data.resource
          this.$emit('checkerResult',this.resourceInfo)
        }
      })
    },
  }
})
</script>

<style lang="scss">
@use 'checker';
</style>