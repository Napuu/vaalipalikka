import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)
    /* eslint-disable no-alert, no-console */
export default new Vuex.Store({
  state: {
    tokenIsVisible: false,
    //debugging
    authenticated: true,
    token: "123",
    role: "admin",
    // authenticated: false,
    // token: "",
    // role: "",
    votings: Array<Voting>(),
    probableCandidateId: "",
    probableVotingId: "",
    admin: {
      activeTab: "",
      votings: Array<PureVoting>(),
      votes: Array<PureVote>(),
      candidates: Array<PureCandidate>(),
      tokens: Array<PureToken>(),
      availabilities: Array<PureAvailability>(),
      editableVoting: <PureVoting>{},
    }
  },
  mutations: {
    login(state, {role, token}) {
      state.authenticated = true
      state.role = role
      state.token = token
    },
    logout(state) {
      state.authenticated = false
      state.role = ""
      state.token = ""
    },
    setVotings(state, {votings}) {
      state.votings = votings
    },
    setProbableVotingTarget(state, {votingid, candidateid}) {
      state.probableCandidateId = candidateid
      state.probableVotingId = votingid
    },
    clearProbableVotingTarget(state) {
      state.probableCandidateId = ""
      state.probableVotingId = ""
    },
    toggleTokenVisibility(state) {
      state.tokenIsVisible = !state.tokenIsVisible
    },
    setEditableVoting(state, {voting}) {
      state.admin.editableVoting = voting
    },
    clearEditableVotingTarget(state) {
      state.admin.editableVoting = <PureVoting>{} 
    },
    setActiveAdminTab(state, {tab}) {
      state.admin.activeTab = tab
    }
  },
  actions: {
    async fetchAdminViewableData(state): Promise<Boolean> {
      let options = {headers: {"Authorization": "123"}}
      this.dispatch("fetchVotes", {options})
      this.dispatch("fetchCandidates", {options})
      this.dispatch("fetchTokens", {options})
      this.dispatch("fetchVotings", {options})
      this.dispatch("fetchAvailabilities", {options})
      return true
    },
    async fetchVotes(state, {options}) {
      let _votes = await fetch("/api?action=vote&a=show", options)
      this.state.admin.votes = await _votes.json()
    },
    async fetchCandidates(state, {options}) {
      let _candidates= await fetch("/api?action=candidate&a=show", options)
      this.state.admin.candidates = await _candidates.json()
    },
    async fetchTokens(state, {options}) {
      let _tokens = await fetch("/api?action=token&a=show", options)
      this.state.admin.tokens = await _tokens.json()
    },
    async fetchVotings(state, {options}) {
      let _votings = await fetch("/api?action=voting&a=show", options)
      this.state.admin.votings = await _votings.json()
    },
    async fetchAvailabilities(state, {options}) {
      let _availability = await fetch("/api?action=availability&a=show", options)
      this.state.admin.availabilities = await _availability.json()
    },
    async addOrUpdateVoting(state, {voting}) {
      let newVoting: PureVoting = {
        Name: voting.name,
        Id: voting.id,
        Open: voting.open,
        Ended: voting.ended,
        Description: "",
        VotesPerToken: parseInt(voting.votespertoken)
      }
      let answer = await fetch("/api?action=voting&a=add", {
        headers: {"Authorization": this.state.token},
        method: "POST",
        body: JSON.stringify(newVoting)
      })
      let answerText = await answer.text()
      if (answerText.indexOf("replaced") !== -1) {
        this.state.admin.votings = this.state.admin.votings.map(_voting => {
          if (_voting.Id == voting.id) {
            return newVoting
          } else {
            return _voting
          }
        })
      }
    },
    async deleteVoting(state, {votingid}) {
      await fetch(`/api?action=voting&a=del&t=${votingid}`, {
        headers: {"Authorization": this.state.token},
      });
      this.state.admin.votings = this.state.admin.votings.filter(_voting => {
        return _voting.Id !== votingid
      })
    }
  },
  modules: {
  }
})
