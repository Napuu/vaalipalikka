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
    addVoting(state, {voting}) {
      state.admin.votings.push(voting)
      state.admin.editableVoting = voting
    },
    addCandidate(state, {candidate}) {
      state.admin.candidates.push(candidate)
      state.admin.editableCandidate = candidate
    },
    addAvailability(state, {availability}) {
      state.admin.availabilities.push(availability)
      state.admin.editableAvailability = availability
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
    clearEditableCandidateTarget(state) {
      state.admin.editableCandidate = <PureCandidate>{}
    },
    setAvailabilities(state, {availabilities}) {
      state.admin.availabilities = availabilities
    },
    clearEditableAvailabilityTarget(state) {
      state.admin.editableAvailability = <PureAvailability>{}
    },
    clearEditableTokenTarget(state) {
      state.admin.editableToken = <PureToken>{}
    },
  },
  actions: {
    async login(state, {token}) {
      fetch("/api?action=login", {headers: {"Authorization": token}}).then(async res => {
        let text = await res.text()
        if (text !== "denied") {
          if (text === "admin") {
            this.commit("login", {role: "admin", token: token})
            router.push("admin")
            await this.dispatch("fetchAdminViewableData")
          } else {
            let votings = await fetch("/api?action=voter&a=show", {headers: {"Authorization": token}})
            let votingsJson = await votings.json()
            this.commit("setVotings", {votings: votingsJson})
            this.commit("login", {role: "voter", token: token})
            router.push("voting")
          }
          window.localStorage.setItem("token", token)
        }
      })
    },
    async fetchAdminViewableData(state): Promise<Boolean> {
      let options = {headers: {"Authorization": this.state.token}}
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
      // let a = await _availability.json()
      // state.commit("setAvailabilities", a)
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
      } else {
        this.state.admin.votings.push(newVoting)
      }
    },
    async addOrUpdateCandidate(state, {candidate}) {
      let newCandidate: PureCandidate = {
        Name: candidate.name,
        Id: candidate.id,
        Description: "",
      }
      let answer = await fetch("/api?action=candidate&a=add", {
        headers: {"Authorization": this.state.token},
        method: "POST",
        body: JSON.stringify(newCandidate)
      })
      let answerText = await answer.text()
      if (answerText.indexOf("replaced") !== -1) {
        this.state.admin.candidates = this.state.admin.candidates.map(_candidate => {
          if (_candidate.Id == candidate.id) {
            return newCandidate
          } else {
            return _candidate
          }
        })
      } else {
        this.state.admin.candidates.push(newCandidate)
      }
    },
    async addOrUpdateAvailability(state, {availability}) {
      let newAvailability: PureAvailability = {
        CandidateId: availability.candidateid,
        VotingId: availability.votingid
      }
      let answer = await fetch("/api?action=availability&a=add", {
        headers: {"Authorization": this.state.token},
        method: "POST",
        body: JSON.stringify(newAvailability)
      })
      let answerText = await answer.text()
      if (answerText.indexOf("replaced") !== -1) {
        this.state.admin.availabilities = this.state.admin.availabilities.map(_availability => {
          if (_availability.CandidateId == availability.CandidateId && _availability.VotingId == availability.VotingId) {
            return newAvailability
          } else {
            return _availability
          }
        })
      } else {
        this.state.admin.availabilities.push(newAvailability)
      }
    },

    async deleteVoting(state, {votingid}) {
      await fetch(`/api?action=voting&a=del&t=${votingid}`, {
        headers: {"Authorization": this.state.token},
      });
      this.state.admin.votings = this.state.admin.votings.filter(_voting => {
        return _voting.Id !== votingid
      })
    },
    async deleteCandidate(state, {candidateid}) {
      await fetch(`/api?action=candidate&a=del&t=${candidateid}`, {
        headers: {"Authorization": this.state.token},
      });
      this.state.admin.candidates = this.state.admin.candidates.filter(_candidate => {
        return _candidate.Id !== candidateid
      })
    },
    async deleteToken(state, {tokenvalue}) {
      await fetch(`/api?action=candidate&a=del&t=${tokenvalue}`, {
        headers: {"Authorization": this.state.token},
      });
      this.state.admin.tokens = this.state.admin.tokens.filter(_token => {
        return _token.Value !== tokenvalue
      })
    },
    async clearAndAddAvailabilities(state, {votingid, candidates}) {
      let pairs = candidates.map((a: PureCandidate) => { return {CandidateId: a.Id, VotingId: votingid}})
      await fetch(`/api?action=availability&a=clearadd`, {
        headers: {"Authorization": this.state.token},
        body: JSON.stringify(pairs),
        method: "POST"
      });

    },
    async toggleToken(state, {tokenvalue}) {
      console.log("tokenvalue", tokenvalue)
      let targetToken: any= this.state.admin.tokens.filter(a => a.Value === tokenvalue)
      console.log("token" + targetToken)
      console.log("???")
      targetToken = targetToken[0]
      if (targetToken.Valid === 0) {
        await fetch(`/api?action=token&a=toggle&t=${tokenvalue}&v=1`, {
          headers: {"Authorization": this.state.token}
        })
      } else {
        await fetch(`/api?action=token&a=toggle&t=${tokenvalue}&v=0`, {
          headers: {"Authorization": this.state.token}
        })
      }
      this.state.admin.tokens = this.state.admin.tokens.map((token: PureToken) => {
        let t = token
        if (token.Value === tokenvalue) {
          if (t.Valid === 0) t.Valid = 1
          else t.Valid = 0
        }
        return t
      })
    },
    async generateTokens(state) {
      if (this.state.role === "admin") {
        await fetch("api?action=token&a=generate")
        // use actual action

        let options = {headers: {"Authorization": this.state.token}}
        let _tokens = await fetch("/api?action=token&a=show", options)
        this.state.admin.tokens = await _tokens.json()

      }
    }
  },
  modules: {
  }
})
