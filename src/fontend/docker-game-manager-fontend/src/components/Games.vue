<template>
  <div>
    <b-container>
      <b-alert variant="success" :show="success.curTime" dismissable @dismissed="this.success.curTime = 0" @dismiss-count-down="successCountDownChanged">
        <p>{{success.message}}</p>
        <b-progress variant="success" :max="success.maxTime" :value="success.curTime" height="4px">
        </b-progress>
      </b-alert>
      <b-jumbotron fluid>
        <template slot="header">
          Game Servers
        </template>
        <template slot="lead">
          Edit your available game servers.
        </template>
      </b-jumbotron>
      <div class="outline">
        <b-btn v-b-toggle.newGameCollapse variant="primary">Add New Game</b-btn>
        <b-collapse id="newGameCollapse" class="mt-2">
          <b-form @submit="addNewGame">
            <b-form-group id="gameNameGroup" label="Game Name" label-for="gameNameInput">
              <b-form-input id="gameNameInput" type="text" v-model="gameNewInfo.name" placeholder="New Game Name" required></b-form-input>
            </b-form-group>
            <b-form-group id="dockerImageGroup" label="Docker Image Name" label-for="dockerImageInput">
              <b-form-input id="dockerImageInput" type="text" v-model="gameNewInfo.dockerImage" placeholder="Docker Image Name" required></b-form-input>
            </b-form-group>
            <b-button type="submit" variant="primary">Add</b-button>
            <b-button type="reset" variant="danger">Reset</b-button>
          </b-form>
        </b-collapse>
      </div>
      <div class="outline">
        <b-table hover striped :items="gameList" :fields="gameListFields" style="background-color: white;">
          <template slot="edit" slot-scope="data">
            <a :href="data.value">Edit</a>
          </template>
        </b-table>
      </div>
    </b-container>
  </div>
</template>

<script>
import serverResource from '@/servers/serverResource'
export default {
  data () {
    return {
      gameListFields: [
        { key: 'id', sortable: true },
        { key: 'gameName', sortable: true },
        { key: 'dockerImage', sortable: true },
        'edit'
      ],
      gameNewInfo: { name: '', dockerImage: '' },
      success: { show: false, message: '', maxTime: 10, curTime: 0, doGoTo: false, goToLoc: '' },
      games: []
    }
  },
  async created () {
    this.refreshGames()
  },
  methods: {
    successCountDownChanged (countDown) {
      this.success.curTime = countDown
    },
    async addNewGame (event) {
      var self = this
      event.preventDefault()
      // TODO: loading logic
      await serverResource.addNewGame(self.gameNewInfo)
      self.success.message = 'New game added!'
      self.success.curTime = self.success.maxTime

      // reset form data
      self.gameNewInfo.name = ''
      self.gameNewInfo.dockerImage = ''
    },
    async refreshGames () {
      this.games = await serverResource.getGames()
    }
  },
  computed: {
    gameList () {
      var self = this
      var gameRows = []
      self.games.forEach(function (e) {
        gameRows.push({ id: e.id, gameName: e.name, dockerImage: e.dockerImage, edit: '/games/edit/' + e.id })
      })
      return gameRows
    }
  }
}
</script>