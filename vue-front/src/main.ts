import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import BootstrapVue from 'bootstrap-vue' 
Vue.config.productionTip = false

// these really shouldn't be here but oh well...
import Login from '@/components/Login.vue'
Vue.component('Login', Login);
import TokenThingy from '@/components/TokenThingy.vue'
Vue.component('TokenThingy', TokenThingy);
Vue.use(BootstrapVue)

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
