import {Gender} from './gender'
import {Address} from './address'
import {SportsOrg} from './sports-org'
import { CompetitionType } from './competition-type';
import { DelegationMemberProposition } from './delegation-member-proposition';
import { TeamComposition } from './team-composition';
import { AgeCategory } from './age-category'

export interface Competition{
    id?: string;
    name: string;
    startDate: number;
    endDate: number;
    gender: Gender;
    type: CompetitionType;
    tiebreak: boolean;
    address: Address;
    organizer?: SportsOrg;

    delegationMemberPropositions?: DelegationMemberProposition[];
    teamComposition: TeamComposition;
    ageCategories?: AgeCategory[];
}


export interface CompetitionTable {
    displayedColumns: string[];
    dataSource: Competition[];
    selectedRow: Competition | null;
}



