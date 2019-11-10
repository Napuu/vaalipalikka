
<template>
  <div class="onevoting">
    <div class="col1">
      <div>
        <div class="title">Id</div> 
        <div class="value">{{voting.Id}}</div>
      </div>
      <div>
        <div class="title">Nimi</div> 
        <div v-if="state.admin.editableVoting.Id !== voting.Id" class="value">{{voting.Name}}</div>
        <b-form-input v-else v-model="name" />
      </div>
      <div>
        <div class="title">Avoinna</div> 
        <div v-if="state.admin.editableVoting.Id !== voting.Id" class="value">{{voting.Open}}</div>
        <b-form-select v-else v-model="open" :options="openOptions"/>
      </div>
      <div>
        <div class="title">Päättynyt</div> 
        <div v-if="state.admin.editableVoting.Id !== voting.Id" class="value">{{voting.Ended}}</div>
        <b-form-select v-else v-model="ended" :options="endedOptions"/>
      </div>
      <div>
        <div class="title">Sallitut äänet</div> 
        <div v-if="state.admin.editableVoting.Id !== voting.Id" class="value">{{voting.VotesPerToken}}</div>
        <b-form-input v-else v-model="votespertoken" type="number"/>
      </div>
    </div>
    <div class="col2">
      <b-button variant="warning" v-on:click="modify">M</b-button>
      <b-button variant="info" v-on:click="update">T</b-button>
      <b-button variant="danger">P</b-button>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
    /* eslint-disable no-alert, no-console */
@Component
export default class OneVotingAdmin extends Vue {
  @Prop() private voting!: PureVoting;
  private name: String = "";
  private id: string = "";
  private open: number = 0;
  private ended: number = 0;
  private votespertoken: number = 0;

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
    console.log("Updateing")
    let votings = await fetch("/api?action=voting&a=add", {
      headers: {"Authorization": "123"},
      method: "POST",
      body: JSON.stringify({
        Name: this.name,
        Id: this.id,
        Open: this.open,
        Ended: this.ended,
        Description: "",
        VotesPerToken: this.votespertoken
      })
    })
    this.$store.commit("clearEditableVotingTarget")
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
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.row {
}
.candidateListing {
}
.col1 {
  display: flex;
  max-width: 66%;
  flex-wrap: wrap;
}
.col2 {
  display: flex;
  max-width: 33%;
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
