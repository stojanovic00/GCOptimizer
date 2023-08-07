import { DelegationMemberPosition } from "./delegation-member-position";
import { Gender } from "./gender";
import { SportsOrg } from "./sports-org";

export interface DelegationMember {
    id?: string;
    fullName: string;
    email: string;
    gender:Gender; 
    position: DelegationMemberPosition;
    sportsOrganisation? : SportsOrg;
} 
