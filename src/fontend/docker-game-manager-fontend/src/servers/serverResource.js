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
  },
  deleteServer (id) {
    return serverService.deleteServer(id)
  },
  addNewGame (newInfo) {
    return serverService.postNewGame(newInfo)
  },
  setGameInfo (id, newInfo) {
    return serverService.setGameInfo(id, newInfo)
  },
  getGame (id) {
    return serverService.getGame(id)
  },
  deleteGame (id) {
    return serverService.deleteGame(id)
  }
}
