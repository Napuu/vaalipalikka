<template>
  <div class="onevoting">
    <div class="col1 valueholder">
      <div class="title">Nimi</div>
      <div  class="value">{{voting.Name}}</div>
    </div>
    <div class="col2 candidates valueholder">
      <div class="title">Ehdokkaat</div>
      <div v-bind:key="candidate.Id" v-for="candidate in candidatesHere">
        <div class="col20">
          <div class="col21">
            {{candidate.Name}}
          </div>
          <div class="col22">
            {{(votesByCandidate[candidate.Id] ? votesByCandidate[candidate.Id].length : 0)}}
          </div>
        </div>
      </div>
      <div>
        <div class="col20">
          <div class="col21">
            Yhteensä:
          </div>
          <div class="col22">
            {{Object.values(votesByCandidate).reduce((a, b) => a + b.length, 0)}}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import 'vue-multiselect/dist/vue-multiselect.min.css'
import store from '@/store';
import { BvModal } from 'bootstrap-vue';
    /* eslint-disable no-alert, no-console */
export default {
  // delete seems to be reserved keyword
  methods: {
  },
  props: [
    'voting'
  ],
  data() {
    return {
      name: "",
      id: "",
      open: 0,
      ended: 0,
      votespertoken: 0,
      state: store.state,
      openOptions: [
        { value: 1, text: "Kyllä" },
        { value: 0, text: "Ei" }
      ],
      votesByCandidate: ((store.state.admin.votes
      .filter((vote: PureVote) => {
        return vote.VotingId === this.voting.Id
      }) || [{}])
      .reduce((prev: any, curr: PureVote): any => {
        (prev[curr["CandidateId"]] = prev[curr["CandidateId"]] || []).push(curr)
        return prev
      }, {...(store.state.admin.candidates.filter((candidate: PureCandidate) => {
          return store.state.admin.availabilities.find((availability: PureAvailability) => {
            return availability.CandidateId === candidate.Id && availability.VotingId === this.voting.Id
          }) !== undefined
        }).reduce((candidate: any, item: any): any => {
          let r: any = {}
          r[item.Id] = []
          return r
        }, {}))
      }) as VotingStatusMap),
      candidatesHere: store.state.admin.candidates.filter((a: PureCandidate) => {
        let availabilities = store.state.admin.availabilities
        return availabilities.find((b: PureAvailability) => b.VotingId === this.voting.Id && b.CandidateId === a.Id)
      }),
      candidates: [
        ...store.state.admin.candidates.map((a: PureCandidate) => {return {Id: a.Id, Name: a.Name}})
      ],
      endedOptions: [
        { value: 1, text: "Kyllä" },
        { value: 0, text: "Ei" }
      ]
    }
  },
  components: {
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.candidates {
  width: 400px;
}
.valueholder {
  padding: 3px;
  padding-left: 6px;
  padding-right: 6px;
  margin-right: -1px;
  border: 1px solid black;
  text-align: left;
}
.col1 {
  width: 200px;
}
.col20 {
  display: flex;
}
.col21 {
  min-width: 80%;
}
.col22 {
  text-align: right;
  min-width: 20%;
  flex-direction: row-reverse;
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

