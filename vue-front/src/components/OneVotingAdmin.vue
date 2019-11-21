
<template>
  <div class="onevoting">
    <div class="col1">
      <div class="valueholder">
        <div class="title">Nimi</div>
        <div v-if="state.admin.editableVoting.Id !== voting.Id" class="value">{{voting.Name}}</div>
        <b-form-input v-else v-model="name" />
      </div>
      <div class="valueholder">
        <div class="title">Avoinna</div>
        <div v-if="state.admin.editableVoting.Id !== voting.Id" class="value">{{voting.Open}}</div>
        <b-form-select v-else v-model="open" :options="openOptions"/>
      </div>
      <div class="valueholder">
        <div class="title">Päättynyt</div>
        <div v-if="state.admin.editableVoting.Id !== voting.Id" class="value">{{voting.Ended}}</div>
        <b-form-select v-else v-model="ended" :options="endedOptions"/>
      </div>
      <div class="valueholder">
        <div class="title">Sallitut äänet</div>
        <div v-if="state.admin.editableVoting.Id !== voting.Id" class="value">{{voting.VotesPerToken}}</div>
        <b-form-input v-else v-model="votespertoken" type="number"/>
      </div>
      <div class="candidates valueholder">
        <div class="title">Ehdokkaat</div>
        <div v-if="state.admin.editableVoting.Id !== voting.Id" class="value">{{candidatesSelected.map(a => a.Name).join(", ") }}</div>
        <Multiselect v-else v-model="candidatesSelected" :options="candidates" :searchable="false" :multiple="true" :close-on-select="false" :clear-on-select="false"  label="Name" track-by="Id">
          <template slot="selection" >
            <span class="multiselect__single" v-if="candidatesSelected.length">{{ candidatesSelected.map(a => a.Name).join(", ") }}</span>
          </template>
        </Multiselect>
      </div>
    </div>
    <div class="col2">
      <EditButtons @modify="modify" @delete="_delete" @update="update" />
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import 'vue-multiselect/dist/vue-multiselect.min.css'
import store from '@/store';
import { BvModal } from 'bootstrap-vue';
import Multiselect from 'vue-multiselect';
import EditButtons from '@/components/EditButtonsAdmin.vue'
    /* eslint-disable no-alert, no-console */
export default {
  // delete seems to be reserved keyword
  methods: {
    modify() {
      this.name = this.voting.Name
      this.id = this.voting.Id
      this.open = this.voting.Open
      this.ended = this.voting.Ended
      this.votespertoken = this.voting.VotesPerToken
      this.$store.commit("setEditableVoting", {voting: this.voting})
    },
    async update() {
      this.id = this.voting.Id
      store.commit("clearEditableVotingTarget")
      store.dispatch("addOrUpdateVoting", {voting: {
        name: this.name,
        id: this.id,
        open: this.open,
        ended: this.ended,
        votespertoken: this.votespertoken
      }})
      store.dispatch("fetchVotings")
      store.dispatch("clearAndAddAvailabilities", {votingid: this.id, candidates: this.candidatesSelected})
      // setTimeout(() => {
      //   console.log(this.id, this.voting.Id)
      //   console.log(store.state.admin.availabilities)
      //   console.log(this)
      //   const that = this
      //   this.candidatesSelected = store.state.admin.candidates.filter((a: PureCandidate) => {
      //     let availabilities = store.state.admin.availabilities
      //     // console.log(a.Name)
      //     return availabilities.find((b: PureAvailability) => {
      //       // console.log(b.VotingId, this.id, b.CandidateId, a.Id)
      //       console.log(b.VotingId === this.id, b.CandidateId === a.Id)
      //       return b.VotingId === this.id && b.CandidateId === a.Id
      //     })
      //   })
      //   console.log("????")
      //   console.log(this.candidatess)
      // }, 2000)
      // setTimeout(() => {
      //   console.log("now fetching")
      //   store.dispatch("fetchAvailabilities", {options: { "Authorization": store.state.token}}) 
      // }, 1000)
    },
    async _delete() {
      this.$bvModal.msgBoxConfirm(`Poista äänestys "${this.voting.Name}"`, {
        okTitle: 'Ok',
        cancelTitle: 'Peruuta',
      })
      .then(value => {
        if (value) {
          this.$store.dispatch("deleteVoting", {votingid: this.voting.Id})
        }
      })
    }
  },
  props: [
    'voting'
  ],
  // updated() {
  //   this.candidatesSelected = store.state.admin.candidates.filter((a: PureCandidate) => {
  //     let availabilities = store.state.admin.availabilities
  //     return availabilities.find((b: PureAvailability) => b.VotingId === this.voting.Id && b.CandidateId === a.Id)
  //   })
  // },
  mounted() {
    this.candidatesSelected = this.candidatess
    // this.candidatesSelected = store.state.admin.candidates.filter((a: PureCandidate) => {
    //   let availabilities = store.state.admin.availabilities
    //   return availabilities.find((b: PureAvailability) => b.VotingId === this.voting.Id && b.CandidateId === a.Id)
    // })
  },
  computed: {
    candidatess() {
      return store.state.admin.candidates.filter((a: PureCandidate) => {
        let availabilities = store.state.admin.availabilities
        return availabilities.find((b: PureAvailability) => b.VotingId === this.voting.Id && b.CandidateId === a.Id)
      })
    }
  },
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
      candidatesSelected: [],
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
    Multiselect,
    EditButtons
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
