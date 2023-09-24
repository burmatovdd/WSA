<template>
  <div class="data__right">
    <div class="data__tables reports__tables">
      <div class="data__name">Ресурс</div>
    </div>
    <div class="data__list">
      <ul>
        <li v-for="item in noResolve">
          <div class="data__content">
            <p class="data__content-resource">{{ item.resource }}</p>
            <p class="data__content-date">{{ item.date }}</p>
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
    noResolve: Object,
    newWaf: Object,
  },
})
</script>

<style scoped>

</style>