<template>
  <div class="onetoken">
    <div class="col1">
      <div class="valueholder">
        <div class="value">{{token.Value}}</div>
      </div>
      <div class="valueholder">
        <b-button class="statusbutton" @click="toggle('off')" v-if="token.Valid === 1" variant="success">Käytössä</b-button>
        <b-button class="statusbutton" @click="toggle('on')" v-else variant="danger">Ei käytössä</b-button>
        <!-- <div v-if="state.admin.editableToken.Value !== token.Value" class="valid">{{token.Valid}}</div> -->
        <!-- <b-form-select v-else v-model="open" :options="validOptions"/> -->
      </div>
    </div>
    <div class="col2">
      <!-- <b-button variant="warning" v-on:click="modify">M</b-button> -->
      <!-- <b-button variant="info" v-on:click="update">T</b-button> -->
      <!-- <b-button variant="danger" v-on:click="_delete">Poista</b-button> -->
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import 'vue-multiselect/dist/vue-multiselect.min.css'
import store from '@/store';
import { BvModal } from 'bootstrap-vue';
import Multiselect from 'vue-multiselect';
    /* eslint-disable no-alert, no-console */
export default {
  // delete seems to be reserved keyword
  methods: {
    toggle(status: string) {
      console.log("toggling status")
      this.$store.dispatch("toggleToken", {tokenvalue: this.token.Value})
    },
    modify() {
      this.value = this.token.Value
      this.valid = this.token.Valid
      this.$store.commit("setEditableToken", {token: this.token})
    },
    async update() {
      store.commit("clearEditableTokenTarget")
      store.dispatch("addOrUpdateToken", {token: {
        value: this.value,
        valid: this.valid,
      }})
    },
    async _delete() {
      this.$bvModal.msgBoxConfirm(`Poista äänestys "${this.token.Name}"`, {
        okTitle: 'Ok',
        cancelTitle: 'Peruuta',
      })
      .then(value => {
        if (value) {
          this.$store.dispatch("deleteToken", {tokenvalue: this.token.Value})
        }
      })
    }
  },
  props: [
    'token'
  ],
  data() {
    return {
      value: "",
      valid: 0,
      state: store.state,
      validOpetions: [
        { value: 1, text: "Kyllä" },
        { value: 0, text: "Ei" }
      ],
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
.value {
  padding-top: 6px;
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
.button {
  margin-bottom: 3px;
}
.statusbutton {
  width: 130px;
}
#title {
  background: grey;
  color: white;
  /* shifting a bit right ":D"
  margin-left: 3px;
  margin-right: -3px;
  */
}
.onetoken {
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
