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

export default {
  name: 'voting',
  components: {
    OneVoting
  }
}
</script>
