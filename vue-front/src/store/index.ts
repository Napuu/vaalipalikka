import Vue from 'vue'
import Vuex from 'vuex'
import { LayoutPlugin } from 'bootstrap-vue'
import router from '@/router'

Vue.use(Vuex)
    /* eslint-disable no-alert, no-console */
export default new Vuex.Store({
  state: {
    tokenIsVisible: false,
    //debugging
    // authenticated: true,
    // token: "123",
    // role: "admin",
    authenticated: false,
    token: "",
    role: "",
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
      editableCandidate: <PureCandidate>{},
      editableAvailability: <PureAvailability>{},
      editableToken: <PureToken>{},
    }
  },
  mutations: {
    login(state, {role, token}) {
      state.authenticated = true
      state.role = role
      state.token = token
    },
    logout(state) {
      window.localStorage.clear()
      state.authenticated = false
      state.role = ""
      state.token = ""
    },
    setNonAdminVotings(state, {votings}) {
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
    clearEditableVotingTarget(state) {
      state.admin.editableVoting = <PureVoting>{}
    },
    addVoting(state, {voting}) {
      state.admin.votings.push(voting)
    },
    addCandidate(state, {candidate}) {
      state.admin.candidates.push(candidate)
    },
    addAvailability(state, {availability}) {
      state.admin.availabilities.push(availability)
      state.admin.editableAvailability = availability
    },
    setVotings(state, payload: {votings: Array<PureVoting>}) {
      state.admin.votings = payload.votings
    },
    setCandidates(state, payload: {candidates: Array<PureCandidate>}) {
      state.admin.candidates = payload.candidates
    },
    setTokens(state, payload: {tokens: Array<PureToken>}) {
      state.admin.tokens = payload.tokens
    },
    setVotes(state, payload: {votes: Array<PureVote>}) {
      state.admin.votes = payload.votes
    },
    setAvailabilities(state, payload: {availabilities: Array<PureAvailability>}) {
      state.admin.availabilities = payload.availabilities
    },
    setActiveAdminTab(state, {tab}) {
      state.admin.activeTab = tab
    },
    setEditableCandidate(state, {candidate}) {
      state.admin.editableCandidate = candidate
    },
    setEditableAvailability(state, {availability}) {
      state.admin.editableAvailability = availability
    },
    setEditableToken(state, {token}) {
      state.admin.editableToken =token
    },
    setEditableVoting(state, {voting}) {
      state.admin.editableVoting = voting
    },
    clearEditableCandidateTarget(state) {
      state.admin.editableCandidate = <PureCandidate>{}
    },
    updateCandidate(state, payload: {newCandidate: PureCandidate}) {
      state.admin.candidates = state.admin.candidates.map((_candidate) => {
        if (_candidate.Id == payload.newCandidate.Id) return payload.newCandidate
        else return _candidate
      })
    },
    updateVoting(state, payload: {newVoting: PureVoting}) {
      state.admin.votings = state.admin.votings.map((_voting) => {
        if (_voting.Id == payload.newVoting.Id) return payload.newVoting
        else return _voting
      })
    },
    updateAvailability(state, payload: {newAvailability: PureAvailability}) {
      state.admin.availabilities = state.admin.availabilities.map((_availability) => {
        if (_availability.CandidateId == payload.newAvailability.CandidateId && 
          _availability.VotingId == payload.newAvailability.VotingId) {
          return payload.newAvailability
        } 
        else return _availability
      })
    },
    updateToken(state, payload: {newToken: PureToken}) {
      state.admin.tokens = state.admin.tokens.map((_token) => {
        if (_token.Value == payload.newToken.Value) return payload.newToken
        else return _token
      })
    },
    toggleToken(state, payload: {togglableTokenValue: string}) {
      state.admin.tokens = state.admin.tokens.map((token: PureToken) => {
        let t = token
        if (token.Value === payload.togglableTokenValue) {
          if (t.Valid === 0) t.Valid = 1
          else t.Valid = 0
        }
        return t
      })
    },
    deleteCandidate(state, payload: {deletableCandidateId: string}) {
      state.admin.candidates = state.admin.candidates.filter(_candidate => _candidate.Id !== payload.deletableCandidateId)
    },
    deleteVoting(state, payload: {deletableVotingId: string}) {
      state.admin.votings = state.admin.votings.filter(_voting => _voting.Id !== payload.deletableVotingId)
    },
    deleteToken(state, payload: {deletableTokenValue: string}) {
      state.admin.tokens = state.admin.tokens.filter((_token: PureToken) => _token.Value !== payload.deletableTokenValue)
    },
    clearEditableAvailabilityTarget(state) {
      state.admin.editableAvailability = <PureAvailability>{}
    },
    clearEditableTokenTarget(state) {
      state.admin.editableToken = <PureToken>{}
    },
  },
  actions: {
    async login({commit, dispatch}, {token}) {
      fetch("/api?action=login", {headers: {"Authorization": token}}).then(async res => {
        let text = await res.text()
        if (text !== "denied") {
          if (text === "admin") {
            commit("login", {role: "admin", token: token})
            router.push("admin")
            dispatch("fetchAdminViewableData")
          } else {
            let votings = await fetch("/api?action=voter&a=show", {headers: {"Authorization": token}})
            let votingsJson = await votings.json()
            commit("login", {role: "voter", token: token})
            commit("setNonAdminVotings", {votings: votingsJson})
            router.push("voting")
          }
          window.localStorage.setItem("token", token)
        }
      })
    },
    async fetchAdminViewableData({dispatch, state}){
      let options = {headers: {"Authorization": state.token}}
      dispatch("fetchVotes", {options})
      dispatch("fetchCandidates", {options})
      dispatch("fetchTokens", {options})
      dispatch("fetchVotings", {options})
      dispatch("fetchAvailabilities", {options})
    },
    async fetchVotes({commit}, {options}) {
      let _votes = await fetch("/api?action=vote&a=show", options)
      let votes = await _votes.json()
      commit("setVotes", {votes})
    },
    async fetchCandidates({commit}, {options}) {
      let _candidates= await fetch("/api?action=candidate&a=show", options)
      let candidates = await _candidates.json()
      commit("setCandidates", {candidates})
    },
    async fetchTokens({commit}, {options}) {
      let _tokens = await fetch("/api?action=token&a=show", options)
      let tokens = await _tokens.json()
      commit("setTokens", {tokens})
    },
    async fetchVotings({commit}, {options}) {
      let _votings = await fetch("/api?action=voting&a=show", options)
      let votings = await _votings.json()
      commit("setVotings", {votings})
    },
    async fetchAvailabilities({commit}, {options}) {
      let _availability = await fetch("/api?action=availability&a=show", options)
      let availabilities = await _availability.json()
      commit("setAvailabilities", {availabilities})
    },
    async addOrUpdateVoting({state}, {voting}) {
      let newVoting: PureVoting = {
        Name: voting.name,
        Id: voting.id,
        Open: voting.open,
        Ended: voting.ended,
        Description: "",
        VotesPerToken: parseInt(voting.votespertoken)
      }
      let answer = await fetch("/api?action=voting&a=add", {
        headers: {"Authorization": state.token},
        method: "POST",
        body: JSON.stringify(newVoting)
      })
      let answerText = await answer.text()
      if (answerText.indexOf("replaced") !== -1) {
        state.admin.votings = state.admin.votings.map(_voting => {
          if (_voting.Id == voting.id) {
            return newVoting
          } else {
            return _voting
          }
        })
      } else {
        state.admin.votings.push(newVoting)
      }
    },
    async addOrUpdateCandidate({commit, state}, {candidate}) {
      let newCandidate: PureCandidate = {
        Name: candidate.name,
        Id: candidate.id,
        Description: "",
      }
      let answer = await fetch("/api?action=candidate&a=add", {
        headers: {"Authorization": state.token},
        method: "POST",
        body: JSON.stringify(newCandidate)
      })
      let answerText = await answer.text()
      if (answerText.indexOf("replaced") !== -1) {
        commit("updateCandidate", {newCandidate})
      } else {
        state.admin.candidates.push(newCandidate)
      }
    },
    async addOrUpdateAvailability({state, commit}, {availability}) {
      let newAvailability: PureAvailability = {
        CandidateId: availability.candidateid,
        VotingId: availability.votingid
      }
      let answer = await fetch("/api?action=availability&a=add", {
        headers: {"Authorization": state.token},
        method: "POST",
        body: JSON.stringify(newAvailability)
      })
      let answerText = await answer.text()
      if (answerText.indexOf("replaced") !== -1) {
        commit("updateAvailability", {newAvailability})
      } else {
        state.admin.availabilities.push(newAvailability)
      }
    },

    async deleteVoting({commit, state}, {votingid}) {
      await fetch(`/api?action=voting&a=del&t=${votingid}`, {
        headers: {"Authorization": state.token},
      });
      commit("deleteVoting", {deletableVotingId: votingid})
    },
    async deleteCandidate({commit, state}, {candidateid}) {
      await fetch(`/api?action=candidate&a=del&t=${candidateid}`, {
        headers: {"Authorization": state.token},
      });
      commit("deleteCandidate", {deletableCandidateId: candidateid})
    },
    // unused but working
    // async deleteToken(state, {tokenvalue}) {
    //   await fetch(`/api?action=token&a=del&t=${tokenvalue}`, {
    //     headers: {"Authorization": this.state.token},
    //   });
    //   commit("deleteToken", {deletableTokenValue: tokenvalue})
    // },
    async clearAndAddAvailabilities({state}, {votingid, candidates}) {
      let pairs = candidates.map((a: PureCandidate) => { return {CandidateId: a.Id, VotingId: votingid}})
      await fetch(`/api?action=availability&a=clearadd`, {
        headers: {"Authorization": state.token},
        body: JSON.stringify(pairs),
        method: "POST"
      });

    },
    async toggleToken({commit, state}, {tokenvalue}) {
      console.log("tokenvalue", tokenvalue)
      let targetToken: any = state.admin.tokens.filter(a => a.Value === tokenvalue)
      console.log("token" + targetToken)
      console.log("???")
      targetToken = targetToken[0]
      if (targetToken.Valid === 0) {
        await fetch(`/api?action=token&a=toggle&t=${tokenvalue}&v=1`, {
          headers: {"Authorization": state.token}
        })
      } else {
        await fetch(`/api?action=token&a=toggle&t=${tokenvalue}&v=0`, {
          headers: {"Authorization": state.token}
        })
      }
      commit("toggleToken", {togglableTokenValue: tokenvalue}) 
    },

    async generateTokens({state, commit}) {
      if (state.role === "admin") {
        await fetch("api?action=token&a=generate")
        let _tokens = await fetch("/api?action=token&a=show", {
          headers: {"Authorization": state.token}
        })
        let tokens = await _tokens.json()
        commit("setTokens", {tokens})
      }
    }
  },
  modules: {
  }
})
