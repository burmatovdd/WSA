<template>
  <table class="resource">
    <thead class="resource__thead">
    <tr class="resource__tr">
      <th class="resource__th resource__th--name">
        Доменное имя
      </th>
      <th class="resource__th resource__th--value">
        {{this.$props.resource.resName}}
      </th>
    </tr>
    <tr class="resource__tr">
      <th class="resource__th resource__th--name">
        IP адрес
      </th>
      <th class="resource__th resource__th--value">
        {{this.$props.resource.ip}}
      </th>
    </tr>
    <tr class="resource__tr">
      <th class="resource__th resource__th--name">
        Статус активности
      </th>
      <th v-if="this.$props.resource.status === true" class="resource__th resource__th--value">
        активен
      </th>
      <th v-else class="resource__th resource__th--value">неактивен</th>
    </tr>
    <tr class="resource__tr">
      <th class="resource__th resource__th--name">
        WAF
      </th>
      <th v-if="this.$props.resource.waf === true" class="resource__th resource__th--value">
        за WAF
      </th>
      <th v-else class="resource__th resource__th--value">нет</th>
    </tr>
    <tr class="resource__tr">
      <th class="resource__th resource__th--name">
        SSL
      </th>
      <th v-if="this.$props.resource.ssl.ok === true" class="resource__th resource__th--value">
        {{this.$props.resource.ssl.commonName}}
      </th>
      <th v-else class="resource__th resource__th--value">нет</th>
    </tr>
    <tr class="resource__tr">
      <th class="resource__th resource__th--name">
        Ответственное лицо
      </th>
      <th class="resource__th resource__th--value">
        {{this.$props.resource.fio}}
      </th>
    </tr>
    <tr class="resource__tr">
      <th class="resource__th resource__th--name">
        Организация
      </th>
      <th class="resource__th resource__th--value">
        {{this.$props.resource.owner}}
      </th>
    </tr>
    </thead>

    <thead class="resource__thead">

    </thead>
    <thead class="resource__thead">

    </thead>
    <thead class="resource__thead">

    </thead>
  </table>
</template>

<script>
import {defineComponent} from "vue";
export default defineComponent({
  name: "resource",
  props: {
    resource: Object
  },
})
</script>

<style lang="scss">
@use 'resource';
</style>