import Vue, { VNode } from 'vue'

declare global {
  namespace JSX {
    interface Element extends VNode {}
    // tslint:disable no-empty-interface
    interface ElementClass extends Vue {}
    interface IntrinsicElements {
      [elem: string]: any
    }
  }
  interface VotingStatusMap {
    [CandidateId: string]: PureVote[]
  }
  interface Voting {
    Name: string
    Id: string
    Description: string
    Open: number
    Ended: number
    VotesLeft: number
    Candidates: Array<Candidate>
  }
  interface Candidate {
    Name: string
    Id: string
    Description: string
    Voted: Boolean
  }
  interface Vote {
    Id: string
    VotingId: string
    CandidateId: string
    Token: string
  }
  interface PureVoting {
    Name: string
    Id: string
    Description: string
    Open: number
    Ended: number
    Visible: number
  }
  interface PureCandidate {
    Name: string
    Id: string
    Description: string
  }
  interface PureVote extends Vote {}
  interface PureToken {
    Value: string
    Valid: number
  }
  interface PureAvailability {
    CandidateId: string
    VotingId: string
  }
}
