<template>
  <div>
    <b-alert :show="loading" variant="info">Loading...</b-alert>
    <template v-if="!loading">
      <b-container>
        <b-jumbotron fluid>
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
        <div class="outline">
          <div>
            <b-button-group>
              <b-btn variant="success">Start</b-btn>
              <b-btn variant="warning">Suspend</b-btn>
            </b-button-group>
            <b-btn class="float-right" variant="danger" v-b-modal.removeServerModal>Remove</b-btn>
            <b-modal id="removeServerModal" title="Remove Server">
              <p class="my-6">Are you sure you want to remove this server? You cannot undo this action, all non backed up data will be lost.</p>
              <div slot="modal-footer">
                <b-btn size="sm" class="float-right" variant="primary">No</b-btn>
                <b-btn size="sm" class="float-right" variant="danger">Yes</b-btn>
              </div>
            </b-modal>
          </div>
          <div class="outline">
            <b-btn v-b-toggle.serverInfoCollapse variant="primary">Edit Server Info</b-btn>
            <b-collapse id="serverInfoCollapse" class="mt-2">
              <b-form>
                <b-form-group id="serverNameGroup" label="Server Name" label-for="serverNameInput" description="The human-readable name">
                  <b-form-input id="serverNameInput" type="text" v-model="serverNewInfo.name" placeholder="New Server Name"></b-form-input>
                </b-form-group>
              </b-form>
            </b-collapse>
          </div>
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
      serverData: [ { 'id': 1, 'serverId': 1, 'stats': [ { 'id': 1, 'name': 'cpu', 'value': 0 }, { 'id': 2, 'name': 'ram', 'value': 0 } ] } ],
      serverInfo: { 'id': 1, 'title': '', 'game': '' },
      serverNewInfo: {
        name: ''
      }
    }
  },
  async created () {
    this.refreshServerInfo()
  },
  methods: {
    async refreshServerInfo () {
      this.loading = true
      this.serverInfo = await serverResource.getServer(this.id)
      this.serverData = await serverResource.getStats(this.id)
      this.loading = false
    }
  }
}
</script>