
<template>
  <div class="onecandidate">
    <div class="col1">
      <div class="valueholder">
        <div class="title">Nimi</div>
        <div v-if="state.admin.editableCandidate.Id !== candidate.Id" class="value">{{candidate.Name}}</div>
        <b-form-input v-else v-model="name" />
      </div>
    </div>
    <div class="col2">
      <EditButtons @modify="modify" @delete="_delete" @update="update" />
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
    /* eslint-disable no-alert, no-console */
import EditButtons from '@/components/EditButtonsAdmin.vue'
export default {
  props: [
    'candidate'
  ],
  data() {
    return {
      openOptions: [
        { value: 1, text: "Kyllä" },
        { value: 0, text: "Ei" }
      ],
      endedOptions: [
        { value: 1, text: "Kyllä" },
        { value: 0, text: "Ei" }
      ],
      name: "",
      id: "",
      state: this.$store.state,
    }
  },
  components: {
    EditButtons
  },
  methods: {
    async _delete() {
    this.$bvModal.msgBoxConfirm(`Poista ehdokas"${this.candidate.Name}"`, {
      okTitle: 'Ok',
      cancelTitle: 'Peruuta',
    })
      .then(value => {
        if (value) {
          this.$store.dispatch("deleteCandidate", {candidateid: this.candidate.Id})
        }
      })
    },
    modify() {
      this.name = this.candidate.Name
      this.id = this.candidate.Id
      this.$store.commit("setEditableCandidate", {candidate: this.candidate})
    },
    async update() {
      this.$store.commit("clearEditableCandidateTarget")
      this.$store.dispatch("addOrUpdateCandidate", {candidate: {
        name: this.name,
        id: (this.id === "" ? (new Date().getTime()).toString() : this.id),
        description: ""
      }})
      this.$store.dispatch("fetchCandidates", {headers: {"Authorization": "123"}})
    }
  }
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
.onecandidate {
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
