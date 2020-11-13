
<template>
  <div class="adminhome">
    <b-navbar toggleable="lg" type="dark" variant="info">
      <b-navbar-brand @click="move('candidate')" href="#">Ehdokkaat</b-navbar-brand>
      <b-navbar-brand @click="move('voting')" href="#">Äänestykset</b-navbar-brand>
      <b-navbar-brand @click="move('token')" href="#">Äänestysavaimet</b-navbar-brand>
      <b-navbar-brand @click="move('vote')" href="#">Äänet</b-navbar-brand>
    </b-navbar>
    <div v-if="tab == 'voting'" class="voting">
      <button @click="addVoting">lisää äänestys</button>
      <div v-bind:key="voting.Id" v-for="voting in this.$store.state.admin.votings">
        <OneVotingAdmin v-if="voting.Id !== ''" :voting="voting"/>
      </div>
<!-- TODO make admin buttons (modify, save, delete) disabled or something when they should be -->
    </div>
    <div v-else-if="tab == 'candidate'" class="candidate">
      <button @click="addCandidate">lisää ehdokas</button>
      <div v-bind:key="candidate.Id" v-for="candidate in this.$store.state.admin.candidates">
        <OneCandidateAdmin :candidate="candidate"/>
      </div>
    </div>
    <div v-else-if="tab == 'token'" class="token">
      <div v-bind:key="token.Value" v-for="token in this.$store.state.admin.tokens">
        <OneTokenAdmin :token="token"/>
      </div>
    </div>
    <div v-else-if="tab == 'vote'" class="vote">
      <!-- <button @click="addToken">lisää token</button> -->
      <div>aktiivisia avaimia: {{this.$store.state.admin.tokens.reduce((a, b) => a + b.Valid, 0)}}</div>
      <div v-bind:key="voting.Id" v-for="voting in this.$store.state.admin.votings">
        <VotesAdmin :voting="voting"/>
      </div>
    </div>
    <div v-else>
      {{tab}}
    </div>
  </div>
</template>

<script lang="ts">
// @ is an alias to /src

import OneVotingAdmin from '@/components/OneVotingAdmin.vue'
import OneCandidateAdmin from '@/components/OneCandidateAdmin.vue'
import OneAvailabilityAdmin from '@/components/OneAvailabilityAdmin.vue'
import OneTokenAdmin from '@/components/OneTokenAdmin.vue'
import VotesAdmin from '@/components/VotesAdmin.vue'
import store from '../store'
    /* eslint-disable no-alert, no-console */
let interval;
export default {
  methods: {
    move(_tab: string) {
      clearInterval(interval);
      let ff = this;
      switch (_tab) {
        case "candidate":
          store.dispatch("fetchCandidates")
          break
        case "voting":
          store.dispatch("fetchVotings")
          break
        case "availability":
          store.dispatch("fetchAvailabilities")
          break
        case "token":
          store.dispatch("fetchTokens")
          break
        case "vote":
          interval = setInterval(() => {
            ff.store.dispatch("fetchVotes")
            ff.move("refresh");
          }, 5000);
          break
        // I have no idea about vue lifecycles right now so using this hack to remount component
        case "refresh":
          setTimeout(() => {
            ff.move("vote");
          }, 10);
          break
      }
      this.tab = _tab
    },

    addVoting() {
      let newTarget: PureVoting = {Name: "", Id: (new Date().getTime()).toString(), Description: "", Open: 0, Ended: 0, Visible: 0}
      store.commit("addVoting", {voting: newTarget})
      store.commit("setEditableVoting", {voting: newTarget})
    },
    addCandidate() {
      let newTarget: PureCandidate = {Name: "", Id: (new Date().getTime()).toString(), Description: ""}
      store.commit("addCandidate", {candidate: newTarget})
      store.commit("setEditableCandidate", {candidate: newTarget})
    },
    generateTokens() {
      this.$bvModal.msgBoxConfirm(`Generoi uudet avaimet?`, {
        okTitle: 'Ok',
        cancelTitle: 'Peruuta',
      })
      .then((value: boolean) => {
        if (value) {
          this.$store.dispatch("generateTokens")
        }
      })
    },
    sort(array: Array<any>): Array<any> {
      return Array.prototype.slice.call(array).sort((a: any, b: any): number => {
        if (this.getVotingName(a.VotingId) < this.getVotingName(b.VotingId)) return -1
        else return -1
      })
    },
    getVotingName(votingId: string): string {
    let name = this.$store.state.admin.votings.filter((c: PureVoting) => c.Id === votingId)
    return (name.length > 0 ? name[0].Name : "voting id " + votingId + " not found")
  }
  },
  name: 'voting',
  data(): any {
    return {
      tab: "",
      store: store,
    }
  },
  components: {
    OneVotingAdmin,
    OneCandidateAdmin,
    OneTokenAdmin,
    VotesAdmin
  }
}

</script>
