
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
      <div class="valueholder">
        <Multiselect v-model="candidatesSelected" :options="candidates" :multiple="true" :close-on-select="false" :clear-on-select="false" :preserve-search="true" placeholder="Pick some" label="text" track-by="text">
          <template slot="selection" slot-scope="{ values, search, isOpen }"><span class="multiselect__single" v-if="values.length &amp;&amp; !isOpen">{{ values.length }} options selected</span></template>
        </Multiselect>
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
import store from '@/store';
    /* eslint-disable no-alert, no-console */
export default {
  // delete seems to be reserved keyword
  private _delete() {
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
      store.commit("clearEditableVotingTarget")
      store.dispatch("addOrUpdateVoting", {voting: {
        name: this.name,
        id: (this.id === "" ? (new Date().getTime()).toString() : this.id),
        open: this.open,
        ended: this.ended,
        votespertoken: this.votespertoken
      }})
    // TODO refactor reducer above so state really is SSOT
      setTimeout(() => {
        console.log("??fecth?")
        this.$store.dispatch("fetchVotings", {headers: {"Authorization": "123"}})
        this.$store.dispatch("fetchAvailabilities", {headers: {"Authorization": "123"}})
      }, 1000)
    },
    _delete: (() => {
      this.$bvModal.msgBoxConfirm(`Poista äänestys "${this.voting.Name}"`, {
      okTitle: 'Ok',
      cancelTitle: 'Peruuta',
    })
      .then(value => {
        if (value) {
          this.$store.dispatch("deleteVoting", {votingid: this.voting.Id})
        }
      })
    }).bind(this)

  ]
  private state = this.$store.state;
  private async update() {
    store.commit("clearEditableVotingTarget")
    store.dispatch("addOrUpdateVoting", {voting: {
      name: this.name,
      id: (this.id === "" ? (new Date().getTime()).toString() : this.id),
      open: this.open,
      ended: this.ended,
      votespertoken: this.votespertoken
    }})
    // TODO refactor reducer above so state really is SSOT
    setTimeout(() => {
      console.log("??fecth?")
      this.$store.dispatch("fetchVotings", {headers: {"Authorization": "123"}})
      this.$store.dispatch("fetchAvailabilities", {headers: {"Authorization": "123"}})
    }, 1000)
  },
  props: [
    'voting'
  // @Prop() private voting!: PureVoting;
  ],
  data: () => {
    return {
      name: "",
      id: "",
      open: 0,
      ended: 0,
      votespertoken: 0,
      openOptions: [
        { value: 1, text: "Kyllä" },
        { value: 0, text: "Ei" }
      ],
      endedOptions: [
        { value: 1, text: "Kyllä" },
        { value: 0, text: "Ei" }
      ]
    }
  }
}
/*
export default class OneVotingAdmin extends Vue {
  @Prop() private voting!: PureVoting;
  private name: string = "";
  private id: string = "";
  private open: number = 0;
  private ended: number = 0;
  private votespertoken: number = 0;
  // delete seems to be reserved keyword
  private _delete() {
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
  private state = this.$store.state;
  private modify() {
    this.name = this.voting.Name
    this.id = this.voting.Id
    this.open = this.voting.Open
    this.ended = this.voting.Ended
    this.votespertoken = this.voting.VotesPerToken
    this.$store.commit("setEditableVoting", {voting: this.voting})
  }
  private async update() {
    this.$store.commit("clearEditableVotingTarget")
    this.$store.dispatch("addOrUpdateVoting", {voting: {
      name: this.name,
      id: (this.id === "" ? (new Date().getTime()).toString() : this.id),
      open: this.open,
      ended: this.ended,
      votespertoken: this.votespertoken
    }})
    // TODO refactor reducer above so state really is SSOT
    setTimeout(() => {
      console.log("??fecth?")
      this.$store.dispatch("fetchVotings", {headers: {"Authorization": "123"}})
      this.$store.dispatch("fetchAvailabilities", {headers: {"Authorization": "123"}})
    }, 1000)
  }
  data() {
    return {
      openOptions: [
        { value: 1, text: "Kyllä" },
        { value: 0, text: "Ei" }
      ],
      endedOptions: [
        { value: 1, text: "Kyllä" },
        { value: 0, text: "Ei" }
      ]
    }
  }
  */
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
.candidateListing {
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
