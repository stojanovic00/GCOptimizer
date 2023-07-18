import {DelegationMemberPosition} from'./delegation-member-position'

export interface DelegationMemberProposition{
    id? : string;
    position : DelegationMemberPosition;
    minNumber : number;
    maxNumber : number;
  }
  