<template>
  <div id="app" @touchstart="clicked" @mousedown="clicked">
    <div id="nav">
      <b-button variant="light" v-if="this.$store.state.authenticated" to="/" v-on:click.native="logout()" replace>Kirjaudu ulos</b-button>
      <Login />
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
  background: #fafafa;
}

#nav {
  padding: 30px;
  display: flex;
  background: #072556;
}

#nav a {
  color: black;
}

#nav a.router-link-exact-active {
}
#nav {
  display: flex;
}
</style>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
    /* eslint-disable no-alert, no-console */

import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
var x = document.getElementsByTagName("button");
var i;
for (i = 0; i < x.length; i++) {
  x[i].addEventListener("touchstart", (ev) => {
    console.log("vittu touch start")
  })
}
@Component
export default class App extends Vue {
  @Prop() private msg!: string;
  private clicked(ev: Event) {
    console.log("clicked")
    if (ev.target !== null) {
      const targetClass = ((ev.target as any)._prevClass == null ? "" : (ev.target as any)._prevClass)
      console.log(targetClass)
      if (this.$store.state.probableVotingId != "" && (targetClass.indexOf("voteButton") === -1 && targetClass.indexOf("candidateName") === -1) || targetClass.indexOf("disabled") !== -1) {
        console.log("clearing target")
        this.$store.commit("clearProbableVotingTarget")
      } else {
        console.log("not clearing target")
      }
    }
  }
  private async mounted() {
    console.log("here we go")
    if (window.localStorage.getItem("token") !== null) {
      this.$store.dispatch("login", {token: window.localStorage.getItem("token")})
    }
  }
  private logout() {
    this.$store.commit('logout')
    this.$router.push("/")
  }
}

</script>