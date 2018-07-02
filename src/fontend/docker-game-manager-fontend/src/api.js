// import Vue from 'vue'
import axios from 'axios'

const client = axios.create({
  baseURL: 'http://localhost:3000/',
  json: true
})

export default {
  async execute (method, resource, data) {
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
  },
  getServers () {
    return this.execute('get', '/servers')
  },
  getServer (id) {
    return this.execute('get', '/servers/' + id)
  },
  getStats (id) {
    return this.execute('get', '/servers/' + id + '/stats')
  }
}
