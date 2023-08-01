<template>
  <nav class="nav">
    <ul class="nav__row">
      <li class="nav__item">
        <div class="nav__item-row">
          <button class="nav__button" @click="toDashboard">
              <img src="../../img/dashboard.svg" alt="dashboard-icon" class="nav__img">
          </button>
         <p class="nav__text">Дашборд</p>
        </div>
      </li>
      <li class="nav__item">
        <div class="nav__item-row">
          <button class="nav__button" @click="toOwners">
            <img src="../../img/owner.svg" alt="dashboard-icon" class="nav__img">
          </button>
          <p class="nav__text">Организации</p>
        </div>
      </li>
      <li class="nav__item">
          <div class="nav__item-row">
              <button class="nav__button" @click="exit">
                  <img src="../../img/exit.svg" alt="dashboard-icon" class="nav__img">
              </button>
              <p class="nav__text">Выход</p>
          </div>
      </li>
    </ul>
  </nav>
</template>

<script>
import { defineComponent } from 'vue';
export default defineComponent({
  name: "custom-nav",
  data: function () {
    return {};
  },
  methods: {
    toDashboard() {
      this.$router.push('/dashboard')
    },
    toOwners() {
      console.log("in work")
    },
    exit() {
      this.$router.push('/')
    },
  }
});
</script>

<style lang="scss">
.nav {
  &__row {
    display: flex;
    flex-direction: column;
    list-style: none;
    padding: 0;

    >*:not(script) {
      padding: 22.5px 8px;
    }

    &:not(:empty) {
      margin: -22.5px -8px;
    }
  }

  &__item {
    display: flex;
    flex-direction: column;
    align-items: center;

    &-row {
      display: flex;
      flex-direction: column;
      align-items: center;

      >*:not(script) {
        padding: 10px;
      }

      &:not(:empty) {
        margin: -10px;
      }
    }
  }

  &__text {
    margin: 0;
  }

  &__button {
    background-color: #1A2028;
    border: none;
    cursor: pointer;
    fill: tomato;

    &-img {
      width: 40px;
    }
  }
}
</style>
