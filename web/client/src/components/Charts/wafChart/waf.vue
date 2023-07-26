<template>
  <Doughnut
      v-if="loaded"
      :options="chartOptions"
      :data="chartData"
  />
  <div class="load" v-else>loading...</div>
</template>

<script lang="ts">
import { Chart as ChartJS, ArcElement, Tooltip, Legend, Chart } from 'chart.js'
import { Doughnut } from 'vue-chartjs'
import { getData } from "./PieConfig.js";
ChartJS.register(ArcElement, Tooltip, Legend)

export default {
  name: 'App',
  components: {
    Doughnut
  },
  data: () => ({
    arr: [],
    loaded: false,
    chartData: {
      labels: ['Ok', 'No Ok'],
      datasets: [
        {
          backgroundColor: ['#294486', '#DE3163'],
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
            padding: 8
          }
        }
      }
    },
  }),

  async mounted() {
    this.loaded = false
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
}
</script>
<style lang="scss">
.load {
  width: 70px;
  margin: auto;
  margin-top: 40px;
}
</style>
