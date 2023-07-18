import { DelegationMemberPosition } from "./delegation-member-position";
import { Gender } from "./gender";

export interface DelegationMember {
    id?: string;
    fullName: string;
    email: string;
    gender:Gender; 
    position: DelegationMemberPosition;
} 
