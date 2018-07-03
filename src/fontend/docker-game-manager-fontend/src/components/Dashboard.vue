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
      <b-row>
        <b-btn :size="'sm'" variant="primary" href="/servers/new">Create New Server</b-btn>
      </b-row>
      <b-row>
        <b-col>
          <b-alert :show="loading" variant="info">Loading...</b-alert>
          <table class="table table-striped" v-if="!loading">
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

<style>
  div.row {
    padding-top: 5px;
    padding-bottom: 5px;
  }
</style>

<script>
import serverResource from '@/servers/serverResource'
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
      this.servers = await serverResource.getServers()
      this.loading = false
    }
  }
}
</script>