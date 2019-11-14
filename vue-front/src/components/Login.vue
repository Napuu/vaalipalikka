<template>
  <div v-if="!this.$store.state.authenticated" class="login">
    <div class="loginDescription">Äänestyskoodi: </div>
    <b-input class="inputfield" v-model="token" type="password" />
    <b-button variant="secondary" v-on:click="login">Kirjaudu</b-button>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';

    /* eslint-disable no-alert, no-console */
@Component
export default class Login extends Vue {
  private token: string = ""
  data() {
    return {
      token: ""
    }
  }
  login() {
    if (this.token !== "") {
      fetch("/api?action=login", {headers: {"Authorization": this.token}}).then(async res => {
        let text = await res.text()
        if (text !== "denied") {
          if (text === "admin") {
            this.$store.commit("login", {role: "admin", token: this.token})
            this.$router.push("admin")
            await this.$store.dispatch("fetchAdminViewableData")
          } else {
            let votings = await fetch("/api?action=voter&a=show", {headers: {"Authorization": this.token}})
            let votingsJson = await votings.json()
            this.$store.commit("setVotings", {votings: votingsJson})
            this.$store.commit("login", {role: "voter", token: this.token})
            this.$router.push("voting")
          }
          window.localStorage.setItem("token", this.token)
        }
        this.token = ""
      })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.loginDescription {
  color: lightgrey;
  font-size: 120%;
  padding-top: 7px;
  padding-right: 10px;
}
.login {
  display: flex;
}
input {
  width: 100px;
}
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
    color: #42b983;
}
    </style>
