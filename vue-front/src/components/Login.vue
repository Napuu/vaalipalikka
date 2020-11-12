<template>
  <div>
    <div v-if="!this.$store.state.authenticated" class="login">
      <div class="loginDescription">Äänestyskoodi</div>
      <b-input class="inputfield" v-model="token" type="password" />
      <b-button variant="secondary" v-on:click="login">Kirjaudu</b-button>
    </div>
    <div id="snackbar">Aktivoimaton tai väärä äänestyskoodi</div>
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
    console.log("token should be here " + this.token)
    if (this.token !== "") {
      fetch("/vaalit_api?action=login", {headers: {"Authorization": this.token}}).then(async res => {
        let text = await res.text()
        if (text !== "denied") {
          if (text === "admin") {
            this.$store.commit("login", {role: "admin", token: this.token})
            this.$router.push("admin")
            this.$store.dispatch("fetchAdminViewableData")
            setInterval(() => {
              this.$store.dispatch("fetchAdminViewableData")
            }, 10000);
          } else {
            let votings = await fetch("/vaalit_api?action=voter&a=show", {headers: {"Authorization": this.token}})
            let votingsJson = await votings.json()
            this.$store.commit("setNonAdminVotings", {votings: votingsJson})
            this.$store.commit("login", {role: "voter", token: this.token})
            this.$router.push("voting");
            const _token = this.token; // can't access this inside interval :-)
            setInterval(() => {
              fetch("/vaalit_api?action=voter&a=show", {headers: {"Authorization": _token}})
                .then((res) => res.json())
		.then((votings) => {
                  this.$store.commit("setNonAdminVotings", {votings})
		});
            }, 10000);
          }
          window.localStorage.setItem("token", this.token)
        } else {
          var x = document.getElementById("snackbar");
          x.className = "show";
          setTimeout(function(){ x.className = x.className.replace("show", ""); }, 3000);
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
#snackbar {
  visibility: hidden;
  min-width: 250px;
  background-color: #333;
  color: #fff;
  text-align: center;
  border-radius: 2px;
  padding: 16px;
  position: fixed;
  z-index: 1;
  left: 50%;
  transform: translateX(-50%);
  bottom: 30px;
  font-size: 17px;
}

#snackbar.show {
  visibility: visible;
  -webkit-animation: fadein 0.5s, fadeout 0.5s 2.5s;
  animation: fadein 0.5s, fadeout 0.5s 2.5s;
}

@-webkit-keyframes fadein {
  from {bottom: 0; opacity: 0;} 
  to {bottom: 30px; opacity: 1;}
}

@keyframes fadein {
  from {bottom: 0; opacity: 0;}
  to {bottom: 30px; opacity: 1;}
}

@-webkit-keyframes fadeout {
  from {bottom: 30px; opacity: 1;} 
  to {bottom: 0; opacity: 0;}
}

@keyframes fadeout {
  from {bottom: 30px; opacity: 1;}
  to {bottom: 0; opacity: 0;}
}
</style>
