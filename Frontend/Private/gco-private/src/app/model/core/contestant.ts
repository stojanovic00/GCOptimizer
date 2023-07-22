import { DelegationMember } from "./delegation-member";
export interface Contestant{
    delegationMember : DelegationMember;
    dateOfBirth : number; //protobuf...
}

export interface ContestantTable {
    displayedColumns: string[];
    dataSource: Contestant[];
    selectedRow: Contestant | null;
}