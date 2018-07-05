import serverService from '@/servers/serverService'

/*

This is for manipulating data from the api

*/

export default {
  getServers () {
    return serverService.getServers()
  },
  getServer (id) {
    return serverService.getServer(id)
  },
  getStats (id) {
    return serverService.getStats(id)
  },
  getGames () {
    return serverService.getGames()
  },
  setServerInfo (id, newInfo) {
    return serverService.setServerInfo(id, newInfo)
  },
  postNewServer (newInfo) {
    return serverService.postNewServer(newInfo)
  }
}
