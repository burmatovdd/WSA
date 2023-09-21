<template>
  <Doughnut
      v-if="loaded"
      :options="chartOptions"
      :data="chartData"
      :style="Styles"
  />
  <div class="load" v-else>loading...</div>
</template>

<script lang="ts">
import { Chart as ChartJS, ArcElement, Tooltip, Legend, Chart } from 'chart.js'
import { getData } from "./PieConfig.js";
import { Doughnut } from 'vue-chartjs'
import {defineComponent} from "vue";
ChartJS.register(ArcElement, Tooltip, Legend)
export default defineComponent({
  name: "Chart",
  components: {
    Doughnut
  },
  computed:{
    Styles() {
      return{
        width: `${13}rem`,
        height: `${13}rem`,
      }
    }
  },
  data: function (){
    return {
      arr: [],
      loaded: false,
      chartData: {
        labels: ['за WAF', 'без WAF'],
        datasets: [
          {
            backgroundColor: ['#35D073', '#4A69FF'],
            data: null,
            borderWidth: 0,
          }
        ]
      },
      chartOptions: {
        responsive: true,
        maintainAspectRatio: false,
        plugins: {
          legend: {
            position: 'bottom',
            labels: {
              usePointStyle: true,
              pointStyle: 'circle',
              padding: 24,
            }
          }
        }
      },
    }
  },
  async mounted() {
    try {
      await getData().then(response => {
        this.arr = [response.withWaf, response.noWaf]
        this.chartData.datasets[0].data = this.arr
        this.loaded = true
      })
    } catch (e) {
      console.error(e)
    }
  }
})
</script>

<style lang="scss">
.load {
  width: 70px;
  margin: auto;
  margin-top: 40px;
}
</style>