<template>
<div class="tile">
  <div class="tile__content">
    <slot>
      <pre>Fallback content</pre>
    </slot>
  </div>
</div>
</template>

<script>
import {defineComponent} from 'vue';
export default defineComponent({
  name: "tile"
})
</script>

<style lang="scss">
@use 'tile';
</style>