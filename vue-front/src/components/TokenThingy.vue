<template>
  <div class="token">
    <div v-if="this.$store.state.token !== ''">
      <b-button variant="primary" v-on:click="toggle">
        <div v-if="this.$store.state.tokenIsVisible">{{this.$store.state.token}}</div>
        <div v-else>{{this.$store.state.token.split("").map(c => "*").join("")}}</div>
      </b-button>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';

    /* eslint-disable no-alert, no-console */
@Component
export default class TokenThingy extends Vue {
  private toggle(e: Event) {
    this.$store.commit("toggleTokenVisibility")
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
