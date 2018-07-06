<template>
  <div>
    <b-alert :show="loading" variant="info">Loading...</b-alert>
    <b-alert dismissible @dismissed="error=false" :show="error" variant="danger">{{errorText}}</b-alert>
    <template v-if="!loading">
      <b-container>
        <b-jumbotron fluid>
          <template slot="header">
            Edit Game: {{gameInfo.name}}
          </template>
          <template slot="lead">
            Image: {{gameInfo.dockerImage}}
          </template>
        </b-jumbotron>
        <div class="outline">
          <b-form @submit="setGameInfo">
            <b-form-group id="gameNameGroup" label="Game Name" label-for="gameNameInput">
              <b-form-input id="gameNameInput" type="text" v-model="gameNewInfo.name" placeholder="Change Game Name"></b-form-input>
            </b-form-group>
            <b-form-group id="gameDockerImageGroup" label="DockerImage" label-for="gameDockerImageInput">
              <b-form-input id="gameDockerImageInput" type="text" v-model="gameNewInfo.dockerImage" placeholder="Change the Docker Image"></b-form-input>
            </b-form-group>
            <b-button type="submit" variant="primary">Set Changes</b-button>
            <b-button type="reset" variant="danger">Reset</b-button>
            <div class="float-right">
              <b-btn variant="danger" v-b-modal.removeServerModal>Remove</b-btn>
              <b-modal id="removeServerModal" title="Remove Server">
              <p class="my-6">Are you sure you want to remove this server? You cannot undo this action, all non backed up data will be lost.</p>
              <div slot="modal-footer">
                <b-btn size="sm" class="float-right" variant="primary" @click="hideServerModal">No</b-btn>
                <b-btn size="sm" class="float-right" variant="danger" @click="deleteServer">Yes</b-btn>
              </div>
            </b-modal>
            </div>
          </b-form>
        </div>
      </b-container>
    </template>
  </div>
</template>

<script>
import serverResource from '@/servers/serverResource'
export default {
  props: ['id'],
  data () {
    return {
      loading: false,
      error: false,
      errorText: '',
      gameInfo: { 'id': 1, 'name': '', 'dockerImage': '' },
      gameNewInfo: {}
    }
  },
  async created () {
    this.refreshServerInfo()
  },
  methods: {
    async refreshServerInfo () {
      this.loading = true
      this.gameInfo = await serverResource.getGame(this.id)
      this.gameNewInfo = JSON.parse(JSON.stringify(this.gameInfo))
      this.loading = false
    },
    async setGameInfo () {
      event.preventDefault()
      this.loading = true
      try {
        // send the new server info
        await serverResource.setGameInfo(this.id, this.gameNewInfo)
        // refresh the current page's info to get the updated list from the server
        this.gameInfo = await serverResource.getGame(this.id)
        // TODO: set success alert
      } catch (e) {
        console.log(e)
        this.error = true
        this.errorText = e
      }
      this.loading = false
    },
    removeGameModal () {
      this.$root.$emit('bv::hide::modal', 'removeGameModal')
    },
    async deleteGame () {
      await serverResource.deleteServer(this.id)
      this.$root.$emit('bv::hide::modal', 'removeGameModal')
      this.$router.push({name: 'Games'})
    }
  },
  computed: {
  }
}
</script>