<template>
  <div class="data data__left">
    <div class="data__tables reports__tables">
      <div class="data__name">{{info1.title}}</div>
    </div>
    <div class="data__list">
      <ul>
        <li v-for="item in info1.array">
          <div class="data__content">
            <p class="data__content-resource">{{ item.resource }}</p>
          </div>
        </li>
      </ul>
    </div>
  </div>
  <div class="data data__right">
    <div class="data__tables reports__tables">
      <div class="data__name">{{info2.title}}</div>
    </div>
    <div class="data__list">
      <ul>
        <li v-for="item in info2.array">
          <div class="data__content">
            <p class="data__content-resource">{{ item.resource }}</p>
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
import {defineComponent} from 'vue';

export default defineComponent( {
  name: "list",
  props: {
    info1: Object,
    info2: Object,
  },
})
</script>

<style scoped>

</style>