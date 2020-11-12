<template>
  <div class="onevoting">
    <div class="col1">
      <div id="title" class="button">{{this.voting["Name"]}}</div>
      <div class="candidateListing">
        <div v-for="candidate in this.voting['Candidates']" v-bind:key="candidate.Id">
          <div v-if="voting['Open'] && !voting['Ended'] && voting['VotesLeft'] > 0">
            <b-button block squared variant="primary" class="button candidateName" v-on:click="() => {initialClick(candidate.Id)}">{{candidate.Name}}</b-button>
            <b-button block squared variant="success" class="button voteButton" v-on:click="() => {voteClick(candidate.Id)}" v-if="state.probableCandidateId == candidate['Id'] && state.probableVotingId == voting['Id']">Äänestä</b-button>
          </div>
          <div v-else>
            <div v-if="candidate.Voted">
              <div class="btn button candidateName btn-success disabled btn-block rounded-0" v-on:click.native="vittu" block squared disabled variant="success">{{candidate.Name}}</div>
            </div>
            <div v-else>
              <div class="btn button candidateName btn-secondary disabled btn-block rounded-0" v-on:click.native="vittu" block squared disabled variant="primary">{{candidate.Name}}</div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div disabled variant="outline-primary" class="status disabled rounded-0 btn btn-outline-primary">
      {{(() => {
        if (this.voting["Open"] && !this.voting["Ended"] && this.voting["VotesLeft"] > 0) {
          return "Äänestys on  auki!"
        } else if (this.voting["Ended"]) {
          return "Äänestys on päättynyt."
        } else if (!this.voting["Open"] && !this.voting["Ended"]) {
          return "Äänestys ei ole vielä alkanut"
        } else {
          return "Olet jo äänestänyt tässä äänestyksessä!"
        }
      })()}}
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
    /* eslint-disable no-alert, no-console */
@Component
export default class OneVoting extends Vue {
  @Prop() private voting!: Voting;
  private state = this.$store.state;
  private initialClick(id: string) {
    if (this.voting["Open"] && !this.voting["Ended"] && this.voting["VotesLeft"] > 0) this.$store.commit("setProbableVotingTarget", {votingid: this.voting["Id"], candidateid: id})
  }
  private voteClick(candidateId: string) {
    fetch("/vaalit_api?action=voter&a=vote", {
      method: "POST",
      body: JSON.stringify({
        VotingId: this.voting.Id,
        Id: (new Date().getTime()).toString(),
        CandidateId: candidateId,
        Token: this.$store.state.token
      }),
      headers: {"Authorization": this.$store.state.token}
    }).then(async res => {
      let votings = await fetch("/vaalit_api?action=voter&a=show", {headers: {"Authorization": this.$store.state.token}})
      let votingsJson = await votings.json()
      this.$store.commit("setNonAdminVotings", {votings: votingsJson})
      this.$store.commit("clearProbableVotingTarget")
    })
    //const textAns = await ans.text()
  }
  data() {
    return {

    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.candidateListing {
}
.col1 {
  flex-grow: 3;
}
.button {
  margin-bottom: 3px;
}
#title {
  background: grey;
  color: white;
  /* shifting a bit right ":D"
  margin-left: 3px;
  margin-right: -3px;
  */
}
.onevoting {
  padding: 10px;
  border-bottom: 2px solid black;
  display: flex;
}
.voteButton {
  margin-top: -3px;
  margin-bottom: 10px;
}
.status {
  margin-left: 10px;
  margin-bottom: 2px;
  flex-basis: 40%;
  min-width: 40%;
  max-width: 40%;
}
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: flex;
}
a {
    color: #42b983;
}
    </style>
