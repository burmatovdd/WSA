<template>
  <nav class="nav">
    <ul class="nav__row">
      <li class="nav__item">
        <div class="nav__item-row">
          <RouterLink to="/dashboard" class="nav__item-row-link">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" class="nav__img">
              <g clip-path="url(#clip0_457_1879)">
                <path d="M5 12H3L12 3L21 12H19" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
                <path d="M5 12V19C5 19.5304 5.21071 20.0391 5.58579 20.4142C5.96086 20.7893 6.46957 21 7 21H17C17.5304 21 18.0391 20.7893 18.4142 20.4142C18.7893 20.0391 19 19.5304 19 19V12" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" />
                <path d="M9 21V15C9 14.4696 9.21071 13.9609 9.58579 13.5858C9.96086 13.2107 10.4696 13 11 13H13C13.5304 13 14.0391 13.2107 14.4142 13.5858C14.7893 13.9609 15 14.4696 15 15V21" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" />
              </g>
              <defs>
                <clipPath id="clip0_457_1879">
                  <rect width="24" height="24" fill="white"/>
                </clipPath>
              </defs>
            </svg>
            <p class="nav__text">Главная</p>
          </RouterLink>
        </div>
      </li>
      <li class="nav__item">
        <div class="nav__item-row">
          <RouterLink to="/reports" class="nav__item-row-link">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" class="nav__img">
              <g clip-path="url(#clip0_459_2114)">
                <path d="M9.615 20H7C6.46957 20 5.96086 19.7893 5.58579 19.4142C5.21071 19.0391 5 18.5304 5 18V6C5 5.46957 5.21071 4.96086 5.58579 4.58579C5.96086 4.21071 6.46957 4 7 4H15C15.5304 4 16.0391 4.21071 16.4142 4.58579C16.7893 4.96086 17 5.46957 17 6V14" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
                <path d="M14 19L16 21L20 17" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
                <path d="M9 8H13" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
                <path d="M9 12H11" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
              </g>
              <defs>
                <clipPath id="clip0_459_2114">
                  <rect width="24" height="24" fill="white"/>
                </clipPath>
              </defs>
            </svg>
            <p class="nav__text">Отчеты</p>
          </RouterLink>
        </div>
      </li>
      <li class="nav__item">
        <div class="nav__item-row">
          <RouterLink to="/analytics" class="nav__item-row-link">
            <svg width="24" height="25" viewBox="0 0 24 25" fill="none" xmlns="http://www.w3.org/2000/svg" class="nav__img">
              <g clip-path="url(#clip0_459_2356)">
                <path d="M7 15.5H5C3.89543 15.5 3 16.3954 3 17.5V19.5C3 20.6046 3.89543 21.5 5 21.5H7C8.10457 21.5 9 20.6046 9 19.5V17.5C9 16.3954 8.10457 15.5 7 15.5Z" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
                <path d="M19 15.5H17C15.8954 15.5 15 16.3954 15 17.5V19.5C15 20.6046 15.8954 21.5 17 21.5H19C20.1046 21.5 21 20.6046 21 19.5V17.5C21 16.3954 20.1046 15.5 19 15.5Z" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
                <path d="M13 3.5H11C9.89543 3.5 9 4.39543 9 5.5V7.5C9 8.60457 9.89543 9.5 11 9.5H13C14.1046 9.5 15 8.60457 15 7.5V5.5C15 4.39543 14.1046 3.5 13 3.5Z" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
                <path d="M6 15.5V14.5C6 13.9696 6.21071 13.4609 6.58579 13.0858C6.96086 12.7107 7.46957 12.5 8 12.5H16C16.5304 12.5 17.0391 12.7107 17.4142 13.0858C17.7893 13.4609 18 13.9696 18 14.5V15.5" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
                <path d="M12 9.5V12.5" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
              </g>
              <defs>
                <clipPath id="clip0_459_2356">
                  <rect width="24" height="24" fill="white" transform="translate(0 0.5)"/>
                </clipPath>
              </defs>
            </svg>
            <p class="nav__text">Аналитика</p>
          </RouterLink>
        </div>
      </li>
    </ul>
  </nav>
</template>

<script>
import { defineComponent } from 'vue';
import {RouterLink} from 'vue-router'
export default defineComponent({
  name: "custom-nav",
  data: function () {
    return {};
  },
});
</script>

<style lang="scss">
@use 'nav';
</style>