<template>
  <div>
    <b-container>
      <b-alert :show="loading" variant="info">Loading...</b-alert>
      <b-jumbotron fluid :show="!loading">
        <template slot="header">
          Edit Server: {{serverInfo.title}}
        </template>
        <template slot="lead">
          Running: {{serverInfo.game.toUpperCase()}}
        </template>
        <hr class="my-4">
        <b-container>
          <div v-for="stat in serverData[0].stats" :key="stat.id">
            <h5>{{stat.name.toUpperCase()}}</h5>
            <b-progress :max="stat.max" show-value>
              <b-progress-bar :value="stat.value" :label="stat.value.toFixed(0) + '%'"></b-progress-bar>
            </b-progress>
          </div>
        </b-container>
      </b-jumbotron>
    </b-container>
  </div>
</template>

<style>
  .progress {
    background-color: #00000061
  }
</style>

<script>
import api from '@/api'
export default {
  props: ['id'],
  data () {
    return {
      loading: false,
      serverData: {},
      serverInfo: {}
    }
  },
  async created () {
    this.refreshServerInfo()
  },
  methods: {
    async refreshServerInfo () {
      this.loading = true
      this.serverInfo = await api.getServer(this.id)
      this.serverData = await api.getStats(this.id)
      this.loading = false
    }
  }
}
</script>