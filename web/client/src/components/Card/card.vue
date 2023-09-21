<template>
  <div class="card">
    <div class="card__content">
      <slot>
        <pre>Fallback content</pre>
      </slot>
    </div>
  </div>
</template>

<script>
import {defineComponent} from 'vue';
export default defineComponent({
  name: "card"
})
</script>

<style lang="scss">
@use 'card';
</style>