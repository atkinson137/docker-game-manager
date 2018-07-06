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
          <b-table hover striped :items="serverList" :fields="serverListFields">
            <template slot="edit" slot-scope="data">
              <a :href="data.value">Edit</a>
            </template>
          </b-table>
        </b-col>
      </b-row>
    </b-container>
  </div>
</template>

<script>
import serverResource from '@/servers/serverResource'
export default {
  data () {
    return {
      loading: false,
      servers: [{
        id: 1,
        title: 'Test',
        gameId: 1
      }],
      games: [{
        id: 1,
        name: 'ark',
        dockerImage: 'hello-world'
      }],
      model: {},
      serverListFields: [
        { key: 'id', sortable: true },
        { key: 'serverName', sortable: true },
        { key: 'game', sortable: true },
        'edit'
      ]
    }
  },
  async created () {
    this.rerfreshServers()
  },
  methods: {
    async rerfreshServers () {
      this.loading = true
      this.servers = await serverResource.getServers()
      this.games = await serverResource.getGames()
      this.loading = false
    },
    log (item) {
      console.log(item)
    }
  },
  computed: {
    serverList () {
      var self = this
      var serverRows = []
      self.servers.forEach(function (e) {
        if (self.games !== undefined) {
          self.games.forEach(function (game) {
            if (game.id === e.gameId) {
              serverRows.push({ id: e.id, serverName: e.title, game: game.name.toUpperCase(), edit: '/servers/edit/' + e.id })
            }
          })
        }
      })
      return serverRows
    }
  }
}
</script>