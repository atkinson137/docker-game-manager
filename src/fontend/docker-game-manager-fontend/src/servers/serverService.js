// import Vue from 'vue'
import axios from 'axios'

const client = axios.create({
  baseURL: 'http://localhost:3000/',
  json: true
})

async function execute (method, resource, data) {
  // inject the accessToken for each request
  // let accessToken = await Vue.prototype.$auth.getAccessToken()
  return client({
    method,
    url: resource,
    data
    // headers: {
    //  Authorization: `Bearer ${accessToken}`
    // }
  }).then(req => {
    return req.data
  })
}

export default {
  getServers () {
    return execute('get', '/servers')
  },
  getServer (id) {
    return execute('get', '/servers/' + id)
  },
  getStats (id) {
    return execute('get', '/servers/' + id + '/stats')
  },
  getGames () {
    return execute('get', '/games')
  },
  setServerInfo (id, serverInfo) {
    return execute('put', '/servers/' + id, serverInfo)
  },
  postNewServer (newServerInfo) {
    return execute('post', '/servers', newServerInfo)
  }
}
