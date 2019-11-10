<template>
  <div class="voting">
    <div v-bind:key="voting.Id" v-for="voting in this.$store.state.votings">
      <OneVoting :voting="voting"/>
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
