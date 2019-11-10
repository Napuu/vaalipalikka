
<template>
  <div class="adminhome">
      asdf
    <div class="voting">
      <div v-bind:key="voting.Id" v-for="voting in this.$store.state.admin.votings">
        <OneVotingAdmin :voting="voting"/>
      </div>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src

import OneVotingAdmin from '@/components/OneVotingAdmin.vue'

export default {
  name: 'voting',
  components: {
    OneVotingAdmin
  }
}

</script>