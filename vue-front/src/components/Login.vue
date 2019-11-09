<template>
  <div v-if="!this.$store.state.authenticated" class="login">
    <div class="loginDescription">Äänestyskoodi: </div>
    <input v-model="token" type="password" />
    <button v-on:click="login">Kirjaudu</button>
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
            this.$store.commit("login", {role: "admin"})
            this.$router.push("admin")
          } else {
            this.$store.commit("login", {role: "voter", token: this.token})
            this.$router.push("voting")
          }
        }
        this.token = ""
      })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.login {
  display: flex;
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
