<template>
  <div id="app">
    <div id="nav">
      <Login />
      <router-link v-if="this.$store.state.authenticated" to="/" v-on:click.native="logout()" replace>Logout</router-link>
      <TokenThingy />
    </div>
    <router-view/>
  </div>
</template>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}
#nav {
  display: flex;
}
</style>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
    /* eslint-disable no-alert, no-console */
import Login from '@/components/Login.vue'

@Component
export default class App extends Vue {
  @Prop() private msg!: string;
  private logout() {
    this.$store.commit('logout')
    this.$router.push("/")
  }
}

</script>