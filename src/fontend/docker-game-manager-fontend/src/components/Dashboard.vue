<template>
  <div class="container-fluid mt-4">
    <h1 class="h1">Servers Dashboard</h1>
    <b-jumbotron fluid>
      <template slot="lead">
        Coming soon: System Resource Overview
      </template>
    </b-jumbotron>
    <b-jumbotron fluid>
      <template slot="lead">
        Coming soon: Available Deployment Targets
      </template>
    </b-jumbotron>
    <b-container>
      <b-alert :show="loading" variant="info">Loading...</b-alert>
      <b-row>
          <b-col>
              <table class="table table-striped">
                  <thead>
                      <tr>
                          <th>ID</th>
                          <th>Server Name</th>
                          <th>Game</th>
                          <th>Edit</th>
                      </tr>
                  </thead>
                  <tbody>
                      <tr v-for="server in servers" :key="server.id">
                        <td>{{server.id}}</td>
                        <td>{{server.title}}</td>
                        <td>{{server.game}}</td>
                        <td><b-link :href="'/servers/edit/' + server.id">Edit</b-link></td>
                      </tr>
                  </tbody>
              </table>
          </b-col>
      </b-row>
    </b-container>
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