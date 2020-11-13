<template>
  <div class="voting">
    <div style="padding: 10px; font-size: large" v-if="this.$store.state.votings.filter((v) => v.Visible).length == 0">Ei meneillään olevia äänestyksiä</div>
    <div v-bind:key="voting.Id" v-for="voting in this.$store.state.votings">
      <OneVoting v-if="voting.Visible === 1" :voting="voting"/>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import OneVoting from '@/components/OneVoting.vue'
let interval;
export default {
  name: 'voting',
  mounted() {
    let store = this.$store;
    async function f() {
      const r = await fetch("/vaalit_api?action=voter&a=show", {headers: {"Authorization": store.state.token}});
      const votings = await r.json();
      store.commit("setNonAdminVotings", {votings})
    }
    interval = setInterval(f, 5000);
  },
  beforeDestroy() {
    clearInterval(interval);
  },
  components: {
    OneVoting
  }
}
</script>
