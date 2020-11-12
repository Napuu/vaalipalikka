<template>
  <div class="voting">
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
