import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    authenticated: false,
    tokenIsVisible: false,
    token: "",
    role: "",
  },
  mutations: {
    login(state, {role , token}) {
      state.authenticated = true
      state.role = role
      state.token = token
    },
    logout(state) {
      state.authenticated = false
      state.role = ""
      state.token = ""
    },
    toggleTokenVisibility(state) {
      state.tokenIsVisible = !state.tokenIsVisible
    }
  },
  actions: {
  },
  modules: {
  }
})
