<template>
  <div
    class='dialog'
    :class="{ 'dialog--open': this.$props.open }"
    :aria-hidden="this.$props.open === true ? 'false' : 'true'"
    role="dialog"
  >
    <div class='dialog__backdrop' @click="$emit('closeModalDialog')" />
    <div class='dialog__container'>
      <button class="dialog__close" @click="$emit('closeModalDialog')">Close button</button>
      <slot>
        <pre>Fallback content</pre>
      </slot>
    </div>
  </div>
</template>

<script>
import { defineComponent } from 'vue';
export default defineComponent({
  name: "dialogcomp",
  props: {
    open: {
      type: Boolean,
      default: false
    }
  },
  mounted() {
    this.onEsc = document.addEventListener('keyup', e => {
      if (e.key == "Escape") {
        this.$emit('closeModalDialog');
      }
    });
  },
  unmounted() {
    document.removeEventListener('keyup', this.onEsc);
  },
  emits: ['closeModalDialog']
});
</script>

<style lang="scss">
@use 'dialog';

</style>
