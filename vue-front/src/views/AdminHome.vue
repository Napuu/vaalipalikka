
<template>
  <div class="adminhome">
    <b-navbar toggleable="lg" type="dark" variant="info">
      <b-navbar-brand @click="move('voting')" href="#">Äänestykset</b-navbar-brand>
      <b-navbar-brand @click="move('notvoting')" href="#">muuta</b-navbar-brand>
    </b-navbar>
    <div v-if="state.admin.activeTab == 'voting'" class="voting">
      <div v-bind:key="voting.Id" v-for="voting in this.$store.state.admin.votings">
        <OneVotingAdmin :voting="voting"/>
      </div>
    </div>
    <div v-else>
      asdfasdfsd
    </div>
  </div>
</template>

<script>
// @ is an alias to /src

import OneVotingAdmin from '@/components/OneVotingAdmin.vue'
    /* eslint-disable no-alert, no-console */
export default {
  methods: {
    move: (tab) => {
      console.log(tab)
    }
  },
  name: 'voting',
  data: () => {
    return {
      state: this.$store.state
    }
  },
  components: {
    OneVotingAdmin
  }
}

</script>