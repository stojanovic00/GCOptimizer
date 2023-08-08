import {SportsOrg} from './sports-org'
import {Apparatus} from './apparatus'

export interface ContestantScoring{
    id : string;
    competingId : number;
    fullName : string;
    sportsOrganization : SportsOrg;
    teamNumber : number;
    ageCategory : string;
}

