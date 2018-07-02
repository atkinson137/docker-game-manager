<template>
  <div class="container-fluid mt-4">
    <h1 class="h1">Posts Manager</h1>
    <b-alert :show="loading" variant="info">Loading...</b-alert>
    <b-row>
        <b-col>
            <table class="table table-striped">
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>Server Name</th>
                        <th>Game</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="server in servers" :key="server.id">
                        <td>{{server.id}}</td>
                        <td>{{server.title}}</td>
                        <td>{{server.game}}</td>
                    </tr>
                </tbody>
            </table>
        </b-col>
    </b-row>
  </div>
</template>

<script>
import api from '@/api'
export default {
  data () {
    return {
      loading: false,
      servers: [],
      model: {}
    }
  },
  async created () {
    this.rerfreshServers()
  },
  methods: {
    async rerfreshServers () {
      this.loading = true
      this.servers = await api.getServers()
      this.loading = false
    }
  }
}
</script>