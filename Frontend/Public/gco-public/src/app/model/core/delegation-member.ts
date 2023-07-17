import { DelegationMemberPosition } from "./delegation-member-position";

export interface DelegationMember {
    id: string;
    fullName: string;
    email: string;
    position: DelegationMemberPosition;
} 
