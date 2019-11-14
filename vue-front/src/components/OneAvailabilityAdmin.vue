<template>
  <!-- <div v-if="getVotingName(availability.VotingId) !== '' && getCandidateName(availability.CandidateId) !== ''" class="oneavailability"> -->
  <div class="oneavailability">
    <div class="col1">

      <!-- <div class="valueholder">
        <div class="title">Nimi</div>
        <div v-if="state.admin.editableVoting.Id !== voting.Id" class="value">{{voting.Name}}</div>
        <b-form-input v-else v-model="name" />
      </div> -->

      <div class="valueholder">
        <!-- <b-form-input v-else v-model="name" /> -->
        <div class="title">Äänestys</div>
        <!-- <div v-if="state.admin.editableAvailability.Id !== availability.Id" class="value"> -->
        <div class="value">
          {{voting.Name}}
        </div>
        <!-- <b-form-select v-else v-model="voting" :options="votings" /> -->
      </div>
      <div class="valueholder">
        <div class="title">Ehdokkaat</div>
        <div v-if="state.admin.editableAvailability.CandidateId !== availability.CandidateId || state.admin.editableAvailability.VotingId !== availability.VotingId" class="value">
        <Multiselect v-model="candidatesSelected" :options="candidates" :multiple="true" :close-on-select="false" :clear-on-select="false" :preserve-search="true" placeholder="Pick some" label="text" track-by="text">
          <template slot="selection" slot-scope="{ values, search, isOpen }"><span class="multiselect__single" v-if="values.length &amp;&amp; !isOpen">{{ values.length }} options selected</span></template>
        </Multiselect>
        </div>
        <div v-else>
          {{candidatesSelected}}
        </div>
        <!-- <div v-if="state.admin.editableAvailability.CandidateId !== availability.CandidateId || state.admin.editableAvailability.VotingId !== availability.VotingId" class="value"> -->
        <!-- <div v-if="state.admin.editableAvailability.Id !== availability.Id" class="value"> -->
          <!-- {{getCandidateName(availability.CandidateId)}} -->
        <!-- </div> -->
        <!-- <b-form-select v-else v-model="candidate" :options="candidates" /> -->
      </div>
    </div>
    <div class="col2">
      <b-button variant="warning" v-on:click="modify">M</b-button>
      <b-button variant="info" v-on:click="update">T</b-button>
      <b-button variant="danger" v-on:click="_delete">P</b-button>
    </div>
  </div>

</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import Multiselect from 'vue-multiselect';
import store from '@/store';
    /* eslint-disable no-alert, no-console */
export default {
  props: [
    'voting'
  ],
  // props:
  // @Prop() private voting!: PureVoting;
  // private voting: {value: string, text: string} = {value: "", text: ""}
  // private candidate: {value: string, text: string} = {value: "", text: ""}
  // private voting: string = ""
  // private candidate: string = ""
  // delete seems to be reserved keyword
  // private getCandidateName(candidateId: string): string {
  //   let name = this.$store.state.admin.candidates.filter((c: PureCandidate) => c.Id === candidateId)
  //   return (name.length > 0 ? name[0].Name : "")
  // }
  // private getVotingName(votingId: string): string {
  //   let name = this.$store.state.admin.votings.filter((c: PureVoting) => c.Id === votingId)
  //   return (name.length > 0 ? name[0].Name : "")
  // }
  // private _delete() {

  //   this.$bvModal.msgBoxConfirm(`Ota ehdokas "${this.availability.Name}"`, {
  //     okTitle: 'Ok',
  //     cancelTitle: 'Peruuta',
  //   })
  //     .then(value => {
  //       if (value) {
  //         this.$store.dispatch("deleteAvailability", {availabilityid: this.availability.Id})
  //       }
  //     })
  // }
  // private state = this.$store.state;
  // private modify() {
    // this.candidate = this.availability.CandidateId
    // this.voting = this.availability.VotingId
    // this.candidate = {value: this.availability.CandidateId, text: this.getCandidateName(this.availability.CandidateId)},
    // this.voting = {value: this.availability.CandidateId, text: this.getCandidateName(this.availability.CandidateId)},
    // this.$store.commit("setEditableAvailability", {availability: this.availability})
  // }
  // private async update() {
    // this.$store.commit("clearEditableAvailabilityTarget")
    // this.$store.dispatch("addOrUpdateAvailability", {availability: {
      // candidateid: this.candidate,
      // votingid: this.voting
    // }})
    // TODO refactor reducer above so state really is SSOT
    // this.$store.dispatch("fetchAvailabilities", {headers: {"Authorization": "123"}})
  // }
  data: () => {
    return {
      state: store.state,
      votings: [
        // ...this.$store.state.admin.votings.map((a: PureVoting) => {return {value: a.Id, text: a.Name}})
      ],
      candidatesSelected: [],
      candidates: [
        ...store.state.admin.candidates.map((a: PureCandidate) => {return {value: a.Id, text: a.Name}})
      ]
    }
  },
  components : {Multiselect}
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.valueholder {
  padding: 3px;
  padding-left: 6px;
  padding-right: 6px;
  margin-right: -1px;
  border: 1px solid black;
}
.availabilityListing {
}
.col1 {
  display: flex;
  max-width: 66%;
  min-width: 66%;
  flex-wrap: wrap;
}
.col2 {
  display: flex;
  text-align: right;
  max-width: 33%;
  min-width: 33%;
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
.oneavailability {
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
