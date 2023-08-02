import { SportsOrg } from './sports-org'

export interface TeamScoreboardSlot{
    id: string;
    place: number;
    sportsOrganization: SportsOrg;
    teamNumber: number;
    apparatusTotalScores: Map<string, number>;
    totalScore: number;
}