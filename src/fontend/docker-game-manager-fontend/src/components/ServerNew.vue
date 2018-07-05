<template>
  <div>
    <b-container>
      <b-alert variant="success" :show="success.show" dismissable @dismissed="success.curTime = 0" @dismiss-count-down="countDownChanged">
        <p>{{success.message}}</p>
        <b-progress variant="success" :max="success.maxTime" :value="success.curTime" height="4px">
        </b-progress>
      </b-alert>
      <b-jumbotron fluid>
        <template slot="header">
          Create New Server
        </template>
        <template slot="lead">
          Spin up a new server.
        </template>
      </b-jumbotron>
      <div class="outline">
        <b-form @submit="submitNewServer">
          <b-form-group id="serverNameGroup" label="Server Name" label-for="serverNameInput" description="The human-readable name">
            <b-form-input id="serverNameInput" type="text" v-model="serverNewInfo.name" placeholder="New Server Name" required></b-form-input>
          </b-form-group>
          <b-form-select v-model="serverNewInfo.game" :options="gameOptions" class="mb-3" required />
          <b-button type="submit" variant="primary">Spin Up</b-button>
          <b-button type="reset" variant="danger">Reset</b-button>
        </b-form>
      </div>
    </b-container>
  </div>
</template>

<script>
import serverResource from '@/servers/serverResource'
export default {
  data () {
    return {
      loading: false,
      success: { show: false, message: '', maxTime: 10, curTime: 0 },
      games: [ { 'id': 1, 'name': 'ark', 'dockerImage': 'hello-world' }, { 'id': 2, 'name': 'factorio', 'dockerImage': 'hello-world' } ],
      serverNewInfo: {
        name: '',
        game: null
      }
    }
  },
  async created () {
    this.getGames()
  },
  methods: {
    async getGames () {
      this.loading = true
      this.games = await serverResource.getGames()
      this.loading = false
    },
    async submitNewServer (event) {
      event.preventDefault()
      this.loading = true
      try {
        await serverResource.postNewServer(this.serverNewInfo)
        this.success.show = true
        this.success.message = 'Server creation success'
        this.success.curTime = this.success.maxTime
      } catch (e) {
        console.log(e)
        alert('There was an error, check the log')
      }
      this.loading = false
    },
    countDownChanged (countDown) {
      this.success.curTime = countDown
    }
  },
  computed: {
    gameOptions: function () {
      var options = []
      options.push({ value: null, text: 'Please select an option' })
      this.games.forEach(function (e) {
        options.push({ value: e.id, text: e.name.toUpperCase() })
      })
      return options
    }
  }
}
</script>