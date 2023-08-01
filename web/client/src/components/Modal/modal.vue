<template>
  <div
    class='dialog'
    :class="{ 'dialog--open': this.$props.open }"
    :aria-hidden="this.$props.open === true ? 'false' : 'true'"
    role="dialog"
  >
    <div class='dialog__backdrop' @click="$emit('closeModalDialog')" />
    <div class='dialog__container'>
      <button class="dialog__close" @click="$emit('closeModalDialog')">Close modal dialog</button>
      <slot>
        <pre>Fallback content</pre>
      </slot>
    </div>
  </div>
</template>

<script>
import { defineComponent } from 'vue';
export default defineComponent({
  name: "modal.vue",
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
.dialog {
  display: none;
  position: fixed;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  flex-direction: column;
  padding: 1rem;
  align-items: center;
  justify-content: center;

  &__backdrop {
    position: absolute;
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
    background-color: rgba(0, 0, 0, .618);
  }

  &__container {
    position: relative;
    width: 100%;
    max-width: 600px;
    padding: 2.5rem;
    background-color: #1A2028;
  }

  &__close {
    min-height: 1.375rem;
    min-width: 1.375rem;
    font-size: 0;
    line-height: 0;
    color: rgba(0, 0, 0, 0);
    position: absolute;
    right: 1rem;
    top: 1rem;
  }

  &--open {
    display: flex;
  }
}
</style>
