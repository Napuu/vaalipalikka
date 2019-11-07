<template>
  <div class="login">
    <input v-model="token" type="text" />
    <button v-on:click="login">kirjaudu</button>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';

    /* eslint-disable no-alert, no-console */
@Component
export default class HelloWorld extends Vue {
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
        console.log(text)
        if (text !== "denied") this.$store.commit("login")
      })

    }
  }
  private clicked(e: Event) {
    /*eslint no-console: "error"*/
    // eslint-disable-next-line no-console
    console.log("asdf")
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
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
