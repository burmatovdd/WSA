<template>
  <Doughnut v-if="loaded"
            :options="chartOptions"
            :data="chartData"/>
</template>

<script lang="ts">
import {Chart as ChartJS, ArcElement, Tooltip, Legend, Chart} from 'chart.js'
import { Doughnut } from 'vue-chartjs'
import {getData} from "./PieConfig.js";
ChartJS.register(ArcElement, Tooltip, Legend)

export default {
  name: 'App',
  components: {
    Doughnut
  },
  data: () => ({
    arr: [],
    loaded: false,
    chartData: null,
    chartOptions: {
      responsive: true,
      maintainAspectRatio: false,
      plugins: {
        legend: {
          position: 'bottom',
          labels: {
            usePointStyle: true,
            pointStyle: 'circle',
            padding: 50
          }
        }
      }
    },
  }),
  async mounted() {
    this.loaded = false
    try {
      await getData().then(response => {
        this.arr = [response.okCertificates, response.noOkCertificates]
      })
      this.chartData = {
        labels: ['Ok', 'No Ok'],
        datasets: [
          {
            backgroundColor: ['#294486', '#DE3163'],
            data: this.arr,
            borderWidth: 0,
          }
        ]
      }

      this.loaded = true
    } catch (e) {
      console.error(e)
    }
  }
}
</script>
