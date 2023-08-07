import {SportsOrg} from './sports-org'
import {Apparatus} from './apparatus'

export interface ContestantScoring{
    id : string;
    competingId : number;
    fullName : string;
    sportsOrganization : SportsOrg;
    competingApparatuses : Apparatus[];
    teamNumber : number;
    ageCategory : string;
}

