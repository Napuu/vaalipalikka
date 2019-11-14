<template>
  <div>
    <b-button variant="warning" @click="emit('modify')">Muokkaa</b-button>
    <b-button variant="info" @click="emit('update')">Tallenna</b-button>
    <b-button variant="danger" @click="emit('delete')">Poista</b-button>
  </div>
</template>
<script lang="ts">
import Vue from 'vue'
export default{
  methods: {
    emit(msg: string) {
      this.$emit(msg)
    }
  }
}
</script>
<style scoped>

</style>