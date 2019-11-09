import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

Vue.config.productionTip = false

// these really shouldn't be here but oh well...
import Login from '@/components/Login.vue'
Vue.component('Login', Login);
import TokenThingy from '@/components/TokenThingy.vue'
Vue.component('TokenThingy', TokenThingy);

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
